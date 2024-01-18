package args

import (
	"embed"
)

// 调试模式

var Debug bool

// 嵌入目录

var Efs *embed.FS

// 机器人参数

var Bot = struct {
	Enable      bool       `yaml:"enable"`
	Welcome     string     `yaml:"welcome"`
	Revoke      string     `yaml:"revoke"`
	Managers    []string   `yaml:"managers"`
	BlackList   []string   `yaml:"blackList"`
	WhiteList   []string   `yaml:"whiteList"`
	HostedRooms []*BotRoom `yaml:"hostedRooms"`
}{
	Enable: true,
}

type BotRoom struct {
	Mask    string `yaml:"mask"`
	Name    string `yaml:"name"`
	RoomId  string `yaml:"roomId"`
	Welcome string `yaml:"welcome"`
}

// 大语言模型

var LLM = struct {
	HistoryNum  int    `yaml:"historyNum"`
	RoleContext string `yaml:"roleContext"`
	Models      []*LLModel
}{
	HistoryNum: 20,
}

type LLModel struct {
	Name     string `yaml:"name"`
	Provider string `yaml:"provider"`
	Endpoint string `yaml:"endpoint"`
	Model    string `yaml:"model"`
	Key      string `yaml:"key"`
}

// 日志配置

var Log = struct {
	Dir    string `yaml:"dir"`
	Level  string `yaml:"level"`
	Target string `yaml:"target"`
}{
	Dir:    "logs",
	Level:  "info",
	Target: "stdout",
}

// Web 服务

var Web = struct {
	Address string `yaml:"address"`
	Swagger bool   `yaml:"swagger"`
	Token   string `yaml:"token"`
}{
	Address: "127.0.0.1:7600",
	Swagger: true,
}

// Wcf 服务

var Wcf = struct {
	Address    string `yaml:"address"`
	WeChatAuto bool   `yaml:"wechatAuto"`
	MsgPrinter bool   `yaml:"msgPrinter"`
}{
	Address:    "127.0.0.1:7601",
	WeChatAuto: true,
}
