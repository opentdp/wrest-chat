package tables

// 群聊配置

type Chatroom struct {
	Rd           uint   `json:"rd" gorm:"primaryKey"`      // 主键
	Roomid       string `json:"roomid" gorm:"uniqueIndex"` // 群聊 id
	Name         string `json:"name"`                      // 群聊名称
	Level        int32  `json:"level" gorm:"default:-1"`   // 等级 [1:待验证, 2:已注册]
	Remark       string `json:"remark"`                    // 备注
	JoinArgot    string `json:"join_argot"`                // 加群指令
	PatReturn    string `json:"pat_return"`                // 响应拍拍我
	RevokeMsg    string `json:"revoke_msg"`                // 防撤回消息
	WelcomeMsg   string `json:"welcome_msg"`               // 欢迎消息
	ModelContext string `json:"model_context"`             // 定义模型扮演的身份
	ModelDefault string `json:"model_default"`             // 定义默认模型
	ModelHistory int    `json:"model_history"`             // 定义历史消息数量
	CreatedAt    int64  `json:"created_at"`                // 创建时间戳
	UpdatedAt    int64  `json:"updated_at"`                // 最后更新时间戳
}
