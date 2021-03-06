package service

import (
	"github.com/SuanCaiYv/my-app-backend/config"
	"github.com/SuanCaiYv/my-app-backend/db"
	"github.com/SuanCaiYv/my-app-backend/entity"
	"github.com/SuanCaiYv/my-app-backend/entity/resp"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/go-ego/gse"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type ArticleApi interface {
	AddArticle(context *gin.Context)

	UploadDraft(context *gin.Context)

	UpdateArticle(context *gin.Context)

	DeleteArticle(context *gin.Context)

	GetArticle(context *gin.Context)

	ListArticleNoAuth(context *gin.Context)

	ListArticle(context *gin.Context)

	ListDraft(context *gin.Context)

	ArticleDetail(context *gin.Context)

	ExportArticle(context *gin.Context)

	KindAndTagList(context *gin.Context)

	AddKind(context *gin.Context)

	AddTag(context *gin.Context)

	KindList(context *gin.Context)

	TagList(context *gin.Context)

	DeleteKind(context *gin.Context)

	DeleteTag(context *gin.Context)
}

type ArticleApiHandler struct {
	articleDao db.ArticleDao
	kindDao    db.KindDao
	tagDao     db.TagDao
	gridFsDao  db.GridFSDao
	sysUserDao db.SysUserDao
	cutter     *gse.Segmenter
	logger     *logrus.Logger
}

func NewArticleApiHandler() *ArticleApiHandler {
	cutter := &gse.Segmenter{}
	err := cutter.LoadDictEmbed("zh")
	util.JustPanic(err)
	err = cutter.LoadDictEmbed("zh")
	util.JustPanic(err)
	return &ArticleApiHandler{
		articleDao: db.NewArticleDaoService(),
		kindDao:    db.NewKindDaoService(),
		tagDao:     db.NewTagDaoService(),
		gridFsDao:  db.NewGridFSDaoService(),
		sysUserDao: db.NewSysUserDaoService(),
		cutter:     cutter,
		logger:     util.NewLogger(),
	}
}

type newArticle struct {
	ArticleId   string   `json:"article_id"`
	ArticleName string   `json:"article_name"`
	Summary     string   `json:"summary"`
	CoverImg    string   `json:"cover_img"`
	Content     string   `json:"content"`
	Kind        string   `json:"kind"`
	TagList     []string `json:"tag_list"`
	Visibility  int      `bson:"visibility" json:"visibility"`
}

