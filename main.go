package main

import (
	"embed"

	"github.com/rehiy/wechat-rest-api/config"
	"github.com/rehiy/wechat-rest-api/server"
)

//go:embed public
var efs embed.FS

func main() {

	config.Efs = &efs
	config.Viper()

	server.Start()

}
