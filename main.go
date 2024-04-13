package main

import (
	"IM_Server/core"
	"IM_Server/im_chat/chat_models"
	"IM_Server/im_group/group_models"
	"IM_Server/im_user/user_models"

	"flag"
	"fmt"
)

type Options struct {
	DB bool
}

func main() {

	var opt Options
	flag.BoolVar(&opt.DB, "db", false, "db")
	flag.Parse()

	if opt.DB {
		db := core.InitMysql("")
		err := db.AutoMigrate(
			&user_models.User{},         // 用户表
			&user_models.Friend{},       // 好友表
			&user_models.FriendVerify{}, // 好友验证表
			&user_models.UserConf{},     // 用户配置表

			&chat_models.Chat{}, // 对话表

			&group_models.Group{},       // 群组表
			&group_models.GroupMember{}, // 群成员表
			&group_models.GroupMsg{},    // 群消息表
			&group_models.GroupVerify{}, // 群验证表
		)
		if err != nil {
			fmt.Println("表结构生成失败", err)
			return
		}
		fmt.Println("表结构生成成功！")
	}
}
