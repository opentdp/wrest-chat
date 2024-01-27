package model

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

// 联系人

type Contact struct {
	Rd        uint   `gorm:"primaryKey"`
	Wxid      string `gorm:"uniqueIndex"` // 微信 id
	Code      string // 微信号
	Remark    string // 备注
	Name      string // 微信昵称
	Country   string // 国家
	Province  string // 省/州
	City      string // 城市
	Gender    int32  // 性别
	CreatedAt int64
	UpdatedAt int64
}

// 消息

type Message struct {
	Rd        uint   `gorm:"primaryKey"`
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
	CreatedAt int64
	UpdatedAt int64
}

// 配置

type Profile struct {
	Rd        uint   `gorm:"primaryKey"`
	Wxid      string `gorm:"uniqueIndex"` // 微信 id
	Roomid    string `gorm:"uniqueIndex"` // 群聊 id
	Level     int32  // 等级
	AiArgot   string // 唤醒词
	AiModel   string // 会话模型
	CreatedAt int64
	UpdatedAt int64
}
