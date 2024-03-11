package tables

// 消息

type Message struct {
	Rd        uint   `json:"rd" gorm:"primaryKey"`    // 主键
	Id        uint64 `json:"id" gorm:"uniqueIndex"`   // 消息 id
	IsSelf    bool   `json:"is_self"`                 // 是否自己发送的
	IsGroup   bool   `json:"is_group"`                // 是否群消息
	Type      uint32 `json:"type"`                    // 消息类型
	Ts        uint32 `json:"ts"`                      // 消息类型
	Roomid    string `json:"roomid" gorm:"default:-"` // 群 id（如果是群消息的话）
	Content   string `json:"content"`                 // 消息内容
	Sender    string `json:"sender"`                  // 消息发送者
	Sign      string `json:"sign"`                    // Sign
	Thumb     string `json:"thumb"`                   // 缩略图
	Extra     string `json:"extra"`                   // 附加内容
	Xml       string `json:"xml"`                     // 消息 xml
	Remark    string `json:"remark"`                  // 备注，非微信字段
	CreatedAt int64  `json:"created_at"`              // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`              // 最后更新时间戳
}
