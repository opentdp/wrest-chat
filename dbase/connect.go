package dbase

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/dbase/tables"
)

func Connect() {

	// 连接数据库
	dborm.Connect(&dborm.Config{
		Type:   "sqlite",
		DbName: "wrest.db3",
	})

	// 实施自动迁移
	dborm.Db.AutoMigrate(
		&tables.Chatroom{},
		&tables.Contact{},
		&tables.Keyword{},
		&tables.LLModel{},
		&tables.Message{},
		&tables.Profile{},
	)

	// 全局配置迁移
	setting.DataMigrate()

}
