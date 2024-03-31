package dbase

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/dbase/tables"
)

func Connect() {

	// 连接数据库
	db := dborm.Connect(&dborm.Config{
		Type:   "sqlite",
		DbName: "wrest.db3",
	})

	// 设置默认参数
	db.Exec("PRAGMA foreign_keys=ON;")
	db.Exec("PRAGMA journal_mode=WAL;")
	db.Exec("PRAGMA busy_timeout=5000;")

	// 实施自动迁移
	db.AutoMigrate(
		&tables.Chatroom{},
		&tables.Cronjob{},
		&tables.Contact{},
		&tables.Keyword{},
		&tables.LLModel{},
		&tables.Message{},
		&tables.Profile{},
		&tables.Setting{},
		&tables.Webhook{},
	)

	// 加载全局配置
	setting.DataMigrate()
	setting.Laod()

}
