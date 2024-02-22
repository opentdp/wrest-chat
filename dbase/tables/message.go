package tables

// 消息

type Message struct {
	Rd        uint   `gorm:"primaryKey"`  // 主键
	Id        uint64 `gorm:"uniqueIndex"` // 消息 id
	IsSelf    bool   // 是否自己发送的
	IsGroup   bool   // 是否群消息
	Type      uint32 // 消息类型
	Ts        uint32 // 消息类型
	Roomid    string // 群 id（如果是群消息的话）
	Content   string // 消息内容
	Sender    string // 消息发送者
	Sign      string // Sign
	Thumb     string // 缩略图
	Extra     string // 附加内容
	Xml       string // 消息 xml
	CreatedAt int64  // 创建时间戳
	UpdatedAt int64  // 最后更新时间戳
}
