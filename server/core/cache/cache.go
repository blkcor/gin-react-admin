package cache

import (
	"context"
	"github.com/blkcor/gin-react-admin/config/section"
	"github.com/blkcor/gin-react-admin/core/logger"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func Init() {
	rdb := redis.NewClient(&redis.Options{
		Addr: section.RedisConfig.Host + ":" + section.RedisConfig.Port,
		DB:   section.RedisConfig.Db,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.Error("Redis连接失败: ", err.Error())
		return
	}
	RDB = rdb
	logger.Info("Redis连接已建立!")
}
