package user_models

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	Pwd      string `gorm:"size:64" json:"pwd"`        // 用户密码
	Nickname string `gorm:"size:32" json:"nickname"`   // 用户昵称
	Abstract string `gorm:"size:128"  json:"abstract"` // 用户简介
	Avatar   string `gorm:"size:256"  json:"avatar"`   // 用户头像
	IP       string `gorm:"size:32" json:"ip"`         // 用户ip
	Addr     string `gorm:"size:64" json:"addr"`       // 用户地址
	Role     int    `json:"role"`                      // 角色 1管理员 2普通用户
}
