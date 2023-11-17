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
// @description 微信 REST 接口文档

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api

func main() {

	config.Efs = &efs
	config.Viper()

	server.Start()

}
