package wcf

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/request"
	"github.com/opentdp/wechat-rest/config"
	"github.com/opentdp/wechat-rest/wcf-sdk"
)

var wc *wcf.Client

func initWcfService() {

	var err error

	wl := &wcf.Launcher{
		Address: config.Wcf.Address,
		Wcfexe:  config.Wcf.Executable,
	}

	if wc, err = wl.Start(); err != nil {
		logman.Fatal("failed to start wcf", "error", err)
	}

	wl.AutoDestory()

}

func isLogin(c *gin.Context) {

	c.Set("Payload", wc.IsLogin())

}

func getSelfWxid(c *gin.Context) {

	c.Set("Payload", wc.GetSelfWxid())

}

func getUserInfo(c *gin.Context) {

	c.Set("Payload", wc.GetUserInfo())

}

func getMsgTypes(c *gin.Context) {

	c.Set("Payload", wc.GetMsgTypes())

}

func getContacts(c *gin.Context) {

	c.Set("Payload", wc.GetContacts())

}

func getFriends(c *gin.Context) {

	c.Set("Payload", wc.GetFriends())

}

func getDbNames(c *gin.Context) {

	c.Set("Payload", wc.GetDbNames())

}

func getDbTables(c *gin.Context) {

	db := c.Param("db")
	c.Set("Payload", wc.GetDbTables(db))

}

func refreshPyq(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))
	c.Set("Payload", wc.RefreshPyq(id))

}

func getChatRooms(c *gin.Context) {

	c.Set("Payload", wc.GetChatRooms())

}

func getChatRoomMembers(c *gin.Context) {

	roomid := c.Param("roomid")
	c.Set("Payload", wc.GetChatRoomMembers(roomid))

}

func getAliasInChatRoom(c *gin.Context) {

	wxid := c.Param("wxid")
	roomid := c.Param("roomid")
	c.Set("Payload", wc.GetAliasInChatRoom(wxid, roomid))

}

func sendTxt(c *gin.Context) {

	var req wcf.TextMsg
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	status := wc.SendTxt(req.Msg, req.Receiver, req.Aters)

	c.Set("Payload", gin.H{
		"success": status == 0,
	})

}

func sendImg(c *gin.Context) {

	var req wcf.PathMsg
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	status := wc.SendImg(req.Path, req.Receiver)

	c.Set("Payload", gin.H{
		"success": status == 0,
	})

}

func sendFile(c *gin.Context) {

	var req wcf.PathMsg
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	status := wc.SendFile(req.Path, req.Receiver)

	c.Set("Payload", gin.H{
		"success": status == 0,
	})

}

type DbSqlQueryRequest struct {
	Db  string `json:"db"`
	Sql string `json:"sql"`
}

func dbSqlQuery(c *gin.Context) {

	var req DbSqlQueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	c.Set("Payload", wc.DbSqlQueryMap(req.Db, req.Sql))

}

func acceptNewFriend(c *gin.Context) {

	var req wcf.Verification
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	status := wc.AcceptNewFriend(req.V3, req.V4, req.Scene)

	c.Set("Payload", gin.H{
		"success": status == 1,
	})

}

func receiveTransfer(c *gin.Context) {

	var req wcf.Transfer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	status := wc.ReceiveTransfer(req.Wxid, req.Tfid, req.Taid)

	c.Set("Payload", gin.H{
		"success": status == 1,
	})

}

func addChatRoomMembers(c *gin.Context) {

	var req wcf.AddMembers
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	status := wc.AddChatRoomMembers(req.Roomid, req.Wxids)

	c.Set("Payload", gin.H{
		"success": status == 1,
	})

}

func delChatRoomMembers(c *gin.Context) {

	var req wcf.AddMembers
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	status := wc.DelChatRoomMembers(req.Roomid, req.Wxids)

	c.Set("Payload", gin.H{
		"success": status == 1,
	})

}

func decryptImage(c *gin.Context) {

	var req wcf.DecPath
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	status := wc.DecryptImage(req.Src, req.Dst)

	c.Set("Payload", gin.H{
		"success": status == 1,
	})

}

type ForwardMsgRequest struct {
	Url string `json:"url"`
}

func enableForwardMsg(c *gin.Context) {

	var req ForwardMsgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	cb := func(msg *wcf.WxMsg) {
		logman.Info("forward msg", "url", req.Url, "msg", msg)
		request.JsonPost(req.Url, msg, request.H{})
	}

	status := wc.ReceiverEnroll(true, cb)

	c.Set("Payload", gin.H{
		"success": status == 0,
	})

}

func disableForwardMsg(c *gin.Context) {

	status := wc.ReceiverDisable()

	c.Set("Payload", gin.H{
		"success": status == 0,
	})

}
