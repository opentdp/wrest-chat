package tables

// 联系人

type Contact struct {
	Rd        uint   `gorm:"primaryKey"`  // 主键
	Wxid      string `gorm:"uniqueIndex"` // 微信 id
	Code      string // 微信号
	Remark    string // 备注
	Name      string // 微信昵称
	Country   string // 国家
	Province  string // 省/州
	City      string // 城市
	Gender    int32  // 性别
	CreatedAt int64  // 创建时间戳
	UpdatedAt int64  // 最后更新时间戳
}
