package args

import (
	"os"

	"github.com/opentdp/go-helper/logman"
)

// 机器人参数

var Bot = &IBot{
	Enable: true,
}

type IBot struct {
	Enable       bool   // 是否启用内置机器人
	PatReturn    bool   // 是否自动回应拍一拍
	FriendAccept bool   // 是否自动同意新的好友请求
	RevokeMsg    string // 私聊撤回消息时响应的内容，留空则忽略
	FriendHello  string // 添加好友后的响应内容，留空则忽略
	WhiteLimit   bool   // 开启后只有白名单内的群或好友可以使用机器人
}

// 大语言模型

var LLM = &ILLM{
	HistoryNum: 20,
}

type ILLM struct {
	Default     string // 默认模型
	HistoryNum  int    // 历史消息数量
	RoleContext string // 定义模型扮演的身份
}

// 日志配置

var Log = &ILog{
	Dir:    "logs",
	Level:  "info",
	Target: "stdout",
}

type ILog struct {
	Dir    string // 日志目录
	Level  string // 日志级别 DEBUG|INFO|WARN|ERROR
	Target string // 日志输出方式 both|file|null|stdout|stderr
}

// Wcf 服务

var Wcf = &IWcf{
	Address:    "127.0.0.1:7601",
	WeChatAuto: true,
}

type IWcf struct {
	Address    string // Rpc 监听地址
	MsgStore   bool   // 是否存储收到的消息
	MsgPrint   bool   // 是否打印收到的消息
	WcfBinary  string // 留空则不注入微信
	WeChatAuto bool   // 是否自动启停微信
}

// Web 服务

var Web = &IWeb{
	Address: "127.0.0.1:7600",
	Swagger: true,
}

type IWeb struct {
	Address string // Web 监听地址
	Swagger bool   // 是否启用 Api 文档
	Token   string // 使用 Token 验证请求
}

// 初始化

func init() {

	println(AppName, AppSummary)
	println("Version:", Version, "build", BuildVersion)

	// 初始化配置

	debug := os.Getenv("TDP_DEBUG")
	Debug = debug == "1" || debug == "true"

	if err := LoadConfig(); err != nil {
		panic(err)
	}

	// 初始化日志

	if Log.Dir != "" && Log.Dir != "." {
		os.MkdirAll(Log.Dir, 0755)
	}

	logman.SetDefault(&logman.Config{
		Level:    Log.Level,
		Target:   Log.Target,
		Storage:  Log.Dir,
		Filename: "wrest",
	})

}
