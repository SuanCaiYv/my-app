package service

import (
	"bytes"
	"github.com/SuanCaiYv/my-app-backend/db"
	"github.com/SuanCaiYv/my-app-backend/entity/resp"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"path"
	"strings"
)

type StaticSrcApi interface {
	ADownloadFile(context *gin.Context)

	UploadFile(context *gin.Context)

	ExistFile(context *gin.Context)
}

type StaticSrcApiHandler struct {
	gridFSDao db.GridFSDao
	logger    *logrus.Logger
}

func NewStaticSrcApiHandler() *StaticSrcApiHandler {
	return &StaticSrcApiHandler{
		gridFSDao: db.NewGridFSDaoService(),
		logger:    util.NewLogger(),
	}
}

func (s *StaticSrcApiHandler) ADownloadFile(context *gin.Context) {
	filename := context.Param("filename")
	data, _, err := s.gridFSDao.DownloadFile(filename)
	if err != nil {
		s.logger.Errorf("下载文件: %s 失败", filename)
		context.AbortWithStatus(500)
		return
	}
	reader := bytes.NewReader(data)
	contentLength := len(data)
	contentType := util.MIMEType(filename)

	extraHeaders := map[string]string{
		"Content-Disposition": "attachment; filename=" + `"` + filename + `"`,
	}

	context.DataFromReader(http.StatusOK, int64(contentLength), contentType, reader, extraHeaders)
}

func (s *StaticSrcApiHandler) UploadFile(context *gin.Context) {
	username := context.MustGet("username")
	formFile, err := context.FormFile("file")
	if err != nil {
		s.logger.Errorf("获取文件失败: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("获取文件失败，文件头应为file"))
		return
	}
	metaMap := make(map[string]interface{})
	uri := context.Request.RequestURI
	queryStr := uri[strings.LastIndex(uri, "?")+1:]
	if queryStr != "" {
		queries := strings.Split(queryStr, "&")
		for _, str := range queries {
			ss := strings.Split(str, "=")
			metaMap[ss[0]] = ss[1]
		}
	}
	metaMap["upload_user"] = username
	metaMap["origin_name"] = formFile.Filename
	if t, ok := metaMap["archive"]; ok {
		if t == "avatar" || t == "doc_img" {
		} else {
			metaMap["archive"] = "other"
		}
	} else {
		metaMap["archive"] = "other"
	}
	file, err := formFile.Open()
	if err != nil {
		s.logger.Errorf("打开文件失败: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("打开文件失败"))
		return
	}
	content := make([]byte, formFile.Size, formFile.Size)
	_, err = file.Read(content)
	if err != nil {
		s.logger.Errorf("读取文件失败: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("读取文件失败"))
		return
	}
	newFilename := util.GenerateUUID() + path.Ext(formFile.Filename)
	err = s.gridFSDao.UploadFile(content, newFilename, metaMap)
	if err != nil {
		s.logger.Errorf("写入文件失败: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("写入文件失败"))
		return
	}
	context.JSON(200, resp.NewOk(struct {
		Filename string `json:"filename"`
	}{Filename: newFilename}))
}

func (s *StaticSrcApiHandler) ExistFile(context *gin.Context) {
	username := context.MustGet("username")
	input := make(map[string]interface{})
	err := context.BindJSON(&input)
	if err != nil {
		s.logger.Errorf("参数绑定失败: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("参数绑定失败"))
		return
	}
	existFile := s.gridFSDao.ExistFile(input["filename"].(string))
	context.JSON(200, resp.NewBoolean(existFile))
}
