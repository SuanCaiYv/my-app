package nosql

import (
	"context"
	"fmt"
	config2 "github.com/SuanCaiYv/my-app-backend/config"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type RedisOps interface {
	Set(key string, value interface{}) error

	SetExp(key string, value interface{}, timeout time.Duration) error

	Get(key string, value interface{}) error

	Del(key string) error
}

type RedisClient struct {
	client *redis.Client
	logger *logrus.Logger
	ctx    context.Context
}

var (
	instanceRedisClient *RedisClient
	onceRedisClient     sync.Once
)

func NewRedisClient() *RedisClient {
	onceRedisClient.Do(newRedisClient)
	return instanceRedisClient
}

func newRedisClient() {
	logger := util.NewLogger()
	config := config2.ApplicationConfiguration()
	client := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%d", config.RedisConfig.Url, config.RedisConfig.Port),
		Username: config.RedisConfig.User,
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.DB,
	})
	ctx := context.Background()
	instanceRedisClient = &RedisClient{
		client,
		logger,
		ctx,
	}
}

func (r *RedisClient) Set(key string, value interface{}) error {
	return r.client.Set(r.ctx, key, value, 0).Err()
}

func (r *RedisClient) SetExp(key string, value interface{}, timeout time.Duration) error {
	return r.client.Set(r.ctx, key, value, timeout).Err()
}

func (r *RedisClient) Get(key string, value interface{}) error {
	return r.client.Get(r.ctx, key).Scan(value)
}

func (r *RedisClient) Del(key string) error {
	//TODO implement me
	panic("implement me")
}
