package redis

import (
	"context"
	"fmt"
	"sync"

	"github.com/jpc901/disk-common/conf"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	conn *redis.Client
}

var (
	once sync.Once
	rdb *RedisClient
)

func GetRDBInstance() *RedisClient {
	once.Do(func() {
		rdb = &RedisClient{}
	})
	return rdb
}

func (rdb *RedisClient)GetConn() *redis.Client {
	return rdb.conn
}

func (rdb *RedisClient) Init(config conf.RedisConfig) {
	rdb.conn = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})
	rdb.conn.Ping(context.Background()).Result()
}