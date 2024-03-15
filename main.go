package main

import (
	"embed"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase"
	"github.com/opentdp/wechat-rest/httpd"
	"github.com/opentdp/wechat-rest/wclient/crond"
	"github.com/opentdp/wechat-rest/wclient/plugin"
	"github.com/opentdp/wechat-rest/wclient/robot"
)

//go:embed public
var efs embed.FS

func main() {

	args.Efs = &efs

	dbase.Connect()

	crond.Daemon()
	plugin.CronjobPluginSetup()
	plugin.KeywordPluginSetup()

	robot.Start()

	httpd.Server()

}
