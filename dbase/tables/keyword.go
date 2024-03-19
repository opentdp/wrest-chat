package tables

// 关键词

type Keyword struct {
	Rd        uint   `json:"rd" gorm:"primaryKey"`                                  // 主键
	Group     string `json:"group" gorm:"uniqueIndex:idx_keyword_g_r_p"`            // 分组
	Roomid    string `json:"roomid" gorm:"uniqueIndex:idx_keyword_g_r_p;default:-"` // 群聊 id
	Phrase    string `json:"phrase" gorm:"uniqueIndex:idx_keyword_g_r_p"`           // 短语
	Level     int32  `json:"level" gorm:"default:-1"`                               // 等级
	Target    string `json:"target"`                                                // 目标
	Remark    string `json:"remark"`                                                // 备注
	CreatedAt int64  `json:"created_at"`                                            // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`                                            // 最后更新时间戳
}
