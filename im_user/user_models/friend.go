package user_models

import "gorm.io/gorm"

// Friend 好友表
type Friend struct {
	gorm.Model
	SendUserID    uint   `json:"sendUserID"`                     // 发起验证方
	SendUserModel User   `gorm:"foreignKey:SendUserID" json:"-"` // 发起验证方
	RevUserID     uint   `json:"revUserID"`                      // 接受验证方
	RevUserModel  User   `gorm:"foreignKey:RevUserID" json:"-"`  // 接受验证方
	Notice        string `gorm:"size:128" json:"notice"`         // 备注
}
