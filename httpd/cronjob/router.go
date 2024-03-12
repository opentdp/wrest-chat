package cronjob

import (
	"github.com/opentdp/go-helper/httpd"

	"github.com/opentdp/wechat-rest/httpd/midware"
	"github.com/opentdp/wechat-rest/wclient/crond"
)

func Route() {

	rg := httpd.Group("/api/cronjob")
	rg.Use(midware.OutputHandle, midware.ApiGuard)

	cronjob := Cronjob{}
	rg.POST("list", cronjob.list)
	rg.POST("detail", cronjob.detail)
	rg.POST("create", cronjob.create)
	rg.POST("update", cronjob.update)
	rg.POST("delete", cronjob.delete)
	rg.POST("status", cronjob.status)

	crond.Daemon() // 启动定时任务

}
