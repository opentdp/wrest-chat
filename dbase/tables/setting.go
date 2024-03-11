package tables

// 全局配置

type Setting struct {
	Rd        uint   `json:"rd" gorm:"primaryKey"`       // 主键
	Name      string `json:"name" gorm:"uniqueIndex"`    // 键
	Type      string `json:"type" gorm:"default:string"` // 类型
	Group     string `json:"group" gorm:"index"`         // 分组
	Value     string `json:"value"`                      // 值
	Title     string `json:"title"`                      // 标题
	Remark    string `json:"remark"`                     // 备注
	CreatedAt int64  `json:"created_at"`                 // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`                 // 最后更新时间戳
}
