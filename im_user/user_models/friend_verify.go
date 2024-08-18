package user_models

import (
	"IM_Server/common/common_Model"
	"gorm.io/gorm"
)

// FriendVerify  好友验证表
type FriendVerify struct {
	gorm.Model
	SendUserID           uint                               `json:"sendUserID"`                         // 发起验证方
	SendUserModel        User                               `gorm:"foreignKey:SendUserID" json:"-"`     // 发起验证方
	RevUserID            uint                               `json:"revUserID"`                          // 接受验证方
	RevUserModel         User                               `gorm:"foreignKey:RevUserID" json:"-"`      // 接受验证方
	Status               int8                               `json:"status"`                             // 状态 0 未操作 1 同意 2 拒绝 3 忽略
	AdditionalMessages   string                             `gorm:"size:128" json:"additionalMessages"` // 附加消息
	VerificationQuestion *common_Model.VerificationQuestion `json:"verificationQuestion"`               // 验证问题  为3和4的时候需要
}
