package svc

import (
	"IM_Server/core"
	"IM_Server/im_auth/auth_api/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := core.InitMysql(c.Mysql.DataSource)
	//mysql.AutoMigrate(&)
	return &ServiceContext{
		Config: c,
		DB:     mysql,
	}
}
