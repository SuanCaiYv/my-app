package service

import (
	"fmt"
	"github.com/SuanCaiYv/my-app-backend/nosql"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/go-redis/redis/v8"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	pullTask()
}

type Params map[string]interface{}
type Func func(params Params)
type FuncMap map[string]Func

var (
	funcMap                    = make(FuncMap)
	lock                       = sync.Mutex{}
	wakeupSignal               = make(chan struct{}, 0)
	cancelLastTaskSignal       = make(chan struct{}, 1)
	lastTaskStatus       int64 = ReadyStart
	newestTimestamp      int64 = math.MaxInt64
	redisClient                = nosql.NewRedisClient()
	scheduleLogger             = util.NewLogger()
)

const (
	_ = iota
	DoNotStart
	Started
	ReadyStart
)

const (
	FunctionName = "function_name"
	Id           = "_id"
	Timestamp    = "timestamp"
	RedisKey     = "schedule_task"
)

func AddFunction(name string, fn Func) {
	defer lock.Unlock()
	lock.Lock()
	funcMap[name] = fn
}

func Add(funcName string, params Params, timestamp time.Time) {
	params[FunctionName] = funcName
	params[Id] = fmt.Sprintf("%d", timestamp.UnixMicro())
	params[Timestamp] = timestamp.UnixMicro()
	err := redisClient.PushSortQueue(RedisKey, params, float64(timestamp.UnixMicro()))
	if err != nil {
		scheduleLogger.Error(err)
	}
	// 实现串形化添加任务
	wakeupSignal <- struct{}{}
}

func pullTask() {
	go func() {
		for {
			select {
			case <-wakeupSignal:
				choreographyTask()
			}
		}
	}()
}

/*
我来解释一下，取消一个任务可能产生的结果：
1. 待取消任务成功接收取消信号，此时取消成功，同时重新添加任务；当前任务因为已经被消耗，要么直接设置定时器，要么重新添加至队列并再次循环编排方法。
2. 待取消任务未能成功接收取消信号：
	2.1. 任务绑定的方法还未执行，此时判断status；如果为禁止启动则重新添加自己，结束定时器；如果已经启动，则说明取消信号发起者对于status的设置慢了一步，只能继续执行，并设置取消信号，告知取消者此次取消失败。
	2.2. 任务执行中，此时取消者取消失败，只好自主消化取消信号，避免影响下一次使用。
	2.3. 任务绑定的方法已经执行完毕，此时取消者的行为如同取消成功，但是没法回退时间，所以只能默许取消失败，然后重新添加取消者任务。
*/

// 这狗屎逻辑过三天就只有上帝知道它是怎么跑的了
// 每次重新添加任务都必须调用编排方法，防止出现多线程下添加后，定时任务无法被唤醒的bug
// 我不知道有没有潜在的bug，先这么写，应该是没问题了，又是原子操作又是禁止重排序的
func choreographyTask() {
	params := make(Params)
	timestamp0 := new(float64)
	err := redisClient.PopSortQueue(RedisKey, &params, timestamp0)
	if err != nil {
		if err == redis.Nil {
		} else {
			scheduleLogger.Error(err)
		}
		return
	}
	timestamp := int64(*timestamp0)
	if timestamp < newestTimestamp {
		if newestTimestamp == math.MaxInt64 {
			newestTimestamp = timestamp
			d := time.Duration(timestamp-time.Now().UnixMicro()) * time.Microsecond
			timer := time.After(d)
			go func(params Params) {
				select {
				case <-timer:
					if !atomic.CompareAndSwapInt64(&lastTaskStatus, ReadyStart, Started) {
						scheduleLogger.Warnf("schedule task: %s is started too fast! But can be canceled.", params[Id])
						err := redisClient.PushSortQueue(RedisKey, params, params[Timestamp].(float64))
						if err != nil {
							scheduleLogger.Error(err)
						}
						break
					}
					if fnName, ok := params[FunctionName]; ok {
						if fn, ok := funcMap[fnName.(string)]; ok {
							fn(params)
						}
					}
					if !atomic.CompareAndSwapInt64(&lastTaskStatus, Started, ReadyStart) {
						panic("something wrong with schedule task!")
					}
					// 允许下一次执行
					if lastTaskStatus == ReadyStart {
						// 使用if语句产生数据依赖，消除instruction-reorder
						newestTimestamp = math.MaxInt64
					}
					scheduleLogger.Infof("schedule task: %s is finished!", params[Id])
				case <-cancelLastTaskSignal:
					// 记得重新添加回去
					err := redisClient.PushSortQueue(RedisKey, params, params[Timestamp].(float64))
					if err != nil {
						scheduleLogger.Error(err)
					} else {
						// 允许下一次执行，使用else语句产生数据依赖，消除instruction-reorder
						newestTimestamp = math.MaxInt64
					}
					scheduleLogger.Infof("this task: %s is canceled!", params[Id])
				}
				// 调用编排任务方法
				choreographyTask()
			}(params)
		} else {
			// 已经有别的任务在等待调度
			cancelLastTaskSignal <- struct{}{}
			if !atomic.CompareAndSwapInt64(&lastTaskStatus, ReadyStart, DoNotStart) {
				// 取消失败，自行消化取消信号
				<-cancelLastTaskSignal
				// 取消失败后，成功的任务会自行调用编排方法
				scheduleLogger.Warnf("schedule task already started.")
			} else {
				atomic.CompareAndSwapInt64(&lastTaskStatus, DoNotStart, ReadyStart)
				scheduleLogger.Infof("success to cancel schedule task: %s", params[Id])
				// 给自己归位
				err := redisClient.PushSortQueue(RedisKey, params, params[Timestamp].(float64))
				if err != nil {
					scheduleLogger.Error(err)
				}
				// 调用编排任务方法
				choreographyTask()
			}
		}
	} else {
		// 给人家归位
		err := redisClient.PushSortQueue(RedisKey, params, params[Timestamp].(float64))
		if err != nil {
			scheduleLogger.Error(err)
		}
		// 调用编排任务方法
		choreographyTask()
	}
}
