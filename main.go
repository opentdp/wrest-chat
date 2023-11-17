package main

import (
	"embed"

	"github.com/opentdp/wechat-rest/config"
	"github.com/opentdp/wechat-rest/server"
)

//go:embed public
var efs embed.FS

// @title Wechat Rest API
// @version v0.0.1
// @description 基于 WeChatFerry RPC 实现的电脑版微信 REST-API，使用 Go 语言编写，无第三方运行时依赖。基于 HTTP 提供操作接口，轻松对接任意编程语言。
// @contact.name OpenTDP
// @contact.url https://github.com/opentdp/wechat-rest
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api

func main() {

	config.Efs = &efs
	config.Viper()

	server.Start()

}
