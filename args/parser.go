package args

import (
	"os"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"
)

// 配置信息操作类

type Config struct {
	Koanf  *koanf.Koanf
	Parser *yaml.YAML
}

func NewConfig() *Config {

	var p = yaml.Parser()
	var k = koanf.NewWithConf(koanf.Conf{
		StrictMerge: true,
		Delim:       ".",
	})

	return &Config{k, p}

}

func (c *Config) Init() {

	c.ReadYaml()

	// debug mode
	debug := os.Getenv("TDP_DEBUG")
	Debug = debug == "1" || debug == "true"

	// init logger
	logman.SetDefault(&logman.Config{
		Level:    Logger.Level,
		Target:   Logger.Target,
		Storage:  Logger.Dir,
		Filename: "global",
	})

	// write config
	c.WriteYaml(false)

}

func (c *Config) ReadYaml() {

	// 读取默认配置
	df := map[string]any{
		"logger": Logger,
		"httpd":  Httpd,
		"wcf":    Wcf,
	}
	c.Koanf.Load(confmap.Provider(df, "."), nil)

	// 不存在则忽略
	_, err := os.Stat(YamlFile)
	if os.IsNotExist(err) {
		return
	}

	// 读取配置文件
	err = c.Koanf.Load(file.Provider(YamlFile), c.Parser)
	if err != nil {
		logman.Fatal("read config error", "error", err)
	}

	// 解析配置信息
	c.Koanf.Unmarshal("logger", &Logger)
	c.Koanf.Unmarshal("httpd", &Httpd)
	c.Koanf.Unmarshal("wcf", &Wcf)

}

func (c *Config) WriteYaml(force bool) {

	// 是否强制覆盖
	if !force && filer.Exists(YamlFile) {
		return
	}

	// 序列化参数信息
	b, err := c.Koanf.Marshal(c.Parser)
	if err != nil {
		logman.Fatal("write config error", "error", err)
	}

	// 将参数写入配置文件
	err = os.WriteFile(YamlFile, b, 0644)
	if err != nil {
		logman.Fatal("write config error", "error", err)
	}

}
