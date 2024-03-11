package main

import (
	"embed"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase"
	"github.com/opentdp/wechat-rest/httpd"
	"github.com/opentdp/wechat-rest/wclient/robot"
)

//go:embed public
var efs embed.FS

func main() {

	args.Efs = &efs

	dbase.Connect()

	robot.Start()
	httpd.Server()

}
