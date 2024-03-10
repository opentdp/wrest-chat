package tables

// 大语言模型

type LLModel struct {
	Rd        uint   `json:"rd" gorm:"primaryKey"`    // 主键
	Mid       string `json:"mid" gorm:"uniqueIndex"`  // 模型 Id，用于生成模型切换指令
	Level     int32  `json:"level" gorm:"default:-1"` // 等级，用于限制用户访问
	Family    string `json:"family"`                  // 模型家族，用于在指令中描述模型类型
	Provider  string `json:"provider"`                // 服务商 [google, openai, xunfei, baidu, tencent]
	Model     string `json:"model"`                   // 模型，必须和服务商对应
	Secret    string `json:"secret"`                  // 密钥，必须和服务商对应，多个字段按指定顺序填写并用逗号隔开
	Endpoint  string `json:"endpoint"`                // 仅 google 和 openai 支持自定义
	CreatedAt int64  `json:"created_at"`              // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`              // 最后更新时间戳
}
