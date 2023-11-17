package wcf

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/request"
	"github.com/opentdp/wechat-rest/config"
	"github.com/opentdp/wechat-rest/wcf-sdk"
)

var wc *wcf.Client

func initService() {

	parts := strings.Split(config.Wcf.Address, ":")
	port := cast.ToInt(parts[1])

	wc = &wcf.Client{
		WcfPath: config.Wcf.Executable,
		WcfAddr: parts[0],
		WcfPort: port,
	}

	if err := wc.Connect(); err != nil {
		logman.Fatal("failed to start wcf", "error", err)
	}

	wc.AutoDestory()

}

// @Summary 检查登录状态
// @Produce json
// @Success 200 {object} bool
// @Router /is_login [get]
func isLogin(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.IsLogin())

}

// @Summary 获取登录账号wxid
// @Produce json
// @Success 200 {object} string
// @Router /self_wxid [get]
func getSelfWxid(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetSelfWxid())

}

// @Summary 获取登录账号个人信息
// @Produce json
// @Success 200 {object} wcf.UserInfo
// @Router /user_info [get]
func getUserInfo(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetUserInfo())

}

// @Summary 获取所有消息类型
// @Produce json
// @Success 200 {object} map[int32]string
// @Router /msg_types [get]
func getMsgTypes(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetMsgTypes())

}

// @Summary 获取完整通讯录
// @Produce json
// @Success 200 {object} []wcf.RpcContact
// @Router /contacts [get]
func getContacts(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetContacts())

}

// @Summary 获取好友列表
// @Produce json
// @Success 200 {object} []wcf.RpcContact
// @Router /friends [get]
func getFriends(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetFriends())

}

// @Summary 获取数据库列表
// @Produce json
// @Success 200 {object} []string
// @Router /db_names [get]
func getDbNames(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetDbNames())

}

// @Summary 获取数据库表列表
// @Produce json
// @Param db path string true "数据库名"
// @Success 200 {object} []wcf.DbTable
// @Router /db_tables/{db} [get]
func getDbTables(c *gin.Context) {

	db := c.Param("db")
	c.Set("Payload", wc.CmdClient.GetDbTables(db))

}

// @Summary 刷新朋友圈
// @Produce json
// @Param id path int true "朋友圈id"
// @Success 200 {object} ActionStatus
// @Router /refresh_pyq/{id} [get]
func refreshPyq(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))
	status := wc.CmdClient.RefreshPyq(id)

	c.Set("Payload", ActionStatus{
		Success: status == 1,
	})

}

// @Summary 获取群列表
// @Produce json
// @Success 200 {object} []wcf.RpcContact
// @Router /chatrooms [get]
func getChatRooms(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetChatRooms())

}

// @Summary 获取群成员列表
// @Produce json
// @Param roomid path string true "群id"
// @Success 200 {object} []wcf.RpcContact
// @Router /chatroom_members/{roomid} [get]
func getChatRoomMembers(c *gin.Context) {

	roomid := c.Param("roomid")
	c.Set("Payload", wc.CmdClient.GetChatRoomMembers(roomid))

}

// @Summary 获取群成员昵称
// @Produce json
// @Param wxid path string true "wxid"
// @Param roomid path string true "群id"
// @Success 200 {object} string
// @Router /alias_in_chatroom/{wxid}/{roomid} [get]
func getAliasInChatRoom(c *gin.Context) {

	wxid := c.Param("wxid")
	roomid := c.Param("roomid")
	c.Set("Payload", wc.CmdClient.GetAliasInChatRoom(wxid, roomid))

}

// @Summary 发送文本消息
// @Produce json
// @Param body body wcf.TextMsg true "文本消息请求参数"
// @Success 200 {object} ActionStatus
// @Router /send_txt [post]
func sendTxt(c *gin.Context) {

	var req wcf.TextMsg
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.SendTxt(req.Msg, req.Receiver, req.Aters)

	c.Set("Payload", ActionStatus{
		Success: status == 0,
	})

}

