package tables

// 关键词

type Keyword struct {
	Rd        uint   `gorm:"primaryKey" json:"rd"`               // 主键
	Roomid    string `gorm:"uniqueIndex:r_phrase" json:"roomid"` // 群聊 id
	Phrase    string `gorm:"uniqueIndex:r_phrase" json:"phrase"` // 短语
	Target    string `gorm:"index;default:-" json:"target"`      // 目标
	Level     int32  `gorm:"default:-1" json:"level"`            // 等级
	CreatedAt int64  `json:"created_at"`                         // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`                         // 最后更新时间戳
}
