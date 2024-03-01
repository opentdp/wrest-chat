package args

// 机器人参数

type IBot struct {
	Enable       bool   `yaml:"Enable"`       // 是否启用内置机器人
	PatReturn    bool   `yaml:"PatReturn"`    // 是否自动回应拍一拍
	FriendAccept bool   `yaml:"FriendAccept"` // 是否自动同意新的好友请求
	RevokeMsg    string `yaml:"RevokeMsg"`    // 私聊撤回消息时响应的内容，留空则忽略
	FriendHello  string `yaml:"FriendHello"`  // 添加好友后的响应内容，留空则忽略
	WhiteLimit   bool   `yaml:"WhiteLimit"`   // 开启后只有白名单内的群或好友可以使用机器人
}

var Bot = &IBot{
	Enable: true,
}

// 大语言模型

type ILLM struct {
	Default     string `yaml:"Default"`     // 默认模型
	HistoryNum  int    `yaml:"HistoryNum"`  // 历史消息数量
	RoleContext string `yaml:"RoleContext"` // 定义模型扮演的身份
}

var LLM = &ILLM{
	HistoryNum: 20,
}

// 日志配置

type ILog struct {
	Dir    string `yaml:"Dir"`    // 日志目录
	Level  string `yaml:"Level"`  // 日志级别 debug|info|warn|error
	Target string `yaml:"Target"` // 日志输出方式 both|file|null|stdout|stderr
}

var Log = &ILog{
	Dir:    "logs",
	Level:  "info",
	Target: "stdout",
}

// Wcf 服务

type IWcf struct {
	Address    string `yaml:"Address"`    // Rpc 监听地址
	MsgStore   bool   `yaml:"MsgStore"`   // 是否存储收到的消息
	MsgPrint   bool   `yaml:"MsgPrint"`   // 是否打印收到的消息
	WcfBinary  string `yaml:"WcfBinary"`  // 留空则不注入微信
	WeChatAuto bool   `yaml:"WeChatAuto"` // 是否自动启停微信
}

var Wcf = &IWcf{
	Address:    "127.0.0.1:7601",
	WeChatAuto: true,
}

// Web 服务

type IWeb struct {
	Address string `yaml:"Address"` // Web 监听地址
	Swagger bool   `yaml:"Swagger"` // 是否启用 Api 文档
	Token   string `yaml:"Token"`   // 使用 Token 验证请求
}

var Web = &IWeb{
	Address: "127.0.0.1:7600",
	Swagger: true,
}
