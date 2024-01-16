package wcfrest

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/strutil"
)

func deprecated(rg *gin.RouterGroup, ctrl *Controller) {
	rg.GET("is_login", ctrl.isLogin)
	rg.GET("self_wxid", ctrl.getSelfWxid)
	rg.GET("self_info", ctrl.getSelfInfo)
	rg.GET("msg_types", ctrl.getMsgTypes)
	rg.GET("db_names", ctrl.getDbNames)
	rg.GET("db_tables/:db", ctrl._getDbTables)
	rg.GET("chatrooms", ctrl.getChatRooms)
	rg.GET("chatroom_members/:roomid", ctrl._getChatRoomMembers)
	rg.GET("alias_in_chatroom/:wxid/:roomid", ctrl._getAliasInChatRoom)
	rg.GET("revoke_msg/:msgid", ctrl._revokeMsg)
	rg.GET("avatars", ctrl.getAvatars)
	rg.GET("contacts", ctrl.getContacts)
	rg.GET("friends", ctrl.getFriends)
	rg.GET("user_info/:wxid", ctrl._getInfoByWxid)
	rg.GET("refresh_pyq/:id", ctrl._refreshPyq)
}

func (wc *Controller) _getDbTables(c *gin.Context) {

	db := c.Param("db")
	c.Set("Payload", wc.CmdClient.GetDbTables(db))

}

func (wc *Controller) _getChatRoomMembers(c *gin.Context) {

	roomid := c.Param("roomid")
	c.Set("Payload", wc.CmdClient.GetChatRoomMembers(roomid))

}

func (wc *Controller) _getAliasInChatRoom(c *gin.Context) {

	wxid := c.Param("wxid")
	roomid := c.Param("roomid")
	c.Set("Payload", wc.CmdClient.GetAliasInChatRoom(wxid, roomid))

}

func (wc *Controller) _revokeMsg(c *gin.Context) {

	id := c.Param("msgid")
	msgid := uint64(strutil.ToUint(id))
	status := wc.CmdClient.RevokeMsg(msgid)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

func (wc *Controller) _getInfoByWxid(c *gin.Context) {

	wxid := c.Param("wxid")
	c.Set("Payload", wc.CmdClient.GetInfoByWxid(wxid))

}

func (wc *Controller) _refreshPyq(c *gin.Context) {

	id := c.Param("id")
	pyqid := uint64(strutil.ToUint(id))
	status := wc.CmdClient.RefreshPyq(pyqid)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}