func (a *ArticleApiHandler) AddArticle(context *gin.Context) {
	username := context.MustGet("username").(string)
	input := &newArticle{}
	err := context.Bind(input)
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	article, err := a.articleDao.Select(input.ArticleId)
	if err != nil {
		a.logger.Errorf("????????????Article?????????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("?????????????????????"))
		return
	}
	kind, err := a.kindDao.Select(input.Kind)
	if err != nil {
		a.logger.Errorf("????????????ArticleKind?????????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("?????????????????????"))
		return
	}
	tagList := make([]entity.Tag, 0, len(input.TagList))
	for i := range input.TagList {
		tag, err := a.tagDao.Select(input.TagList[i])
		if err != nil {
			a.logger.Errorf("????????????ArticleTag?????????: %s; %v", username, err)
			context.JSON(200, resp.NewInternalError("?????????????????????"))
			return
		}
		tagList = append(tagList, *tag)
	}
	article.Name = input.ArticleName
	article.Summary = input.Summary
	article.CoverImg = input.CoverImg
	article.Content = input.Content
	article.Catalog = entity.Catalog{
		Name:     "",
		Children: []entity.Catalog{},
	}
	article.Kind = *kind
	article.TagList = tagList
	article.Visibility = input.Visibility
	arr1 := a.cutter.CutSearch(article.Name)
	arr2 := a.cutter.CutSearch(article.Content)
	ss1 := make([]string, 0, len(arr1))
	ss2 := make([]string, 0, len(arr2))
	for i := range arr1 {
		if arr1[i] == " " {
			continue
		}
		ss1 = append(ss1, arr1[i])
	}
	for i := range arr2 {
		if arr2[i] == " " {
			continue
		}
		ss2 = append(ss2, arr2[i])
	}
	arr3 := a.cutter.Cut(article.Name)
	arr4 := a.cutter.Cut(article.Content)
	ss3 := make([]string, 0, len(arr3))
	ss4 := make([]string, 0, len(arr4))
	for i := range arr3 {
		if arr3[i] == " " {
			continue
		}
		ss3 = append(ss3, arr3[i])
	}
	for i := range arr4 {
		if arr4[i] == " " {
			continue
		}
		ss4 = append(ss4, arr4[i])
	}
	article.FulltextTitle = strings.Join(ss1, " ") + strings.Join(ss3, " ")
	article.FulltextContent = strings.Join(ss2, " ") + strings.Join(ss4, " ")
	if article.ReleaseTime.IsZero() {
		article.ReleaseTime = time.Now()
	}
	err = a.articleDao.Update(article)
	if err != nil {
		a.logger.Errorf("??????Article??????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("??????????????????"))
		return
	}
	context.JSON(200, resp.NewOk(struct {
		ArticleId string `json:"article_id"`
	}{ArticleId: article.Id}))
}

type articleDraft struct {
	ArticleId      string `json:"article_id"`
	ArticleName    string `json:"article_name"`
	ArticleContent string `json:"article_content"`
}

func (a *ArticleApiHandler) UploadDraft(context *gin.Context) {
	username := context.MustGet("username").(string)
	input := &articleDraft{}
	err := context.Bind(input)
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	user, err := a.sysUserDao.SelectByUsername(username)
	if err != nil {
		a.logger.Errorf("????????????SysUser?????????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("?????????????????????"))
		return
	}
	if input.ArticleId == "" {
		if input.ArticleName == "" {
			input.ArticleName = time.Now().String()
		}
		article := entity.Article{
			Name:     input.ArticleName,
			Author:   user.Id,
			Summary:  "",
			CoverImg: "",
			Catalog: entity.Catalog{
				Name:     "",
				Children: []entity.Catalog{},
			},
			Content:     input.ArticleContent,
			Kind:        entity.Kind{},
			TagList:     make([]entity.Tag, 0, 0),
			ReleaseTime: time.Now(),
			Visibility:  entity.VisibilityDraft,
			Available:   false,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		err := a.articleDao.Insert(&article)
		if err != nil {
			a.logger.Errorf("??????Article??????: %s; %v", username, err)
			context.JSON(200, resp.NewInternalError("?????????????????????"))
			return
		}
		input.ArticleId = article.Id
	} else {
		article, err := a.articleDao.Select(input.ArticleId)
		if err != nil {
			a.logger.Errorf("????????????Article?????????: %s; %v", username, err)
			context.JSON(200, resp.NewInternalError("?????????????????????"))
			return
		}
		if input.ArticleName != "" {
			article.Name = input.ArticleName
		}
		article.Content = input.ArticleContent
		err = a.articleDao.Update(article)
		if err != nil {
			a.logger.Errorf("??????Article??????: %s; %v", username, err)
			context.JSON(200, resp.NewInternalError("?????????????????????"))
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
	username := context.MustGet("username").(string)
	input := make(map[string]interface{})
	err := context.BindJSON(&input)
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	article, err := a.articleDao.Select(input["article_id"].(string))
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("????????????"))
		return
	}
	util.UpdateStructObjectWithJsonTag(article, input)
	arr1 := a.cutter.CutAll(article.Name)
	arr2 := a.cutter.CutAll(article.Content)
	ss1 := make([]string, 0, len(arr1))
	ss2 := make([]string, 0, len(arr2))
	for i := range arr1 {
		if arr1[i] == " " {
			continue
		}
		ss1 = append(ss1, arr1[i])
	}
	for i := range arr2 {
		if arr2[i] == " " {
			continue
		}
		ss2 = append(ss2, arr2[i])
	}
	arr3 := a.cutter.Cut(article.Name)
	arr4 := a.cutter.Cut(article.Content)
	ss3 := make([]string, 0, len(arr3))
	ss4 := make([]string, 0, len(arr4))
	for i := range arr3 {
		if arr3[i] == " " {
			continue
		}
		ss3 = append(ss3, arr3[i])
	}
	for i := range arr4 {
		if arr4[i] == " " {
			continue
		}
		ss4 = append(ss4, arr4[i])
	}
	article.FulltextTitle = strings.Join(ss1, " ") + strings.Join(ss3, " ")
	article.FulltextContent = strings.Join(ss2, " ") + strings.Join(ss4, " ")
	err = a.articleDao.Update(article)
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("??????????????????"))
		return
	}
	context.JSON(200, resp.NewBoolean(true))
}

func (a *ArticleApiHandler) DeleteArticle(context *gin.Context) {
	username := context.MustGet("username")
	articleId := context.Param("article_id")
	err := a.articleDao.Delete(articleId)
	if err != nil {
		a.logger.Errorf("??????Article??????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("?????????????????????"))
		return
	}
	context.JSON(200, resp.NewBoolean(true))
}

func (a *ArticleApiHandler) GetArticle(context *gin.Context) {
	articleId := context.Param("article_id")
	article, err := a.articleDao.Select(articleId)
	if err != nil {
		a.logger.Errorf("??????????????????: %v", err)
		context.JSON(200, resp.NewInternalError("??????????????????"))
		return
	}
	if article == nil {
		a.logger.Errorf("??????????????????: %s", articleId)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	if article.Visibility != entity.VisibilityPublic {
		a.logger.Errorf("????????????????????????: %s", articleId)
		context.JSON(200, resp.NewBadRequest("????????????????????????"))
		return
	}
	context.JSON(200, resp.NewOk(article))
}

func (a *ArticleApiHandler) ListArticleNoAuth(context *gin.Context) {
	pageSize, err := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	if err != nil {
		a.logger.Errorf("??????????????????: %v", err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	pageNum, err := strconv.Atoi(context.DefaultQuery("page_num", "1"))
	if err != nil {
		a.logger.Errorf("??????????????????: %v", err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	sort := context.DefaultQuery("sort", "created_time")
	// ????????????
	desc, err := strconv.ParseBool(context.DefaultQuery("desc", "true"))
	if err != nil {
		a.logger.Errorf("??????????????????: %v", err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	searchKey := context.DefaultQuery("search_key", "")
	tagIds := context.DefaultQuery("tag_id_list", "")
	var tagIdList []string
	if tagIds != "" {
		tagIdList = strings.Split(tagIds, ",")
	} else {
		tagIdList = make([]string, 0, 0)
	}
	owner := config.ApplicationConfiguration().Owner
	user, err := a.sysUserDao.SelectByUsername(owner)
	if err != nil {
		a.logger.Errorf("????????????SysUser?????????: %s; %v", owner, err)
		context.JSON(200, resp.NewInternalError("?????????????????????"))
		return
	}
	articles, total, err := a.articleDao.ListByAuthor(user.Id, entity.VisibilityPublic, true, int64(pageNum), int64(pageSize), sort, desc, tagIdList, searchKey)
	if err != nil {
		a.logger.Errorf("????????????????????????: %s; %v", "no-auth", err)
		context.JSON(200, resp.NewInternalError("????????????????????????"))
		return
	}
	endPage := false
	if int64(pageNum*pageSize) > total {
		endPage = true
	}
	context.JSON(200, resp.NewList(total, int64(len(articles)), int64(pageNum), int64(pageSize), int64(pageNum+1), endPage, articles))
}

func (a *ArticleApiHandler) ListArticle(context *gin.Context) {
	username := context.MustGet("username").(string)
	pageSize, err := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	pageNum, err := strconv.Atoi(context.DefaultQuery("page_num", "1"))
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	sort := context.DefaultQuery("sort", "created_time")
	// ????????????
	desc, err := strconv.ParseBool(context.DefaultQuery("desc", "true"))
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	searchKey := context.DefaultQuery("search_key", "")
	tagIds := context.DefaultQuery("tag_id_list", "")
	var tagIdList []string
	if tagIds != "" {
		tagIdList = strings.Split(tagIds, ",")
	} else {
		tagIdList = make([]string, 0, 0)
	}
	user, err := a.sysUserDao.SelectByUsername(username)
	if err != nil {
		a.logger.Errorf("????????????SysUser?????????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("?????????????????????"))
		return
	}
	articles, total, err := a.articleDao.ListByAuthor(user.Id, entity.VisibilityDraft, false, int64(pageNum), int64(pageSize), sort, desc, tagIdList, searchKey)
	if err != nil {
		a.logger.Errorf("????????????????????????: %s; %v", "no-auth", err)
		context.JSON(200, resp.NewInternalError("????????????????????????"))
		return
	}
	endPage := false
	if int64(pageNum*pageSize) > total {
		endPage = true
	}
	context.JSON(200, resp.NewList(total, int64(len(articles)), int64(pageNum), int64(pageSize), int64(pageNum+1), endPage, articles))
}

func (a *ArticleApiHandler) ListDraft(context *gin.Context) {
	username := context.MustGet("username").(string)
	pageSize, err := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	pageNum, err := strconv.Atoi(context.DefaultQuery("page_num", "1"))
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	user, err := a.sysUserDao.SelectByUsername(username)
	if err != nil {
		a.logger.Errorf("????????????SysUser?????????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("?????????????????????"))
		return
	}
	var list []entity.Article
	var total int64
	var endPage bool
	// ??????????????????????????????????????????????????????bug
	Add("clearDraft", make(Params), time.Now())
	if pageNum == -1 {
		list, err = a.articleDao.ListByAuthor0(user.Id, entity.VisibilityDraft, true)
		if err != nil {
			a.logger.Errorf("????????????????????????: %s; %v", "no-auth", err)
			context.JSON(200, resp.NewInternalError("????????????????????????"))
			return
		}
		endPage = true
		total = int64(len(list))
	} else {
		a.logger.Info(pageNum, pageSize)
	}
	context.JSON(200, resp.NewList(total, int64(len(list)), int64(pageNum), int64(pageSize), int64(pageNum+1), endPage, list))
}

func (a *ArticleApiHandler) ArticleDetail(context *gin.Context) {
	username := context.MustGet("username").(string)
	articleId := context.Param("article_id")
	if articleId == "" {
		a.logger.Errorf("??????ID??????: %s", username)
		context.JSON(200, resp.NewBadRequest("??????ID??????"))
		return
	}
	article, err := a.articleDao.Select(articleId)
	if err != nil {
		a.logger.Errorf("????????????Article?????????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("?????????????????????"))
		return
	}
	context.JSON(200, resp.NewOk(article))
}

func (a *ArticleApiHandler) ExportArticle(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

type addKind struct {
	KindName string `json:"kind_name"`
}

func (a *ArticleApiHandler) AddKind(context *gin.Context) {
	username := context.MustGet("username").(string)
	input := &addKind{}
	err := context.BindJSON(input)
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	temp, err := a.kindDao.SelectByName(input.KindName)
	if err != nil {
		a.logger.Errorf("??????ArticleKind??????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("?????????????????????"))
		return
	}
	if temp != nil {
		a.logger.Infof("??????????????????????????????: %s, %s", username, input.KindName)
		context.JSON(200, resp.NewBadRequest("????????????????????????"))
		return
	}
	kind := entity.Kind{
		Name: input.KindName,
	}
	err = a.kindDao.Insert(&kind)
	if err != nil {
		a.logger.Errorf("??????ArticleKind??????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("?????????????????????"))
		return
	}
	context.JSON(200, resp.NewBoolean(true))
}

type addTag struct {
	TagName string `json:"tag_name"`
}

func (a *ArticleApiHandler) AddTag(context *gin.Context) {
	username := context.MustGet("username").(string)
	input := &addTag{}
	err := context.BindJSON(input)
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	temp, err := a.tagDao.SelectByName(input.TagName)
	if err != nil {
		a.logger.Errorf("??????Articletag??????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("?????????????????????"))
		return
	}
	if temp != nil {
		a.logger.Infof("??????????????????????????????: %s, %s", username, input.TagName)
		context.JSON(200, resp.NewBadRequest("????????????????????????"))
		return
	}
	tag := entity.Tag{
		Name: input.TagName,
	}
	err = a.tagDao.Insert(&tag)
	if err != nil {
		a.logger.Errorf("??????ArticleTag??????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("?????????????????????"))
		return
	}
	context.JSON(200, resp.NewBoolean(true))
}

func (a *ArticleApiHandler) KindAndTagList(context *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (a *ArticleApiHandler) KindList(context *gin.Context) {
	username := "no-auth"
	pageNum, err := strconv.Atoi(context.DefaultQuery("page_num", "1"))
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	pageSize, err := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	a.logger.Info(pageNum, pageSize)
	var list []entity.Kind
	var total int64
	var endPage bool
	if pageNum == -1 {
		list, err = a.kindDao.ListAll()
		if err != nil {
			a.logger.Errorf("????????????????????????: %s; %v", username, err)
			context.JSON(200, resp.NewInternalError("????????????????????????"))
			return
		}
		total = int64(len(list))
		endPage = true
	} else {
		// TODO
	}
	context.JSON(200, resp.NewList(total, int64(len(list)), int64(pageNum), int64(pageSize), int64(pageNum), endPage, list))
}

func (a *ArticleApiHandler) TagList(context *gin.Context) {
	username := "no-auth"
	pageNum, err := strconv.Atoi(context.DefaultQuery("page_num", "1"))
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	pageSize, err := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewBadRequest("??????????????????"))
		return
	}
	a.logger.Info(pageNum, pageSize)
	var list []entity.Tag
	var total int64
	var endPage bool
	if pageNum == -1 {
		list, err = a.tagDao.ListAll()
		if err != nil {
			a.logger.Errorf("????????????????????????: %s; %v", username, err)
			context.JSON(200, resp.NewInternalError("????????????????????????"))
			return
		}
		total = int64(len(list))
		endPage = true
	} else {
		// TODO
	}
	context.JSON(200, resp.NewList(total, int64(len(list)), int64(pageNum), int64(pageSize), int64(pageNum), endPage, list))
}

func (a *ArticleApiHandler) DeleteKind(context *gin.Context) {
	username := context.MustGet("username").(string)
	id := context.Param("kind_id")
	if id == "" {
		a.logger.Errorf("KindId??????: %s", username)
		context.JSON(200, resp.NewBadRequest("KindId??????"))
		return
	}
	err := a.kindDao.Delete(id)
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("??????????????????"))
		return
	}
	context.JSON(200, resp.NewBoolean(true))
}

func (a *ArticleApiHandler) DeleteTag(context *gin.Context) {
	username := context.MustGet("username").(string)
	id := context.Param("tag_id")
	if id == "" {
		a.logger.Errorf("TagId??????: %s", username)
		context.JSON(200, resp.NewBadRequest("TagId??????"))
		return
	}
	err := a.tagDao.Delete(id)
	if err != nil {
		a.logger.Errorf("??????????????????: %s; %v", username, err)
		context.JSON(200, resp.NewInternalError("??????????????????"))
		return
	}
	context.JSON(200, resp.NewBoolean(true))
}
