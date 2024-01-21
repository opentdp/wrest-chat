package args

import (
	"os"

	"github.com/opentdp/go-helper/logman"
)

var Co = &Configer{}

// 机器人参数

var Bot = &IBot{
	Enable: true,
}

// 大语言模型

var LLM = &ILLM{
	HistoryNum: 20,
}

// 日志配置

var Log = &ILog{
	Dir:    "logs",
	Level:  "info",
	Target: "stdout",
}

// Web 服务

var Web = &IWeb{
	Address: "127.0.0.1:7600",
	Swagger: true,
}

// Wcf 服务

var Wcf = &IWcf{
	Address:    "127.0.0.1:7601",
	WeChatAuto: true,
}

// 初始化

func init() {

	println(AppName, AppSummary)
	println("Version:", Version, "build", BuildVersion)

	// 初始化配置

	debug := os.Getenv("TDP_DEBUG")
	Debug = debug == "1" || debug == "true"

	if err := Co.Init(); err != nil {
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
