package api

import (
	"fmt"
	"github.com/SuanCaiYv/my-app-backend/auth"
	"github.com/SuanCaiYv/my-app-backend/entity/resp"
	"github.com/SuanCaiYv/my-app-backend/service"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/gin-gonic/gin"
	"strings"
)

var logger = util.NewLogger()

func Route() {
	router := gin.New()
	router.Use(corsMiddleware())
	router.Use(trackPath())
	router.Use(gin.CustomRecovery(func(context *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			context.AbortWithStatusJSON(200, resp.NewInternalError(err))
		} else {
			context.AbortWithStatusJSON(200, resp.NewInternalError("unknown error occurred."))
		}
	}))
	// ApiHandler实例化
	userApiHandler := service.NewUserApiHandler()
	staticSrcApi := service.NewStaticSrcApiHandler()
	articleApi := service.NewArticleApiHandler()
	siteApi := service.NewSiteApiHandler()
	wsApi := service.NewWSApiHandler()
	// 版本分组
	versionOne := router.Group("/v1")
	// 测试用
	versionOne.GET("/t", func(context *gin.Context) {
		fmt.Println(context.Query("name"))
		context.JSON(200, resp.NewBoolean(true))
	})
	versionOne.POST("/t", func(context *gin.Context) {
		m := make(map[string]interface{})
		_ = context.BindJSON(&m)
		for k, v := range m {
			fmt.Println(k, v)
		}
		context.JSON(200, resp.NewBoolean(true))
	})
	versionOne.POST("/t/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		fmt.Println(file.Filename)
		context.JSON(200, resp.NewBoolean(true))
	})
	{
		// 免登陆部分
		versionOne.GET("/ws", wsApi.Generic)
		versionOne.GET("/article/list/no_auth", articleApi.ListArticleNoAuth)
		versionOne.GET("/user/info/no_auth", userApiHandler.GetUserInfoNoAuth)
		versionOne.GET("/article/kind_list", articleApi.KindList)
		versionOne.GET("/article/tag_list", articleApi.TagList)
		versionOne.PUT("/sign", userApiHandler.Login)
		versionOne.POST("/sign", userApiHandler.SignUp)
		versionOne.POST("/sign/ver_code", userApiHandler.SendVerCode)
		versionOne.GET("/article/:id", articleApi.GetArticle)

		// 静态资源接口
		versionOne.GET("/static/a/:filename", staticSrcApi.ADownloadFile)

		// 以下需要登录
		versionOne.Use(authFunc)

		// 用户接口
		userApi := versionOne.Group("/user")
		userApi.GET("/info", userApiHandler.GetUserInfo)
		userApi.PUT("/info", userApiHandler.UpdateUserInfo)
		userApi.PUT("", func(context *gin.Context) {
			context.JSON(200, resp.NewBoolean(true))
		})
		userApi.DELETE("", userApiHandler.Logout)

		// 文章接口
		article := versionOne.Group("/article")
		article.GET("/list", articleApi.ListArticle)
		article.GET("/doc/:article_id", articleApi.ExportArticle)
		article.GET("/tag_or_kind/list", articleApi.KindAndTagList)
		article.GET("/detail/:article_id", articleApi.ArticleDetail)
		article.GET("/draft/list", articleApi.ListDraft)
		article.GET("/img_fetch", wsApi.ImageFetch)
		article.PUT("", articleApi.UpdateArticle)
		article.POST("", articleApi.AddArticle)
		article.POST("/kind", articleApi.AddKind)
		article.POST("/tag", articleApi.AddTag)
		article.POST("/draft", articleApi.UploadDraft)
		article.DELETE("/:article_id", articleApi.DeleteArticle)
		article.DELETE("/kind/:kind_id", articleApi.DeleteKind)
		article.DELETE("tag/:tag_id", articleApi.DeleteTag)

		// 静态资源接口
		file := versionOne.Group("/static")
		file.GET("/file/list", staticSrcApi.ListFile)
		file.POST("/file", staticSrcApi.UploadFile)
		file.DELETE("/file/:filename", staticSrcApi.DeleteFile)

		// 其他接口
		other := versionOne.Group("")
		other.GET("/backup", siteApi.BackupSite)
	}
	err := router.Run(":8190")
	util.JustPanic(err)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, HEAD")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func trackPath() gin.HandlerFunc {
	return func(context *gin.Context) {
		logger.Info(context.Request.URL.Path)
		context.Next()
	}
}

func authFunc(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if !strings.HasPrefix(token, "Bearer ") {
		context.AbortWithStatusJSON(200, resp.NewMissToken())
		return
	}
	token = token[7:]
	username, role, err := auth.ValidAccessToken(token)
	if err != nil {
		context.AbortWithStatusJSON(200, resp.NewAuthFailed())
		return
	}
	logger.Debugf("用户: %s 开始操作", username)
	context.Set("username", username)
	context.Set("role", role)
}

func todoImplement(context *gin.Context) {
	context.JSON(200, resp.NewString("此接口暂未实现"))
}
