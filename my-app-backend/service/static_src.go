package service

import (
	"bytes"
	"github.com/SuanCaiYv/my-app-backend/db"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"path/filepath"
)

type StaticSrcApi interface {
	ADownloadFile(context *gin.Context)

	UploadFile(context *gin.Context)
}

type StaticSrcApiHandler struct {
	gridFsDao db.GridFSDao
	logger    *logrus.Logger
}

func NewStaticSrcApiHandler() *StaticSrcApiHandler {
	return &StaticSrcApiHandler{
		gridFsDao: db.NewGridFSDaoService(),
		logger:    util.NewLogger(),
	}
}

func (s *StaticSrcApiHandler) ADownloadFile(context *gin.Context) {
	filename := context.Param("filename")
	data, _, err := s.gridFsDao.DownloadFile(filename)
	if err != nil {
		s.logger.Errorf("下载文件: %s 失败", filename)
		context.AbortWithStatus(500)
	}
	reader := bytes.NewReader(data)
	contentLength := len(data)
	// TODO 暂时只让支持图片，其他静态资源可以写个资源类型根判断函数实现
	contentType := "image/" + filepath.Ext(filename)[1:]

	extraHeaders := map[string]string{
		"Content-Disposition": "attachment; filename=" + `"` + filename + `"`,
	}

	context.DataFromReader(http.StatusOK, int64(contentLength), contentType, reader, extraHeaders)
}

func (s *StaticSrcApiHandler) UploadFile(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}
