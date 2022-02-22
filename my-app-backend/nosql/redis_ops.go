package nosql

import (
	"context"
	"encoding/json"
	"fmt"
	config2 "github.com/SuanCaiYv/my-app-backend/config"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type RedisOps interface {
	// Set 以下都是string的操作
	Set(key string, value interface{}) error

	SetExp(key string, value interface{}, timeout time.Duration) error

	Get(key string, value interface{}) error

	Del(key string) error

	// PushSortQueue 以下都是list的操作
	PushSortQueue(key string, value interface{}, score float64) error

	PopSortQueue(key string, value interface{}, score *float64) error

	PeekSortQueue(key string, value interface{}, score *float64) error
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
	return r.client.Del(r.ctx, key).Err()
}

func (r *RedisClient) PushSortQueue(key string, value interface{}, score float64) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return r.client.ZAdd(timeout, key, &redis.Z{
		Score:  score,
		Member: bytes,
	}).Err()
}

func (r *RedisClient) PopSortQueue(key string, value interface{}, score *float64) error {
	results, err := r.client.ZPopMin(r.ctx, key, 1).Result()
	if err != nil {
		return err
	}
	if len(results) == 0 {
		return redis.Nil
	}
	err = json.Unmarshal([]byte(results[0].Member.(string)), value)
	if err != nil {
		return err
	}
	*score = results[0].Score
	return nil
}

func (r *RedisClient) PeekSortQueue(key string, value interface{}, score *float64) error {
	results, err := r.client.ZRangeWithScores(r.ctx, key, 0, 1).Result()
	if err != nil {
		return err
	}
	if len(results) == 0 {
		return redis.Nil
	}
	err = json.Unmarshal([]byte(results[0].Member.(string)), value)
	if err != nil {
		return err
	}
	*score = results[0].Score
	return nil
}
