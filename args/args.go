package args

import (
	"embed"
)

// 调试模式

var Debug bool

// 嵌入目录

var Efs *embed.FS

// 配置文件

var ViperFile string

// 日志参数

var Logger struct {
	Dir    string
	Level  string
	Target string
}

// Http 服务参数

var Httpd struct {
	Address string
	Token   string
}

// Wcf 服务参数

var Wcf struct {
	Address    string
	Executable string
}
