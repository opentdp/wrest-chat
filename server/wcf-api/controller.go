package wcf

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/rehiy/wechat-rest-api/config"
	"github.com/rehiy/wechat-rest-api/wcf-sdk"
)

var wc *wcf.Client

func initWCF() {

	var err error

	wc, err = wcf.NewWCF(config.Wcf.Address)
	if err != nil {
		panic(err)
	}

}

func isLogin(c *gin.Context) {

	c.Set("Payload", wc.IsLogin())

}

func wxid(c *gin.Context) {

	c.Set("Payload", wc.GetSelfWxid())

}

func userInfo(c *gin.Context) {

	c.Set("Payload", wc.GetUserInfo())

}

func msgTypes(c *gin.Context) {

	c.Set("Payload", wc.GetMsgTypes())

}

func contacts(c *gin.Context) {

	c.Set("Payload", wc.GetContacts())

}

func friends(c *gin.Context) {

	notFriends := map[string]string{
		"mphelper":    "公众平台助手",
		"fmessage":    "朋友推荐消息",
		"medianote":   "语音记事本",
		"floatbottle": "漂流瓶",
		"filehelper":  "文件传输助手",
		"newsapp":     "新闻",
	}

	result := []*wcf.RpcContact{}
	for _, cnt := range wc.GetContacts() {
		if strings.HasSuffix(cnt.Wxid, "@chatroom") || strings.HasPrefix(cnt.Wxid, "gh_") || notFriends[cnt.Wxid] != "" {
			continue
		}
		result = append(result, cnt)
	}

	c.Set("Payload", result)

}

func dbNames(c *gin.Context) {

	c.Set("Payload", wc.GetDbNames())

}

func dbTables(c *gin.Context) {

	tab := c.Param("tab")
	c.Set("Payload", wc.GetDbTables(tab))

}

func pyqRefresh(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))
	c.Set("Payload", wc.RefreshPYQ(id))

}

func chatroomMembers(c *gin.Context) {

	roomid := c.Param("roomid")
	c.Set("Payload", wc.GetChatRoomMembers(roomid))

}