// @Summary 发送图片消息
// @Produce json
// @Param body body wcf.PathMsg true "图片消息请求参数"
// @Success 200 {object} ActionStatus
// @Router /send_img [post]
func sendImg(c *gin.Context) {

	var req wcf.PathMsg
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.SendImg(req.Path, req.Receiver)

	c.Set("Payload", ActionStatus{
		Success: status == 0,
	})

}

// @Summary 发送文件消息
// @Produce json
// @Param body body wcf.PathMsg true "文件消息请求参数"
// @Success 200 {object} ActionStatus
// @Router /send_file [post]
func sendFile(c *gin.Context) {

	var req wcf.PathMsg
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.SendFile(req.Path, req.Receiver)

	c.Set("Payload", ActionStatus{
		Success: status == 0,
	})

}

// @Summary 执行数据库查询
// @Produce json
// @Param body body DbSqlQueryRequest true "数据库查询请求参数"
// @Success 200 {object} ActionStatus
// @Router /db_query_sql [post]
func dbSqlQuery(c *gin.Context) {

	var req DbSqlQueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", wc.CmdClient.DbSqlQueryMap(req.Db, req.Sql))

}

// @Summary 接受好友请求
// @Produce json
// @Param body body wcf.Verification true "接受好友请求参数"
// @Success 200 {object} ActionStatus
// @Router /accept_new_friend [post]
func acceptNewFriend(c *gin.Context) {

	var req wcf.Verification
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.AcceptNewFriend(req.V3, req.V4, req.Scene)

	c.Set("Payload", ActionStatus{
		Success: status == 1,
	})

}

// @Summary 接受转账
// @Produce json
// @Param body body wcf.Transfer true "接受转账请求参数"
// @Success 200 {object} ActionStatus
// @Router /receive_transfer [post]
func receiveTransfer(c *gin.Context) {

	var req wcf.Transfer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.ReceiveTransfer(req.Wxid, req.Tfid, req.Taid)

	c.Set("Payload", ActionStatus{
		Success: status == 1,
	})

}

// @Summary 添加群成员
// @Produce json
// @Param body body wcf.AddMembers true "增删群成员请求参数"
// @Success 200 {object} ActionStatus
// @Router /add_chatroom_members [post]
func addChatRoomMembers(c *gin.Context) {

	var req wcf.AddMembers
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.AddChatRoomMembers(req.Roomid, req.Wxids)

	c.Set("Payload", ActionStatus{
		Success: status == 1,
	})

}

// @Summary 删除群成员
// @Produce json
// @Param body body wcf.AddMembers true "增删群成员请求参数"
// @Success 200 {object} ActionStatus
// @Router /del_chatroom_members [post]
func delChatRoomMembers(c *gin.Context) {

	var req wcf.AddMembers
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.DelChatRoomMembers(req.Roomid, req.Wxids)

	c.Set("Payload", ActionStatus{
		Success: status == 1,
	})

}

// @Summary 解密图片
// @Produce json
// @Param body body wcf.DecPath true "解密图片请求参数"
// @Success 200 {object} ActionStatus
// @Router /decrypt_image [post]
func decryptImage(c *gin.Context) {

	var req wcf.DecPath
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.DecryptImage(req.Src, req.Dst)

	c.Set("Payload", ActionStatus{
		Success: status == 1,
	})

}

// @Summary 开启消息转发
// @Produce json
// @Param body body ForwardMsgRequest true "消息转发请求参数"
// @Success 200 {object} ActionStatus
// @Router /enable_forward_msg [post]
func enableForwardMsg(c *gin.Context) {

	var req ForwardMsgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	cb := func(msg *wcf.WxMsg) {
		logman.Info("forward msg", "url", req.Url, "msg", msg)
		request.JsonPost(req.Url, msg, request.H{})
	}

	error := wc.EnrollReceiver(true, cb)

	c.Set("Payload", ActionStatus{
		Success: error == nil,
		Error:   error,
	})

}

// @Summary 关闭消息转发
// @Produce json
// @Param body body ForwardMsgRequest true "消息转发请求参数"
// @Success 200 {object} ActionStatus
// @Router /disable_forward_msg [post]
func disableForwardMsg(c *gin.Context) {

	error := wc.DisableReceiver()

	c.Set("Payload", ActionStatus{
		Success: error == nil,
		Error:   error,
	})

}
