package tables

// 大语言模型

type LLModel struct {
	Rd        uint   `gorm:"primaryKey" json:"rd"`   // 主键
	Mid       string `gorm:"uniqueIndex" json:"mid"` // 模型 Id
	Provider  string `json:"provider"`               // 服务商 [google, openai, xunfei]
	Endpoint  string `json:"endpoint"`               // 仅 google 和 openai 支持自定义，留空则使用官方接口
	Family    string `json:"family"`                 // 模型家族，用于生成模型切换指令
	Model     string `json:"model"`                  // 模型，必须和服务商提供的值对应
	Secret    string `json:"secret"`                 // 密钥，google 和 openai 填写 KEY，xunfei 填写 APP-ID,API-KEY,API-SECRET
	CreatedAt int64  `json:"created_at"`             // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`             // 最后更新时间戳
}
