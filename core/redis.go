package core

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func InitRedis(addr string, userName string, pwd string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:        addr,
		Username:    userName,        // default user
		Password:    pwd,             // no password set
		DB:          db,              // use default DB
		DialTimeout: 1 * time.Second, // 1 second
	})
	ping := rdb.Ping(context.Background())
	logx.Info("Redis init: ", ping)
	return rdb
}
