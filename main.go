package main

import (
	"embed"

	"github.com/opentdp/wechat-rest/config"
	"github.com/opentdp/wechat-rest/server"
)

//go:embed public
var efs embed.FS

func main() {

	config.Efs = &efs
	config.Viper()

	server.Start()

}
