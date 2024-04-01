package tables

// Webhook

type Webhook struct {
	Rd        uint   `json:"rd" gorm:"primaryKey"`         // 主键
	TargetId  string `json:"target_id" gorm:"uniqueIndex"` // 群聊/私聊ID
	Token     string `json:"token" gorm:"uniqueIndex"`     // webhook 标识（GUID）
	Remark    string `json:"remark"`                       // 备注
	CreatedAt int64  `json:"created_at"`                   // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`                   // 最后更新时间戳
}
