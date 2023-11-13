package wcf

import (
	"github.com/gin-gonic/gin"
)

func Route(rg *gin.RouterGroup) {

	go initWCF()

	rg.GET("is_login", isLogin)
	rg.GET("self_wxid", getSelfWxid)
	rg.GET("user_info", getUserInfo)
	rg.GET("msg_types", getMsgTypes)
	rg.GET("contacts", getContacts)
	rg.GET("friends", getFriends)

	rg.GET("db_names", getDbNames)
	rg.GET("db_tables/:db", getDbTables)
	rg.POST("db_sql_query", dbSqlQuery)

	rg.GET("chatrooms", getChatRooms)
	rg.GET("chatroom_members/:roomid", getChatRoomMembers)
	rg.GET("alias_in_chatroom/:wxid/:roomid", getAliasInChatRoom)
	rg.POST("add_chatroom_members", addChatRoomMembers)
	rg.POST("del_chatroom_members", delChatRoomMembers)

	rg.POST("send_txt", sendTxt)
	rg.POST("send_img", sendImg)
	rg.POST("send_file", sendFile)
	rg.POST("accept_new_friend", acceptNewFriend)
	rg.POST("receive_transfer", receiveTransfer)

	rg.GET("refresh_pyq/:id", refreshPyq)
	rg.POST("decrypt_image", decryptImage)

	rg.GET("enable_receiving_msg", enableReceivingMsg)
	rg.GET("disable_receiving_msg", disableReceivingMsg)

}
