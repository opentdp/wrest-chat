package args

import (
	"os"

	"github.com/knadh/koanf/providers/confmap"
	"github.com/opentdp/go-helper/logman"
)

func (c *Config) Unmarshal() {

	// 读取默认配置

	df := map[string]any{
		"logger": Logger,
		"httpd":  Httpd,
		"wcf":    Wcf,
	}
	c.Koanf.Load(confmap.Provider(df, "."), nil)

	// 读取配置文件

	c.ReadYaml()
	c.Koanf.Unmarshal("logger", &Logger)
	c.Koanf.Unmarshal("httpd", &Httpd)
	c.Koanf.Unmarshal("wcf", &Wcf)

	// 初始化日志

	if Logger.Dir != "" && Logger.Dir != "." {
		os.MkdirAll(Logger.Dir, 0755)
	}

	logman.SetDefault(&logman.Config{
		Level:    Logger.Level,
		Target:   Logger.Target,
		Storage:  Logger.Dir,
		Filename: "rest",
	})

	// 写入配置文件

	c.WriteYaml()

}
