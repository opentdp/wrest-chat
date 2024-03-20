package main

import (
	"embed"

	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/dbase"
	"github.com/opentdp/wrest-chat/httpd"
	"github.com/opentdp/wrest-chat/wclient/crond"
	"github.com/opentdp/wrest-chat/wclient/plugin"
	"github.com/opentdp/wrest-chat/wclient/robot"
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
