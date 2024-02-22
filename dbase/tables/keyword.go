package tables

// 关键词

type Keyword struct {
	Rd        uint   `gorm:"primaryKey"` // 主键
	Roomid    string `gorm:"index"`      // 群聊 id
	Phrase    string `gorm:"index"`      // 词语或短语
	Level     int32  // 优先级等级
	CreatedAt int64  // 创建时间戳
	UpdatedAt int64  // 最后更新时间戳
}
