package wcf

import (
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

type sendTxtReqeust struct {
	Msg      string   `json:"msg"`
	Receiver string   `json:"receiver"`
	Aters    []string `json:"aters"`
}

func sendTxt(c *gin.Context) {

	var req sendTxtReqeust
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	c.Set("Payload", wc.SendTxt(req.Msg, req.Receiver, req.Aters))

}

type SendImgRequest struct {
	Path     string `json:"path"`
	Receiver string `json:"receiver"`
}

func sendImg(c *gin.Context) {

	var req SendImgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	c.Set("Payload", wc.SendImg(req.Path, req.Receiver))

}

type SendFileRequest struct {
	Path     string `json:"path"`
	Receiver string `json:"receiver"`
}

func sendFile(c *gin.Context) {

	var req SendFileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Payload", err)
		return
	}

	c.Set("Payload", wc.SendFile(req.Path, req.Receiver))

}
