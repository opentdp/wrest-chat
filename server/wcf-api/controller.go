package wcf

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/logman"
	"github.com/spf13/cast"

	"github.com/opentdp/wechat-rest/config"
	"github.com/opentdp/wechat-rest/wcf-sdk"
)

var wc *wcf.Client

func initWCF() {

	var err error

	wl := &wcf.Launcher{
		Address: config.Wcf.Address,
		Wcfexe:  config.Wcf.Executable,
	}

	if wc, err = wl.Start(); err != nil {
		logman.Fatal("start wcf faild", "error", err)
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

	c.Set("Payload", wc.SendTxt(req.Msg, req.Receiver, req.Aters))

}

func sendImg(c *gin.Context) {

	var req wcf.PathMsg
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	c.Set("Payload", wc.SendImg(req.Path, req.Receiver))

}

func sendFile(c *gin.Context) {

	var req wcf.PathMsg
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	c.Set("Payload", wc.SendFile(req.Path, req.Receiver))

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

	c.Set("Payload", wc.AcceptNewFriend(req.V3, req.V4, req.Scene))

}

func receiveTransfer(c *gin.Context) {

	var req wcf.Transfer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	c.Set("Payload", wc.ReceiveTransfer(req.Wxid, req.Tfid, req.Taid))

}

func addChatRoomMembers(c *gin.Context) {

	var req wcf.AddMembers
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	c.Set("Payload", wc.AddChatRoomMembers(req.Roomid, req.Wxids))

}

func delChatRoomMembers(c *gin.Context) {

	var req wcf.AddMembers
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	c.Set("Payload", wc.DelChatRoomMembers(req.Roomid, req.Wxids))

}

func decryptImage(c *gin.Context) {

	var req wcf.DecPath
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	c.Set("Payload", wc.DecryptImage(req.Src, req.Dst))

}

func enableReceivingMsg(c *gin.Context) {

	status := wc.EnableReceivingMsg(true)

	go wc.OnReceivingMsg(func(msg *wcf.WxMsg) {
		logman.Info("OnReceivingMsg", "msg", msg)
	})

	c.Set("Payload", gin.H{
		"status": status,
	})

}

func disableReceivingMsg(c *gin.Context) {

	status := wc.DisableReceivingMsg()

	c.Set("Payload", gin.H{
		"status": status,
	})

}
