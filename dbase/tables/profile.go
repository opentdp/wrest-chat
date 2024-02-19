package tables

// 配置

type Profile struct {
	Rd        uint   `gorm:"primaryKey"`
	Wxid      string `gorm:"uniqueIndex"` // 微信 id
	Roomid    string `gorm:"uniqueIndex"` // 群聊 id
	Level     int32  // 等级
	AiArgot   string // 唤醒词
	AiModel   string // 会话模型
	CreatedAt int64
	UpdatedAt int64
}
