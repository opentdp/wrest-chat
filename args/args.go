package args

import (
	"embed"
)

// 调试模式

var Debug bool

// 嵌入目录

var Efs *embed.FS

// 机器人参数

type BotRoom struct {
	Id      string `yaml:"id"`
	Name    string `yaml:"name"`
	Welcome string `yaml:"welcome"`
}

var Bot = struct {
	Enable         bool       `yaml:"enable"`
	Welcome        string     `yaml:"welcome"`
	InvitableRooms []*BotRoom `yaml:"invitableRooms"`
}{
	Enable: true,
}

// Http 服务参数

var Httpd = struct {
	Address string `yaml:"address"`
	Token   string `yaml:"token"`
	Swag    bool   `yaml:"swag"`
}{
	Address: "127.0.0.1:7600",
	Swag:    true,
}

// 大语言模型参数

type LLMModel struct {
	Name     string `yaml:"name"`
	Provider string `yaml:"provider"`
	Endpoint string `yaml:"endpoint"`
	Model    string `yaml:"model"`
	Key      string `yaml:"key"`
}

var LLM = struct {
	HistoryNum int `yaml:"historyNum"`
	Models     []*LLMModel
}{
	HistoryNum: 20,
}

// 日志参数

var Logger = struct {
	Dir    string `yaml:"dir"`
	Level  string `yaml:"level"`
	Target string `yaml:"target"`
}{
	Dir:    "logs",
	Level:  "info",
	Target: "stdout",
}

// Wcf 服务参数

var Wcf = struct {
	Address    string `yaml:"address"`
	SdkLibrary string `yaml:"sdkLibrary"`
	WeChatAuto bool   `yaml:"wechatAuto"`
	MsgPrinter bool   `yaml:"msgPrinter"`
}{
	Address:    "127.0.0.1:10080",
	SdkLibrary: "sdk.dll",
}
