package args

import (
	"embed"
)

// 调试模式

var Debug bool

// 嵌入目录

var Efs *embed.FS

// 配置文件

var YamlFile = "./config.yml"

// 日志参数

type ILogger struct {
	Dir    string
	Level  string
	Target string
}

var Logger = ILogger{
	Dir:    "./logs",
	Level:  "info",
	Target: "stdout",
}

// Http 服务参数

type IHttpd struct {
	Address string
	Token   string
}

var Httpd = IHttpd{
	Address: "127.0.0.1:7600",
}

// Wcf 服务参数

type IWcf struct {
	Address    string
	SdkLibrary string
}

var Wcf = IWcf{
	Address:    "127.0.0.1:10080",
	SdkLibrary: "./sdk.dll",
}
