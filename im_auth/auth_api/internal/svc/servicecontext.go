package svc

import (
	"IM_Server/core"
	"IM_Server/im_auth/auth_api/internal/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlClient := core.InitMysql(c.Mysql.DataSource)
	redisClient := core.InitRedis(c.Redis.Addr, c.Redis.UserName, c.Redis.Pwd, c.Redis.DB)
	//mysql.AutoMigrate(&)
	return &ServiceContext{
		Config: c,
		DB:     mysqlClient,
		Redis:  redisClient,
	}
}
