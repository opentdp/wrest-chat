package tables

// 用户配置

type Profile struct {
	Rd        uint   `json:"rd" gorm:"primaryKey" `         // 主键
	Wxid      string `json:"wxid" gorm:"index"`             // 微信 id
	Roomid    string `json:"roomid" gorm:"index;default:-"` // 群聊 id
	Level     int32  `json:"level" gorm:"default:-1"`       // 等级
	Remark    string `json:"remark"`                        // 备注
	AiArgot   string `json:"ai_argot"`                      // 唤醒词
	AiModel   string `json:"ai_model"`                      // 会话模型
	BanExpire int64  `json:"ban_expire"`                    // 拉黑截止时间
	CreatedAt int64  `json:"created_at"`                    // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`                    // 最后更新时间戳
}
