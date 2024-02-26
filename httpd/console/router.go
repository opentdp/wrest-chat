package console

import (
	"github.com/opentdp/go-helper/httpd"

	"github.com/opentdp/wechat-rest/httpd/midware"
)

func Route() {

	rg := httpd.Group("/cpi")
	rg.Use(midware.OutputHandle, midware.ApiGuard)

}
