package tables

// 全局配置

type Setting struct {
	Rd        uint   `gorm:"primaryKey" json:"rd"`    // 主键
	Name      string `gorm:"uniqueIndex" json:"name"` // 键
	Value     string `json:"value"`                   // 值
	Title     string `json:"title"`                   // 标题
	Remark    string `json:"remark"`                  // 备注
	CreatedAt int64  `json:"created_at"`              // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`              // 最后更新时间戳
}
