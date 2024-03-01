package args

import (
	"os"

	"github.com/opentdp/go-helper/logman"
)

var Configer *Config

type ConfigData struct {
	Bot *IBot `yaml:"Bot"`
	LLM *ILLM `yaml:"LLM"`
	Log *ILog `yaml:"Log"`
	Wcf *IWcf `yaml:"Wcf"`
	Web *IWeb `yaml:"Web"`
}

func init() {

	println(AppName, AppSummary)
	println("Version:", Version, "build", BuildVersion)

	// 调试模式

	de := os.Getenv("TDP_DEBUG")
	Debug = de == "1" || de == "true"

	// 初始化配置

	Configer = &Config{
		File: "config.yml",
		Data: &ConfigData{Bot, LLM, Log, Wcf, Web},
	}

	if len(os.Args) > 1 {
		Configer.File = os.Args[1]
	}

	if err := Configer.Load(); err != nil {
		panic(err)
	}

	// 初始化日志

	logman.SetDefault(&logman.Config{
		Level:    Log.Level,
		Target:   Log.Target,
		Storage:  Log.Dir,
		Filename: "wrest",
	})

}
