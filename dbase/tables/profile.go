package tables

// 用户配置

type Profile struct {
	Rd        uint   `gorm:"primaryKey"` // 主键
	Wxid      string `gorm:"index"`      // 微信 id
	Roomid    string `gorm:"index"`      // 群聊 id
	Level     int32  // 等级
	Remark    string // 备注
	AiArgot   string // 唤醒词
	AiModel   string // 会话模型
	CreatedAt int64  // 创建时间戳
	UpdatedAt int64  // 最后更新时间戳
}
