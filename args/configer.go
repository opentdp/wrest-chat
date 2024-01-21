package args

import (
	"os"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
	"github.com/opentdp/go-helper/logman"
)

// 配置操作类

type Configer struct {
	Koanf  *koanf.Koanf
	Parser *yaml.YAML
	File   string
}

func (c *Configer) Init() error {

	c.File = "config.yml"
	if len(os.Args) > 1 {
		c.File = os.Args[1]
	}

	c.Koanf = koanf.NewWithConf(koanf.Conf{
		StrictMerge: true,
		Delim:       ".",
	})

	c.Parser = yaml.Parser()

	// 从文件加载
	return c.LoadYaml()

}

func (c *Configer) LoadYaml() error {

	logman.Info("load config", "file", c.File)

	// 文件不存在
	if _, err := os.Stat(c.File); os.IsNotExist(err) {
		logman.Warn("load config", "skip", c.File)
		return nil // 忽略错误
	}

	// 从文件读取参数
	err := c.Koanf.Load(file.Provider(c.File), c.Parser)
	if err == nil {
		c.Koanf.Unmarshal("bot", Bot)
		c.Koanf.Unmarshal("llm", LLM)
		c.Koanf.Unmarshal("log", Log)
		c.Koanf.Unmarshal("Web", Web)
		c.Koanf.Unmarshal("Wcf", Wcf)
		return nil
	}

	logman.Error("load config", "error", err)
	return err

}

func (c *Configer) SaveYaml() error {

	logman.Info("save config", "file", c.File)

	// 从内存读取参数
	tmp := &Config{
		Bot, LLM, Log, Web, Wcf,
	}
	err := c.Koanf.Load(structs.Provider(tmp, ""), nil)
	if err != nil {
		logman.Error("load struct", "error", err)
		return err
	}

	// 序列化参数信息
	buf, err := c.Koanf.Marshal(c.Parser)
	if err != nil {
		logman.Error("save config", "error", err)
		return err
	}

	// 将参数写入文件
	err = os.WriteFile(c.File, buf, 0644)
	if err != nil {
		logman.Error("save config", "error", err)
		return err
	}

	return nil

}
