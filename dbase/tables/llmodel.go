package tables

// 大语言模型

type LLModel struct {
	Rd        uint   `gorm:"primaryKey" json:"rd"`   // 主键
	Mid       string `gorm:"uniqueIndex" json:"mid"` // 模型 Id，用于生成模型切换指令
	Family    string `json:"family"`                 // 模型家族，用于在指令中描述模型类型
	Provider  string `json:"provider"`               // 服务商 [google, openai, xunfei]
	Model     string `json:"model"`                  // 模型，必须和服务商对应
	Secret    string `json:"secret"`                 // 密钥，google 和 openai 填写 KEY，xunfei 填写 APP-ID,API-KEY,API-SECRET
	Endpoint  string `json:"endpoint"`               // 仅 google 和 openai 支持自定义，留空则使用官方接口
	CreatedAt int64  `json:"created_at"`             // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`             // 最后更新时间戳
}
