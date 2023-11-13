package wcf

import (
	"github.com/gin-gonic/gin"
)

func Route(rg *gin.RouterGroup) {

	go initWCF()

	rg.GET("is_login", isLogin)
	rg.GET("wxid", wxid)
	rg.GET("user_info", userInfo)
	rg.GET("msg_types", msgTypes)
	rg.GET("contacts", contacts)
	rg.GET("friends", friends)
	rg.GET("db_names", dbNames)
	rg.GET("db_tables/:tab", dbTables)
	rg.GET("pyq_refresh/:id", pyqRefresh)
	rg.GET("chatroom_members/:roomid", chatroomMembers)

}
