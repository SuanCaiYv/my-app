package service

import (
	"encoding/base64"
	"encoding/json"
	"github.com/SuanCaiYv/my-app-backend/db"
	"github.com/SuanCaiYv/my-app-backend/entity/resp"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

type SiteApi interface {
	BackupSite(context *gin.Context)
}

type SiteApiHandler struct {
	sysUserDao db.SysUserDao
	articleDao db.ArticleDao
	gridFSDao  db.GridFSDao
	kindDao    db.KindDao
	tagDao     db.TagDao
	logger     *logrus.Logger
}

var (
	siteApiHandler *SiteApiHandler
	siteApiOnce    sync.Once
)

func NewSiteApiHandler() SiteApi {
	newSiteApiHandler()
	return siteApiHandler
}

func newSiteApiHandler() {
	siteApiOnce.Do(func() {
		siteApiHandler = &SiteApiHandler{
			sysUserDao: db.NewSysUserDaoService(),
			articleDao: db.NewArticleDaoService(),
			gridFSDao:  db.NewGridFSDaoService(),
			kindDao:    db.NewKindDaoService(),
			tagDao:     db.NewTagDaoService(),
			logger:     util.NewLogger(),
		}
	})
}

func (s *SiteApiHandler) BackupSite(context *gin.Context) {
	username := context.MustGet("username").(string)
	user, err := s.sysUserDao.SelectByUsername(username)
	if err != nil {
		s.logger.Errorf("无法读取SysUser数据表: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("无法读取用户表"))
		return
	}
	articles, err := s.articleDao.ListAll0(user.Id)
	if err != nil {
		s.logger.Errorf("无法读取Article数据表: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("无法读取文章表"))
		return
	}
	docImgs0, err := s.gridFSDao.ListByArchive0("doc_img")
	if err != nil {
		s.logger.Errorf("无法读取GridFS数据表: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("无法读取GridFS表"))
		return
	}
	avatars0, err := s.gridFSDao.ListByArchive0("avatar")
	if err != nil {
		s.logger.Errorf("无法读取GridFS数据表: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("无法读取GridFS表"))
		return
	}
	kinds, err := s.kindDao.ListAll()
	if err != nil {
		s.logger.Errorf("无法读取Kind数据表: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("无法读取分类表"))
		return
	}
	tags, err := s.tagDao.ListAll()
	if err != nil {
		s.logger.Errorf("无法读取Tag数据表: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("无法读取标签表"))
		return
	}
	docImages := make([]map[string]interface{}, len(docImgs0), len(docImgs0))
	for i := range docImgs0 {
		bytes, m, err := s.gridFSDao.DownloadFile(docImgs0[i])
		if err != nil {
			continue
		}
		docImages[i] = make(map[string]interface{})
		base64Str := base64.StdEncoding.EncodeToString(bytes)
		docImages[i]["file"] = base64Str
		docImages[i]["meta_dada"] = m
	}
	avatars := make([]map[string]interface{}, len(avatars0), len(avatars0))
	for i := range avatars0 {
		bytes, m, err := s.gridFSDao.DownloadFile(avatars0[i])
		if err != nil {
			continue
		}
		avatars[i] = make(map[string]interface{})
		base64Str := base64.StdEncoding.EncodeToString(bytes)
		avatars[i]["file"] = base64Str
		avatars[i]["meta_dada"] = m
	}
	data := make(map[string]interface{})
	data["article_list"] = articles
	data["kind_list"] = kinds
	data["tag_list"] = tags
	data["doc_img_list"] = docImages
	data["avatar_list"] = avatars
	file, err := os.Create("/tmp/backup.json")
	if err != nil {
		s.logger.Errorf("无法创建备份文件: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("无法创建备份文件"))
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			s.logger.Errorf("无法关闭备份文件: %s; %v", username, err)
		}
	}(file)
	bytes, err := json.Marshal(data)
	if err != nil {
		s.logger.Errorf("无法序列化数据: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("无法序列化数据"))
		return
	}
	_, err = file.Write(bytes)
	if err != nil {
		s.logger.Errorf("无法写入数据: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("无法写入数据"))
		return
	}
	context.FileAttachment(file.Name(), time.Now().Format("2006-01-02 15:01:05")+"-backup.json")
}
