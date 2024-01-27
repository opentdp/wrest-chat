package model

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wechat-rest/dbms/model"
)

func Connect() {

	// 连接数据库
	dborm.Connect(&dborm.Config{
		Type:   "sqlite",
		DbName: "wrest.db3",
	})

	// 实施自动迁移
	dborm.Db.AutoMigrate(
		&model.Chatroom{},
		&model.Contact{},
		&model.Message{},
		&model.Profile{},
	)

}
