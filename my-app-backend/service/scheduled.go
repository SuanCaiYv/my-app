package service

import (
	"github.com/SuanCaiYv/my-app-backend/nosql"
	"github.com/SuanCaiYv/my-app-backend/util"
	"sync"
	"time"
)

var scheduledLogger = util.NewLogger()
var redisOps = nosql.NewRedisClient()
var funcMap = make(map[string]func())
var lock = sync.Mutex{}

func AddFunc(funcName string, fn func()) {
	defer lock.Unlock()
	lock.Lock()
	funcMap[funcName] = fn
}

func AddTask(funcName string, params map[string]interface{}, timestamp time.Time) {
	params["funcName"] = funcName
	err := redisOps.PushSortQueue("scheduled_task", &params, float64(timestamp.UnixMilli()))
	if err != nil {
		scheduledLogger.Error("AddTask error:", err)
	}
	RunTask()
}

func RunTask() {
	params := make(map[string]interface{})
	timestamp0 := new(float64)
	err := redisOps.PeekSortQueue("scheduled_task", &params, timestamp0)
	if err != nil {
		scheduledLogger.Error("RunTask error:", err)
		return
	}
	timestamp := int64(*timestamp0)
	if time.UnixMilli(timestamp).After(time.Now()) {
		after := time.After(time.Duration(timestamp-time.Now().UnixMilli()) * time.Millisecond)
		go func() {
			<-after
			RunTask()
		}()
	} else {
		err := redisOps.PopSortQueue("scheduled_task", &params, timestamp0)
		if err != nil {
			scheduledLogger.Error("RunTask error:", err)
			return
		}
		if fnName, ok := params["funcName"]; ok {
			fn := funcMap[fnName.(string)]
			fn()
			RunTask()
		}
	}
}
