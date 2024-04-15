package core

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

//func InitMysql() *gorm.DB {
//
//	dsn := "root:mysql_MtJE6x@tcp(154.8.197.123:3306)/im_server?charset=utf8mb4&parseTime=True&loc=Local"
//
//	var mysqlLogger logger.Interface
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: mysqlLogger,
//	})
//	if err != nil {
//		log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
//	}
//	sqlDB, _ := db.DB()
//	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
//	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
//	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout
//	return db
//}

func InitMysql(MysqlDataSource string) *gorm.DB {
	var mysqlLogger logger.Interface
	db, err := gorm.Open(mysql.Open(MysqlDataSource), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", MysqlDataSource))
	} else {
		logx.Info("Mysql init success")
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout
	return db
}
