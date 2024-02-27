package tables

// 用户配置

type Profile struct {
	Rd        uint   `gorm:"primaryKey" json:"rd"` // 主键
	Wxid      string `gorm:"index" json:"wxid"`    // 微信 id
	Roomid    string `gorm:"index" json:"roomid"`  // 群聊 id
	Level     int32  `json:"level"`                // 等级
	Remark    string `json:"remark"`               // 备注
	AiArgot   string `json:"ai_argot"`             // 唤醒词
	AiModel   string `json:"ai_model"`             // 会话模型
	BanExpire int64  `json:"ban_expire"`           // 禁言截止时间
	CreatedAt int64  `json:"created_at"`           // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`           // 最后更新时间戳
}
