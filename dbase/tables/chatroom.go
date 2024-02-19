package tables

// 群聊

type Chatroom struct {
	Rd        uint   `gorm:"primaryKey"`
	Roomid    string `gorm:"uniqueIndex"` // 群聊 id
	Name      string // 群聊名称
	Level     int32  // 等级
	Remark    string // 备注
	JoinArgot string // 入群口令
	Welcome   string // 欢迎词
}
