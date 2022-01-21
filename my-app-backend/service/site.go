package service

import (
	"github.com/SuanCaiYv/my-app-backend/db"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SiteApi interface {
	BackupSite(context *gin.Context)
}

type SiteApiHandler struct {
	sysUserDao db.SysUserDao
	articleDao db.ArticleDao
	gridFSDao  db.GridFSDao
	logger     *logrus.Logger
}

func (s *SiteApiHandler) BackupSite(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewSiteApiHandler() *SiteApiHandler {
	return &SiteApiHandler{
		sysUserDao: db.NewSysUserDaoService(),
		articleDao: db.NewArticleDaoService(),
		gridFSDao:  db.NewGridFSDaoService(),
		logger:     util.NewLogger(),
	}
}
