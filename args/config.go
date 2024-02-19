package args

import (
	"os"

	"github.com/opentdp/go-helper/logman"
)

// 机器人参数

var Bot = &IBot{
	Enable:        true,
	BadWord:       map[string]int{},
	InvitableRoom: []string{},
}

type IBot struct {
	BadWord       map[string]int // 命中该关键词时警告
	Enable        bool           // 是否启用内置机器人
	InvitableRoom []string       // 可邀请的群聊，必须在 Usr.ChatRoom 配置中
	Revoke        string         // 有人撤回消息时响应的内容，留空则不响应
	Welcome       string         // 接受好友申请时时响应的内容，留空则不响应
	WhiteChatRoom bool           // 白名单模式，仅允许 Level > 1 的群使用
	WhiteMember   bool           // 白名单模式，仅允许 Level > 1 的好友使用
}

// 大语言模型

var LLM = &ILLM{
	HistoryNum: 20,
	Models:     map[string]*LLModel{},
}

type ILLM struct {
	Default     string              // 默认模型
	HistoryNum  int                 // 历史消息数量
	RoleContext string              // 定义模型扮演的身份
	Models      map[string]*LLModel // 模型列表
}

type LLModel struct {
	Provider string // 服务商 [google, openai, xunfei]
	Endpoint string // 仅 google 和 openai 支持自定义，留空则使用官方接口
	Family   string // 模型家族，用于生成模型切换指令
	Model    string // 模型，必须和服务商提供的值对应
	Key      string // 密钥，google 和 openai 填写 KEY，xunfei 填写 APP-ID,API-KEY,API-SECRET
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

// 用户资料

var Usr = &IUsr{
	ChatRoom: map[string]*ChatRoom{},
	Member:   map[string]*Member{},
}

type IUsr struct {
	ChatRoom map[string]*ChatRoom // 群聊列表
	Member   map[string]*Member   // 用户列表
}

// Wcf 服务

var Wcf = &IWcf{
	Address:    "127.0.0.1:7601",
	WeChatAuto: true,
}

type IWcf struct {
	Address    string // Rpc 监听地址
	MsgBackup  bool   // 是否存储收到的消息
	MsgPrinter bool   // 是否打印收到的消息
	SdkLibrary string // 留空则禁止注入进程
	WeChatAuto bool   // 微信是否跟随启停
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
