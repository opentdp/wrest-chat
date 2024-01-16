package wcfrest

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/opentdp/wechat-rest/wcferry"
)

type Controller struct {
	*wcferry.Client
}

// 通用结果
type CommonPayload struct {
	Success bool   `json:"success,omitempty"`
	Result  string `json:"result,omitempty"`
	Error   error  `json:"error,omitempty"`
}

// @Summary 检查登录状态
// @Produce json
// @Success 200 {object} bool
// @Router /is_login [post]
func (wc *Controller) isLogin(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.IsLogin())

}

// @Summary 获取登录账号wxid
// @Produce json
// @Success 200 {object} string
// @Router /self_wxid [post]
func (wc *Controller) getSelfWxid(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetSelfWxid())

}

// @Summary 获取登录账号个人信息
// @Produce json
// @Success 200 {object} wcferry.UserInfo
// @Router /self_info [post]
func (wc *Controller) getSelfInfo(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetSelfInfo())

}

// @Summary 获取所有消息类型
// @Produce json
// @Success 200 {object} map[int32]string
// @Router /msg_types [post]
func (wc *Controller) getMsgTypes(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetMsgTypes())

}

// @Summary 获取数据库列表
// @Produce json
// @Success 200 {object} []string
// @Router /db_names [post]
func (wc *Controller) getDbNames(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetDbNames())

}

// @Summary 获取数据库表列表
// @Produce json
// @Param body body GetDbTablesRequest true "获取数据库表列表参数"
// @Success 200 {object} []wcferry.DbTable
// @Router /db_tables [post]
func (wc *Controller) getDbTables(c *gin.Context) {

	var req GetDbTablesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", wc.CmdClient.GetDbTables(req.Db))

}

type GetDbTablesRequest struct {
	Db string `json:"db"`
}

// @Summary 执行数据库查询
// @Produce json
// @Param body body DbSqlQueryRequest true "数据库查询参数"
// @Success 200 {object} []map[string]any
// @Router /db_query_sql [post]
func (wc *Controller) dbSqlQuery(c *gin.Context) {

	var req DbSqlQueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", wc.CmdClient.DbSqlQuery(req.Db, req.Sql))

}

type DbSqlQueryRequest struct {
	Db  string `json:"db"`
	Sql string `json:"sql"`
}

// @Summary 获取群列表
// @Produce json
// @Success 200 {object} []wcferry.RpcContact
// @Router /chatrooms [post]
func (wc *Controller) getChatRooms(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetChatRooms())

}

// @Summary 获取群成员列表
// @Produce json
// @Param body body GetChatRoomMembersRequest true "获取群成员列表参数"
// @Success 200 {object} []wcferry.RpcContact
// @Router /chatroom_members [post]
func (wc *Controller) getChatRoomMembers(c *gin.Context) {

	var req GetChatRoomMembersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", wc.CmdClient.GetChatRoomMembers(req.Roomid))

}

type GetChatRoomMembersRequest struct {
	Roomid string `json:"roomid"`
}

// @Summary 获取群成员昵称
// @Produce json
// @Param body body GetAliasInChatRoomRequest true "获取群成员昵称参数"
// @Success 200 {object} string
// @Router /alias_in_chatroom [post]
func (wc *Controller) getAliasInChatRoom(c *gin.Context) {

	var req GetAliasInChatRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", wc.CmdClient.GetAliasInChatRoom(req.Wxid, req.Roomid))

}

type GetAliasInChatRoomRequest struct {
	Roomid string `json:"roomid"`
	Wxid   string `json:"wxid"`
}

