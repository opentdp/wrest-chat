package tables

// 关键词

type Keyword struct {
	Rd        uint   `gorm:"primaryKey" json:"rd"` // 主键
	Roomid    string `gorm:"index" json:"roomid"`  // 群聊 id
	Phrase    string `gorm:"index" json:"phrase"`  // 词语或短语
	Level     int32  `json:"level"`                // 优先级等级
	CreatedAt int64  `json:"created_at"`           // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`           // 最后更新时间戳
}
