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
	"strconv"
	"strings"
)

type StaticSrcApi interface {
	ADownloadFile(context *gin.Context)

	UploadFile(context *gin.Context)

	DeleteFile(context *gin.Context)

	ExistFile(context *gin.Context)

	ListFile(context *gin.Context)
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
		s.logger.Errorf("获取文件失败: %lastTaskStatus; %v", username, err)
		context.JSON(200, resp.NewBadRequest("获取文件失败，文件头应为file"))
		return
	}
	metaMap := make(map[string]interface{})
	uri := context.Request.RequestURI
	idx := strings.LastIndex(uri, "?") + 1
	if idx != 0 {
		queryStr := uri[idx:]
		if queryStr != "" {
			queries := strings.Split(queryStr, "&")
			for _, str := range queries {
				ss := strings.Split(str, "=")
				metaMap[ss[0]] = ss[1]
			}
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
		s.logger.Errorf("打开文件失败: %lastTaskStatus; %v", username, err)
		context.JSON(200, resp.NewInternalError("打开文件失败"))
		return
	}
	content := make([]byte, formFile.Size, formFile.Size)
	_, err = file.Read(content)
	if err != nil {
		s.logger.Errorf("读取文件失败: %lastTaskStatus; %v", username, err)
		context.JSON(200, resp.NewInternalError("读取文件失败"))
		return
	}
	newFilename := util.GenerateUUID() + path.Ext(formFile.Filename)
	err = s.gridFSDao.UploadFile(content, newFilename, metaMap)
	if err != nil {
		s.logger.Errorf("写入文件失败: %lastTaskStatus; %v", username, err)
		context.JSON(200, resp.NewInternalError("写入文件失败"))
		return
	}
	context.JSON(200, resp.NewOk(struct {
		Filename string `json:"filename"`
	}{Filename: newFilename}))
}

func (s *StaticSrcApiHandler) DeleteFile(context *gin.Context) {
	username := context.MustGet("username")
	filename := context.Param("filename")
	err := s.gridFSDao.DeleteFile(filename)
	if err != nil {
		s.logger.Errorf("删除文件失败: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("删除文件失败"))
		return
	}
	context.JSON(200, resp.NewBoolean(true))
}

func (s *StaticSrcApiHandler) ExistFile(context *gin.Context) {
	username := context.MustGet("username")
	input := make(map[string]interface{})
	err := context.BindJSON(&input)
	if err != nil {
		s.logger.Errorf("参数绑定失败: %lastTaskStatus; %v", username, err)
		context.JSON(200, resp.NewBadRequest("参数绑定失败"))
		return
	}
	existFile := s.gridFSDao.ExistFile(input["filename"].(string))
	context.JSON(200, resp.NewBoolean(existFile))
}

func (s *StaticSrcApiHandler) ListFile(context *gin.Context) {
	username := context.MustGet("username").(string)
	pageSize, err := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	if err != nil {
		s.logger.Errorf("获取分页大小失败: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("获取分页大小失败"))
		return
	}
	pageNum, err := strconv.Atoi(context.DefaultQuery("page_num", "1"))
	if err != nil {
		s.logger.Errorf("获取分页页码失败: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("获取分页页码失败"))
		return
	}
	archive := context.DefaultQuery("archive", "other")
	if err != nil {
		s.logger.Errorf("获取分页页码失败: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("获取分页页码失败"))
		return
	}
	var list []string
	var total int64
	var endPage bool
	if pageNum == -1 {
		list, err = s.gridFSDao.ListByArchive0(archive)
		if err != nil {
			s.logger.Errorf("获取归档列表失败: %s; %v", "no-auth", err)
			context.JSON(200, resp.NewInternalError("获取归档列表失败"))
			return
		}
		endPage = true
		total = int64(len(list))
	} else {
		s.logger.Info(pageNum, pageSize)
	}
	context.JSON(200, resp.NewList(total, int64(len(list)), int64(pageNum), int64(pageSize), int64(pageNum+1), endPage, list))
}
