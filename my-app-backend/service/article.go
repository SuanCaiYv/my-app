package service

import (
	"fmt"
	"github.com/SuanCaiYv/my-app-backend/config"
	"github.com/SuanCaiYv/my-app-backend/db"
	"github.com/SuanCaiYv/my-app-backend/entity"
	"github.com/SuanCaiYv/my-app-backend/entity/resp"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type ArticleApi interface {
	AddArticle(context *gin.Context)

	UploadDraft(context *gin.Context)

	UpdateArticle(context *gin.Context)

	DeleteArticle(context *gin.Context)

	ListArticleNoAuth(context *gin.Context)

	ListArticle(context *gin.Context)

	ExportArticle(context *gin.Context)

	KindAndTagList(context *gin.Context)
}

type ArticleApiHandler struct {
	articleDao db.ArticleDao
	gridFsDao  db.GridFSDao
	sysUserDao db.SysUserDao
	logger     *logrus.Logger
}

func NewArticleApiHandler() *ArticleApiHandler {
	return &ArticleApiHandler{
		articleDao: db.NewArticleDaoService(),
		gridFsDao:  db.NewGridFSDaoService(),
		sysUserDao: db.NewSysUserDaoService(),
		logger:     util.NewLogger(),
	}
}

func (a *ArticleApiHandler) AddArticle(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

type articleDraft struct {
	ArticleId      string `json:"article_id"`
	ArticleName    string `json:"article_name"`
	ArticleContent string `json:"article_content"`
}

func (a *ArticleApiHandler) UploadDraft(context *gin.Context) {
	username := context.MustGet("username").(string)
	input := &articleDraft{}
	err := context.BindJSON(input)
	if err != nil {
		a.logger.Errorf("参数解析失败: %v", err)
		context.JSON(200, resp.NewBadRequest("参数解析失败"))
		return
	}
	user, err := a.sysUserDao.SelectByUsername(username)
	if err != nil {
		a.logger.Errorf("无法读取SysUser数据表: %v", err)
		context.JSON(200, resp.NewInternalError("无法读取用户表"))
		return
	}
	fmt.Println(input)
	if input.ArticleId == "" {
		if input.ArticleName == "" {
			input.ArticleName = time.Now().String()
		}
		article := entity.Article{
			Name:        input.ArticleName,
			Author:      user.Id,
			Summary:     "",
			CoverImg:    "",
			Catalog:     entity.Catalog{},
			Content:     input.ArticleContent,
			Kinds:       make([]entity.Kind, 0, 0),
			Tags:        make([]entity.Tag, 0, 0),
			ReleaseTime: time.Now(),
			Visibility:  entity.VisibilityDraft,
			Available:   false,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		err := a.articleDao.Insert(&article)
		if err != nil {
			a.logger.Errorf("插入Article失败: %s; %v", username, err)
			context.JSON(200, resp.NewInternalError("插入文档表失败"))
			return
		}
		input.ArticleId = article.Id
	} else {
		article, err := a.articleDao.Select(input.ArticleId)
		if err != nil {
			a.logger.Errorf("无法读取Article数据表: %v", err)
			context.JSON(200, resp.NewInternalError("无法读取文档表"))
			return
		}
		if input.ArticleName != "" {
			article.Name = input.ArticleName
		}
		article.Content = input.ArticleContent
		err = a.articleDao.Update(article)
		if err != nil {
			a.logger.Errorf("更新Article失败: %s; %v", username, err)
			context.JSON(200, resp.NewInternalError("更新用户文档表失败"))
			return
		}
		input.ArticleId = article.Id
	}
	context.JSON(200, resp.NewOk(struct {
		ArticleId string `json:"article_id"`
	}{
		ArticleId: input.ArticleId,
	}))
}

func (a *ArticleApiHandler) UpdateArticle(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleApiHandler) DeleteArticle(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleApiHandler) ListArticleNoAuth(context *gin.Context) {
	pgSize, _ := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	pgNum, _ := strconv.Atoi(context.DefaultQuery("page_num", "1"))
	sort := context.DefaultQuery("sort", "created_time")
	// 是否倒序
	desc, _ := strconv.ParseBool(context.DefaultQuery("desc", "true"))
	owner := config.ApplicationConfiguration().Owner
	articles, _, err := a.articleDao.ListByAuthor(owner, int64(pgNum), int64(pgSize), sort, desc)
	if err != nil {
		a.logger.Errorf("获取文章列表失败: %v", err)
		context.JSON(200, resp.NewInternalError("获取文章列表失败"))
		return
	}
	context.JSON(200, resp.NewOk(articles))
}

func (a *ArticleApiHandler) ListArticle(context *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (a *ArticleApiHandler) ExportArticle(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleApiHandler) KindAndTagList(context *gin.Context) {
	// TODO implement me
	panic("implement me")
}
