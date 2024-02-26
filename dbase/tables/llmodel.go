package tables

// 大语言模型

type LLModel struct {
	Rd        uint   `gorm:"primaryKey"`  // 主键
	Mid       string `gorm:"uniqueIndex"` // 模型 Id
	Provider  string // 服务商 [google, openai, xunfei]
	Endpoint  string // 仅 google 和 openai 支持自定义，留空则使用官方接口
	Family    string // 模型家族，用于生成模型切换指令
	Model     string // 模型，必须和服务商提供的值对应
	Secret    string // 密钥，google 和 openai 填写 KEY，xunfei 填写 APP-ID,API-KEY,API-SECRET
	CreatedAt int64  // 创建时间戳
	UpdatedAt int64  // 最后更新时间戳
}
