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
	Enable      bool
	Welcome     string
	Revoke      string
	Managers    []string
	BlackList   []string
	WhiteList   []string
	HostedRooms []*BotRoom
}

type BotRoom struct {
	Mask    string
	Name    string
	RoomId  string
	Welcome string
}

// 大语言模型

var LLM = &ILLM{
	HistoryNum: 20,
}

type ILLM struct {
	HistoryNum  int
	RoleContext string
	Models      []*LLModel
}

type LLModel struct {
	Name     string
	Provider string
	Endpoint string
	Model    string
	Key      string
}

// 日志配置

var Log = &ILog{
	Dir:    "logs",
	Level:  "info",
	Target: "stdout",
}

type ILog struct {
	Dir    string
	Level  string
	Target string
}

// Wcf 服务

var Wcf = &IWcf{
	Address:    "127.0.0.1:7601",
	WeChatAuto: true,
}

type IWcf struct {
	Address    string
	WeChatAuto bool
	MsgPrinter bool
}

// Web 服务

var Web = &IWeb{
	Address: "127.0.0.1:7600",
	Swagger: true,
}

type IWeb struct {
	Address string
	Swagger bool
	Token   string
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