// @Summary 邀请群成员
// @Produce json
// @Param body body wcferry.MemberMgmt true "管理群成员参数"
// @Success 200 {object} CommonPayload
// @Router /invite_chatroom_members [post]
func (wc *Controller) inviteChatroomMembers(c *gin.Context) {

	var req wcferry.MemberMgmt
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.InviteChatroomMembers(req.Roomid, req.Wxids)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

// @Summary 添加群成员
// @Produce json
// @Param body body wcferry.MemberMgmt true "管理群成员参数"
// @Success 200 {object} CommonPayload
// @Router /add_chatroom_members [post]
func (wc *Controller) addChatRoomMembers(c *gin.Context) {

	var req wcferry.MemberMgmt
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.AddChatRoomMembers(req.Roomid, req.Wxids)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

// @Summary 删除群成员
// @Produce json
// @Param body body wcferry.MemberMgmt true "管理群成员参数"
// @Success 200 {object} CommonPayload
// @Router /del_chatroom_members [post]
func (wc *Controller) delChatRoomMembers(c *gin.Context) {

	var req wcferry.MemberMgmt
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.DelChatRoomMembers(req.Roomid, req.Wxids)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

// @Summary 撤回消息
// @Produce json
// @Param body body RevokeMsgRequest true "撤回消息参数"
// @Success 200 {object} CommonPayload
// @Router /revoke_msg [post]
func (wc *Controller) revokeMsg(c *gin.Context) {

	var req RevokeMsgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.RevokeMsg(req.Msgid)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

type RevokeMsgRequest struct {
	Msgid uint64 `json:"msgid"`
}

// @Summary 转发消息
// @Produce json
// @Param body body wcferry.ForwardMsg true "转发消息参数"
// @Success 200 {object} CommonPayload
// @Router /forward_msg [post]
func (wc *Controller) forwardMsg(c *gin.Context) {

	var req wcferry.ForwardMsg
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.ForwardMsg(req.Id, req.Receiver)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

// @Summary 发送文本消息
// @Produce json
// @Param body body wcferry.TextMsg true "发送文本消息参数"
// @Success 200 {object} CommonPayload
// @Router /send_txt [post]
func (wc *Controller) sendTxt(c *gin.Context) {

	var req wcferry.TextMsg
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.SendTxt(req.Msg, req.Receiver, req.Aters)

	c.Set("Payload", CommonPayload{
		Success: status == 0,
	})

}

// @Summary 发送图片消息
// @Produce json
// @Param body body wcferry.PathMsg true "发送图片消息参数"
// @Success 200 {object} CommonPayload
// @Router /send_img [post]
func (wc *Controller) sendImg(c *gin.Context) {

	var req wcferry.PathMsg
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.SendImg(req.Path, req.Receiver)

	c.Set("Payload", CommonPayload{
		Success: status == 0,
	})

}

// @Summary 发送文件消息
// @Produce json
// @Param body body wcferry.PathMsg true "发送文件消息参数"
// @Success 200 {object} CommonPayload
// @Router /send_file [post]
func (wc *Controller) sendFile(c *gin.Context) {

	var req wcferry.PathMsg
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.SendFile(req.Path, req.Receiver)

	c.Set("Payload", CommonPayload{
		Success: status == 0,
	})

}

// @Summary 发送卡片消息
// @Produce json
// @Param body body wcferry.RichText true "发送卡片消息参数"
// @Success 200 {object} CommonPayload
// @Router /send_rich_text [post]
func (wc *Controller) sendRichText(c *gin.Context) {

	var req wcferry.RichText
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.SendRichText(req.Name, req.Account, req.Title, req.Digest, req.Url, req.Thumburl, req.Receiver)

	c.Set("Payload", CommonPayload{
		Success: status == 0,
	})

}

// @Summary 拍一拍群友
// @Produce json
// @Param body body wcferry.PatMsg true "拍一拍群友参数"
// @Success 200 {object} CommonPayload
// @Router /send_pat_msg [post]
func (wc *Controller) sendPatMsg(c *gin.Context) {

	var req wcferry.PatMsg
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.SendPatMsg(req.Roomid, req.Wxid)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

// @Summary 获取语音消息
// @Produce json
// @Param body body GetAudioMsgRequest true "获取语音消息参数"
// @Success 200 {object} CommonPayload
// @Router /get_audio_msg [post]
func (wc *Controller) getAudioMsg(c *gin.Context) {

	var req GetAudioMsgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	if req.Timeout > 0 {
		resp, err := wc.CmdClient.GetAudioMsgTimeout(req.Msgid, req.Dir, req.Timeout)
		c.Set("Payload", CommonPayload{
			Success: resp != "",
			Result:  resp,
			Error:   err,
		})
	} else {
		resp := wc.CmdClient.GetAudioMsg(req.Msgid, req.Dir)
		c.Set("Payload", CommonPayload{
			Success: resp != "",
			Result:  resp,
		})
	}

}

type GetAudioMsgRequest struct {
	Msgid   uint64 `json:"msgid"`
	Dir     string `json:"path"`
	Timeout int    `json:"timeout"`
}

// @Summary 获取OCR识别结果
// @Produce json
// @Param body body GetOcrRequest true "获取OCR识别结果参数"
// @Success 200 {object} CommonPayload
// @Router /get_ocr_result [post]
func (wc *Controller) getOcrResult(c *gin.Context) {

	var req GetOcrRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	if req.Timeout > 0 {
		resp, err := wc.CmdClient.GetOcrResultTimeout(req.Extra, req.Timeout)
		c.Set("Payload", CommonPayload{
			Success: resp != "",
			Result:  resp,
			Error:   err,
		})
	} else {
		resp, stat := wc.CmdClient.GetOcrResult(req.Extra)
		c.Set("Payload", CommonPayload{
			Success: stat == 0,
			Result:  resp,
		})
	}

}

type GetOcrRequest struct {
	Extra   string `json:"extra"`
	Timeout int    `json:"timeout"`
}

// @Summary 下载图片
// @Produce json
// @Param body body DownloadImageRequest true "下载图片参数"
// @Success 200 {object} CommonPayload
// @Router /download_image [post]
func (wc *Controller) downloadImage(c *gin.Context) {

	var req DownloadImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	resp, err := wc.CmdClient.DownloadImage(req.Msgid, req.Extra, req.Dir, req.Timeout)

	c.Set("Payload", CommonPayload{
		Success: resp != "",
		Result:  resp,
		Error:   err,
	})

}

type DownloadImageRequest struct {
	Msgid   uint64 `json:"msgid"`
	Extra   string `json:"extra"`
	Dir     string `json:"dir"`
	Timeout int    `json:"timeout"`
}

// @Summary 下载附件
// @Produce json
// @Param body body DownloadAttachRequest true "下载附件参数"
// @Success 200 {object} CommonPayload
// @Router /download_attach [post]
func (wc *Controller) downloadAttach(c *gin.Context) {

	var req DownloadAttachRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.DownloadAttach(req.Msgid, req.Thumb, req.Extra)

	c.Set("Payload", CommonPayload{
		Success: status == 0,
	})

}

type DownloadAttachRequest struct {
	Msgid uint64 `json:"msgid"`
	Thumb string `json:"thumb"`
	Extra string `json:"extra"`
}

// @Summary 获取头像列表
// @Produce json
// @Success 200 {object} []wcferry.ContactHeadImgUrlTable
// @Router /avatars [post]
func (wc *Controller) getAvatars(c *gin.Context) {

	sql := "SELECT usrName,bigHeadImgUrl,smallHeadImgUrl FROM ContactHeadImgUrl"
	res := wc.CmdClient.DbSqlQuery("MicroMsg.db", sql)

	c.Set("Payload", res)

}

// @Summary 获取完整通讯录
// @Produce json
// @Success 200 {object} []wcferry.RpcContact
// @Router /contacts [post]
func (wc *Controller) getContacts(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetContacts())

}

// @Summary 获取好友列表
// @Produce json
// @Success 200 {object} []wcferry.RpcContact
// @Router /friends [post]
func (wc *Controller) getFriends(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetFriends())

}

// @Summary 根据wxid获取个人信息
// @Produce json
// @Param body body GetInfoByWxidRequest true "根据wxid获取个人信息参数"
// @Success 200 {object} wcferry.RpcContact
// @Router /user_info [post]
func (wc *Controller) getInfoByWxid(c *gin.Context) {

	var req GetInfoByWxidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", wc.CmdClient.GetInfoByWxid(req.Wxid))

}

type GetInfoByWxidRequest struct {
	Wxid string `json:"wxid"`
}

// @Summary 刷新朋友圈
// @Produce json
// @Param body body RefreshPyqRequest true "刷新朋友圈参数"
// @Success 200 {object} CommonPayload
// @Router /refresh_pyq [post]
func (wc *Controller) refreshPyq(c *gin.Context) {

	var req RefreshPyqRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.RefreshPyq(req.Id)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

type RefreshPyqRequest struct {
	Id uint64 `json:"id"`
}

// @Summary 接受好友请求
// @Produce json
// @Param body body wcferry.Verification true "接受好友参数"
// @Success 200 {object} CommonPayload
// @Router /accept_new_friend [post]
func (wc *Controller) acceptNewFriend(c *gin.Context) {

	var req wcferry.Verification
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.AcceptNewFriend(req.V3, req.V4, req.Scene)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

// @Summary 接受转账
// @Produce json
// @Param body body wcferry.Transfer true "接受转账参数"
// @Success 200 {object} CommonPayload
// @Router /receive_transfer [post]
func (wc *Controller) receiveTransfer(c *gin.Context) {

	var req wcferry.Transfer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.ReceiveTransfer(req.Wxid, req.Tfid, req.Taid)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

// @Summary 开启推送消息到URL
// @Produce json
// @Param body body ReceiverRequest true "推送消息到URL参数"
// @Success 200 {object} CommonPayload
// @Router /enable_receiver [post]
func (wc *Controller) enabledReceiver(c *gin.Context) {

	var req ReceiverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	if !strings.HasPrefix(req.Url, "http") {
		c.Set("Error", "url must start with http(s)://")
		return
	}

	err := wc.enableUrlReceiver(req.Url)
	c.Set("Payload", CommonPayload{
		Success: err == nil,
		Error:   err,
	})

}

type ReceiverRequest struct {
	Url string `json:"url"`
}

// @Summary 关闭推送消息到URL
// @Produce json
// @Param body body ReceiverRequest true "推送消息到URL参数"
// @Success 200 {object} CommonPayload
// @Router /disable_receiver [post]
func (wc *Controller) disableReceiver(c *gin.Context) {

	var req ReceiverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	err := wc.disableUrlReceiver(req.Url)
	c.Set("Payload", CommonPayload{
		Success: err == nil,
		Error:   err,
	})

}
