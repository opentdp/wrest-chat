package tables

// 群聊配置

type Chatroom struct {
	Rd         uint   `gorm:"primaryKey" json:"rd"`      // 主键
	Roomid     string `gorm:"uniqueIndex" json:"roomid"` // 群聊 id
	Name       string `json:"name"`                      // 群聊名称
	Level      int32  `json:"level"`                     // 等级
	Remark     string `json:"remark"`                    // 备注
	JoinArgot  string `json:"join_argot"`                // 入群口令
	PatReturn  string `json:"pat_return"`                // 响应拍拍我，"-"或空表示忽略
	RevokeMsg  string `json:"revoke_msg"`                // 防撤回消息
	WelcomeMsg string `json:"welcome_msg"`               // 欢迎消息
	CreatedAt  int64  `json:"created_at"`                // 创建时间戳
	UpdatedAt  int64  `json:"updated_at"`                // 最后更新时间戳
}
