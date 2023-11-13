package config

import (
	"log"
	"os"

	"github.com/opentdp/go-helper/logman"
	"github.com/spf13/viper"
)

func init() {

	ViperFile = "./config.yml"

	viper.SetDefault("logger.dir", ".")
	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.target", "stdout")

	viper.SetDefault("httpd.address", "127.0.0.1:7600")

	viper.SetDefault("wcf.address", "127.0.0.1:10080")

}

func Viper() {

	viper.SetConfigFile(ViperFile)

	// 读取配置

	viper.SetEnvPrefix("WXAPI")
	viper.AutomaticEnv()

	if _, err := os.Stat(ViperFile); err == nil {
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
	}

	// 获取参数

	Debug = viper.GetBool("debug")

	Logger.Dir = viper.GetString("logger.dir")
	Logger.Level = viper.GetString("logger.level")
	Logger.Target = viper.GetString("logger.target")

	Httpd.Address = viper.GetString("httpd.address")
	Httpd.Token = viper.GetString("httpd.token")

	Wcf.Address = viper.GetString("wcf.address")
	Wcf.Executable = viper.GetString("wcf.executable")

	// 写入配置文件

	if err := viper.WriteConfig(); err != nil {
		logman.Fatal("write config error", err)
	}

}
