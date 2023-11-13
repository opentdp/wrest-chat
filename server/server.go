package server

import (
	"github.com/opentdp/go-helper/httpd"

	"github.com/rehiy/wechat-rest-api/config"
	"github.com/rehiy/wechat-rest-api/server/midware"
	"github.com/rehiy/wechat-rest-api/server/wcf-api"
)

func Start() {

	engine := httpd.Engine(config.Debug)

	api := engine.Group("/api")
	api.Use(midware.OutputHandle)
	api.Use(midware.AuthGuard)

	// 注册 WCF
	wcf.Route(api)

	// 前端文件路由
	httpd.StaticEmbed("/", "public", config.Efs)

	// 启动 HTTP 服务
	httpd.Server(config.Httpd.Address)

}
