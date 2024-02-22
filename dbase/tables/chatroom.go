package tables

// 群聊

type Chatroom struct {
	Rd        uint   `gorm:"primaryKey"`  // 主键
	Roomid    string `gorm:"uniqueIndex"` // 群聊 id
	Name      string // 群聊名称
	Level     int32  // 等级
	Remark    string // 备注
	JoinArgot string // 入群口令
	Welcome   string // 欢迎词
	CreatedAt int64  // 创建时间戳
	UpdatedAt int64  // 最后更新时间戳
}
