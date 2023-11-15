package server

import (
	"github.com/opentdp/go-helper/httpd"

	"github.com/opentdp/wechat-rest/config"
	"github.com/opentdp/wechat-rest/server/midware"
	"github.com/opentdp/wechat-rest/server/wcf-api"
)

func Start() {

	httpd.Engine(config.Debug)

	api := httpd.Group("/api")
	api.Use(midware.OutputHandle)
	api.Use(midware.AuthGuard)

	// 注册 WCF
	wcf.Route(api)

	// 前端文件路由
	httpd.StaticEmbed("/", "public", config.Efs)

	// 启动 HTTP 服务
	httpd.Server(config.Httpd.Address)

}
