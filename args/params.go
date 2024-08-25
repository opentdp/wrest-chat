package args

// 日志配置

type ILog struct {
	Dir    string `yaml:"Dir"`    // 存储目录
	Level  string `yaml:"Level"`  // 记录级别 debug|info|warn|error
	Target string `yaml:"Target"` // 输出方式 both|file|null|stdout|stderr
}

var Log = &ILog{
	Dir:    "logs",
	Level:  "info",
	Target: "stdout",
}

// Wcf 服务

type IWcf struct {
	Address      string `yaml:"Address"`      // Rpc 监听地址
	MsgPrint     bool   `yaml:"MsgPrint"`     // 是否打印收到的消息
	MsgStore     bool   `yaml:"MsgStore"`     // 是否存储收到的消息
	MsgStoreDays int    `yaml:"MsgStoreDays"` // 消息留存天数
	SdkLibrary   string `yaml:"SdkLibrary"`   // 留空则不注入微信
}

var Wcf = &IWcf{
	Address: "127.0.0.1:7601",
}

// Web 服务

type IWeb struct {
	Address string `yaml:"Address"` // Web 监听地址
	PushUrl string `yaml:"PushUrl"` // 消息推送地址
	Storage string `yaml:"Storage"` // 附件存储路径
	Swagger bool   `yaml:"Swagger"` // 是否启用 Api 文档
	Token   string `yaml:"Token"`   // 使用 Token 验证请求
}

var Web = &IWeb{
	Address: "127.0.0.1:7600",
	Storage: "storage",
	Swagger: true,
}
