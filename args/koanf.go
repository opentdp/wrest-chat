package args

import (
	"os"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/opentdp/go-helper/logman"
)

// 配置信息操作类

type Config struct {
	Koanf  *koanf.Koanf
	Parser *yaml.YAML
	File   string
}

func (c *Config) Init() {

	debug := os.Getenv("TDP_DEBUG")
	Debug = debug == "1" || debug == "true"

	c.Koanf = koanf.NewWithConf(koanf.Conf{
		StrictMerge: true,
		Delim:       ".",
	})
	c.Parser = yaml.Parser()

	c.File = "config.yml"
	if len(os.Args) > 1 {
		c.File = os.Args[1]
	}

	c.Unmarshal()

}

func (c *Config) LoadYaml() error {

	// 配置不存在则忽略
	_, err := os.Stat(c.File)
	if os.IsNotExist(err) {
		return nil
	}

	logman.Warn("load config", "file", c.File)

	// 从配置文件读取参数
	err = c.Koanf.Load(file.Provider(c.File), c.Parser)
	if err != nil {
		logman.Error("load config", "error", err)
		return err
	}

	return nil

}

func (c *Config) WriteYaml() error {

	logman.Warn("write config", "file", c.File)

	// 序列化参数信息
	buf, err := c.Koanf.Marshal(c.Parser)
	if err != nil {
		logman.Error("write config", "error", err)
		return err
	}

	// 将参数写入配置文件
	err = os.WriteFile(c.File, buf, 0644)
	if err != nil {
		logman.Error("write config", "error", err)
		return err
	}

	return nil

}

func (c *Config) Unmarshal() {

	// 读取默认配置

	mp := map[string]any{
		"bot": &Bot,
		"llm": &LLM,
		"log": &Log,
		"web": &Web,
		"wcf": &Wcf,
	}
	c.Koanf.Load(confmap.Provider(mp, "."), nil)

	// 读取配置文件

	c.LoadYaml()
	for k, v := range mp {
		c.Koanf.Unmarshal(k, v)
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
