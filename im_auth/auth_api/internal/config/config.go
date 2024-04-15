package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	Redis struct {
		Addr     string
		UserName string
		Pwd      string
		DB       int
	}
	Auth struct {
		AccessSecret string
		AccessExpire int
	}
}
