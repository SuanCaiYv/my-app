package service

import (
	"bytes"
	"fmt"
	"github.com/SuanCaiYv/my-app-backend/nosql"
	"github.com/SuanCaiYv/my-app-backend/util"
	"strings"
	"time"
)

var scheduledLogger = util.NewLogger()
var redisOps = nosql.NewRedisClient()

type Task interface {
	Run()
}

type EmptyTask struct {
}

func (e *EmptyTask) Run() {
}

func (e *EmptyTask) MarshalBinary() ([]byte, error) {
	return []byte("{}"), nil
}

func (e *EmptyTask) UnmarshalBinary(i []byte) error {
	return nil
}

type ClearDraftArticle struct {
	ArticleId string
}

func (c *ClearDraftArticle) MarshalBinary() (data []byte, err error) {
	builder := strings.Builder{}
	builder.WriteString("{\"articleId\":\"")
	builder.WriteString(c.ArticleId)
	builder.WriteString("\"}")
	fmt.Println(builder.String())
	return []byte(builder.String()), nil
}

func (c *ClearDraftArticle) UnmarshalBinary(data []byte) error {
	ss := bytes.Split(data[1:len(data)-1], []byte(":"))
	c.ArticleId = string(ss[1][1 : len(ss[1])-1])
	return nil
}

func (c *ClearDraftArticle) Run() {
	fmt.Println("run task")
	//articleDao := db.NewArticleDaoService()
	//article, err := articleDao.Select(c.ArticleId)
	//if err != nil {
	//	scheduledLogger.Errorf("ClearDraftArticle.Run() error: %v", err)
	//	return
	//}
	//if article.Content == "" && article.Kind.Name == "" && len(article.TagList) == 0 {
	//	err := articleDao.Delete0(c.ArticleId)
	//	if err != nil {
	//		scheduledLogger.Errorf("ClearDraftArticle.Run() error: %v", err)
	//		return
	//	}
	//}
}

func AddTask(task Task, timeToRun time.Time) {
	err := redisOps.PushSortQueue("scheduled_task", task, float64(timeToRun.UnixMilli()))
	if err != nil {
		scheduledLogger.Errorf("AddTask() error: %v", err)
	}
	RunTask()
}

func RunTask() {
	z, err := redisOps.PeeksSortQueue("scheduled_task")
	if err != nil {
		scheduledLogger.Errorf("RunTask() error: %v", err)
		return
	}
	if time.UnixMilli(int64(z.Score)).After(time.Now()) {
		return
	} else {
		task, err := redisOps.PopSortQueue("scheduled_task")
		if err != nil {
			scheduledLogger.Errorf("RunTask() error: %v", err)
			return
		}
		fmt.Println(task.Member)
		task.Member.(Task).Run()
		RunTask()
	}
}
