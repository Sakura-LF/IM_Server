package group_models

import (
	"IM_Server/common/common_Model"
	"gorm.io/gorm"
)

// GroupMsg  群消息表
type GroupMsg struct {
	gorm.Model
	GroupID    uint                    `json:"groupID"`                     // 群id
	Group      Group                   `gorm:"foreignKey:GroupID" json:"-"` // 群
	SendUserID uint                    `json:"sendUserID"`                  // 发送者id
	MsgType    int8                    `json:"msgType"`                     // 消息类型 1 文本类型  2 图片消息  3 视频消息 4 文件消息 5 语音消息  6 语言通话  7 视频通话  8 撤回消息 9回复消息 10 引用消息 11 at消息
	MsgPreview string                  `gorm:"64" json:"msgPreview"`        // 消息预览
	Msg        common_Model.Msg        `json:"msg"`                         // 消息内容
	SystemMsg  *common_Model.SystemMsg `json:"systemMsg"`                   // 系统提示
}
