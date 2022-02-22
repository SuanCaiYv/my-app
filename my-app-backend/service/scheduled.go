package service

import (
	"github.com/SuanCaiYv/my-app-backend/nosql"
	"github.com/SuanCaiYv/my-app-backend/util"
	"math"
	"sync"
	"time"
)

func init() {
	newestTimestamp = math.MaxInt64
	go loopQueue()
}

var scheduledLogger = util.NewLogger()
var redisOps = nosql.NewRedisClient()
var funcMap = make(map[string]func())
var lock = sync.Mutex{}

// 确保完全取消
var cancel = make(chan struct{}, 0)

// 确保顺序执行
var wakeup = make(chan struct{}, 0)
var done = make(chan struct{}, 1)
var newestTimestamp int64

func AddFunc(funcName string, fn func()) {
	defer lock.Unlock()
	lock.Lock()
	funcMap[funcName] = fn
}

func AddTask(funcName string, params map[string]interface{}, timestamp time.Time) {
	params["funcName"] = funcName
	params["_id"] = util.GenerateUUID()
	err := redisOps.PushSortQueue("scheduled_task", &params, float64(timestamp.UnixMilli()))
	if err != nil {
		scheduledLogger.Error("AddTask error:", err)
	}
	wakeup <- struct{}{}
	scheduledLogger.Info("scheduled task added:")
}

// 获取最近的任务并运行
func runTask() {
	params := make(map[string]interface{})
	timestamp0 := new(float64)
	err := redisOps.PopSortQueue("scheduled_task", &params, timestamp0)
	if err != nil {
		scheduledLogger.Error("runTask error:", err)
		return
	}
	if fnName, ok := params["funcName"]; ok {
		fn := funcMap[fnName.(string)]
		go fn()
	}
}

func loopQueue() {
	for {
		select {
		case <-wakeup:
			getNewestTask()
		case <-done:
			getNewestTask()
		}
	}
}

// 之所以设计这么复杂是为了实现任意时刻，系统最多只有一个任务在执行
func getNewestTask() {
	params := make(map[string]interface{})
	timestamp0 := new(float64)
	err := redisOps.PeekSortQueue("scheduled_task", &params, timestamp0)
	if err != nil {
		scheduledLogger.Error("runTask error:", err)
		return
	}
	timestamp := int64(*timestamp0)
	// 更新最新任务
	if timestamp < newestTimestamp {
		if newestTimestamp != math.MaxInt64 {
			scheduledLogger.Info("newest task is canceled")
			// 取消上一个最新任务
			cancel <- struct{}{}
		}
		newestTimestamp = timestamp
		if time.UnixMilli(timestamp).After(time.Now()) {
			scheduledLogger.Info("newest task is scheduled")
			after := time.After(time.Duration(timestamp-time.Now().UnixMilli()) * time.Millisecond)
			go func() {
				select {
				case <-after:
					runTask()
					// 允许执行下一个任务
					newestTimestamp = math.MaxInt64
				case <-cancel:
					scheduledLogger.Info("cancel newest task")
				}
				done <- struct{}{}
			}()
		} else {
			scheduledLogger.Info("newest task is running")
			runTask()
			// 允许执行下一个任务
			newestTimestamp = math.MaxInt64
			done <- struct{}{}
		}
	} else {
		return
	}
}
