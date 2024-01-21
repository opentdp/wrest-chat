package args

type Config struct {
	Bot *IBot `koanf:"bot"`
	LLM *ILLM `koanf:"llm"`
	Log *ILog `koanf:"log"`
	Web *IWeb `koanf:"web"`
	Wcf *IWcf `koanf:"wcf"`
}

// 机器人参数

type IBot struct {
	Enable      bool       `koanf:"enable"`
	Welcome     string     `koanf:"welcome"`
	Revoke      string     `koanf:"revoke"`
	Managers    []string   `koanf:"managers"`
	BlackList   []string   `koanf:"blackList"`
	WhiteList   []string   `koanf:"whiteList"`
	HostedRooms []*BotRoom `koanf:"hostedRooms"`
}

type BotRoom struct {
	Mask    string `koanf:"mask"`
	Name    string `koanf:"name"`
	RoomId  string `koanf:"roomId"`
	Welcome string `koanf:"welcome"`
}

// 大语言模型

type ILLM struct {
	HistoryNum  int    `koanf:"historyNum"`
	RoleContext string `koanf:"roleContext"`
	Models      []*LLModel
}

type LLModel struct {
	Name     string `koanf:"name"`
	Provider string `koanf:"provider"`
	Endpoint string `koanf:"endpoint"`
	Model    string `koanf:"model"`
	Key      string `koanf:"key"`
}

// 日志配置

type ILog struct {
	Dir    string `koanf:"dir"`
	Level  string `koanf:"level"`
	Target string `koanf:"target"`
}

// Web 服务

type IWeb struct {
	Address string `koanf:"address"`
	Swagger bool   `koanf:"swagger"`
	Token   string `koanf:"token"`
}

// Wcf 服务

type IWcf struct {
	Address    string `koanf:"address"`
	WeChatAuto bool   `koanf:"wechatAuto"`
	MsgPrinter bool   `koanf:"msgPrinter"`
}
