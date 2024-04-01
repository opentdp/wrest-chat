package tables

// 联系人

type Contact struct {
	Rd        uint   `json:"rd" gorm:"primaryKey"`    // 主键
	Wxid      string `json:"wxid" gorm:"uniqueIndex"` // 微信 id
	Code      string `json:"code"`                    // 微信号
	Remark    string `json:"remark"`                  // 备注
	Name      string `json:"name"`                    // 微信昵称
	Country   string `json:"country"`                 // 国家
	Province  string `json:"province"`                // 省/州
	City      string `json:"city"`                    // 城市
	Gender    int32  `json:"gender"`                  // 性别
	CreatedAt int64  `json:"created_at"`              // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`              // 最后更新时间戳
}
