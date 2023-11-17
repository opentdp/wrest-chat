package wcf

import (
	"strings"

	"google.golang.org/protobuf/proto"
)

type CmdClient struct {
	pbSocket // RPC 客户端
}

// 关闭 RPC 连接
// return error 错误信息
func (c *CmdClient) Close() error {
	return c.close()
}

// 检查登录状态
// return bool 是否已登录
func (c *CmdClient) IsLogin() bool {
	req := genFunReq(Functions_FUNC_IS_LOGIN)
	recv := c.call(req.build())
	if recv.GetStatus() == 1 {
		return true
	}
	return false
}

// 获取登录账号wxid
// return string 登录账号wxid
func (c *CmdClient) GetSelfWxid() string {
	req := genFunReq(Functions_FUNC_GET_SELF_WXID)
	recv := c.call(req.build())
	return recv.GetStr()
}

// 获取登录账号个人信息
// return *UserInfo 登录账号个人信息
func (c *CmdClient) GetUserInfo() *UserInfo {
	req := genFunReq(Functions_FUNC_GET_USER_INFO)
	recv := c.call(req.build())
	return recv.GetUi()
}

// 获取所有消息类型
// return map[int32]string 所有消息类型
func (c *CmdClient) GetMsgTypes() map[int32]string {
	req := genFunReq(Functions_FUNC_GET_MSG_TYPES)
	recv := c.call(req.build())
	return recv.GetTypes().GetTypes()
}

// 获取完整通讯录
// return []*RpcContact 完整通讯录
func (c *CmdClient) GetContacts() []*RpcContact {
	req := genFunReq(Functions_FUNC_GET_CONTACTS)
	recv := c.call(req.build())
	return recv.GetContacts().GetContacts()
}

// 获取好友列表
// return []*RpcContact 好友列表
func (c *CmdClient) GetFriends() []*RpcContact {
	notFriends := map[string]string{
		"mphelper":    "公众平台助手",
		"fmessage":    "朋友推荐消息",
		"medianote":   "语音记事本",
		"floatbottle": "漂流瓶",
		"filehelper":  "文件传输助手",
		"newsapp":     "新闻",
	}
	friends := []*RpcContact{}
	for _, cnt := range c.GetContacts() {
		if strings.HasSuffix(cnt.Wxid, "@chatroom") || strings.HasPrefix(cnt.Wxid, "gh_") || notFriends[cnt.Wxid] != "" {
			continue
		}
		friends = append(friends, cnt)
	}
	return friends
}

// 获取群聊列表
// return []*RpcContact 群聊列表
func (c *CmdClient) GetChatRooms() []*RpcContact {
	chatrooms := []*RpcContact{}
	for _, cnt := range c.GetContacts() {
		if strings.HasSuffix(cnt.Wxid, "@chatroom") {
			chatrooms = append(chatrooms, cnt)
		}
	}
	return chatrooms
}

// 获取所有数据库
// return []string 所有数据库名
func (c *CmdClient) GetDbNames() []string {
	req := genFunReq(Functions_FUNC_GET_DB_NAMES)
	recv := c.call(req.build())
	return recv.GetDbs().Names
}

// 获取数据库中所有表
// param db string 数据库名
// return []*DbTable `db` 下的所有表名及对应建表语句
func (c *CmdClient) GetDbTables(db string) []*DbTable {
	req := genFunReq(Functions_FUNC_GET_DB_TABLES)
	req.Msg = &Request_Str{
		Str: db,
	}
	recv := c.call(req.build())
	return recv.GetTables().GetTables()
}

// 执行 SQL 查询，如果数据量大注意分页
// param db string 要查询的数据库
// param sql string 要执行的 SQL
// return []*DbRow 查询结果
func (c *CmdClient) DbSqlQuery(db, sql string) []*DbRow {
	req := genFunReq(Functions_FUNC_EXEC_DB_QUERY)
	req.Msg = &Request_Query{
		Query: &DbQuery{
			Db:  db,
			Sql: sql,
		},
	}
	recv := c.call(req.build())
	return recv.GetRows().GetRows()
}

// 执行 SQL 查询，如果数据量大注意分页
// param db string 要查询的数据库
// param sql string 要执行的 SQL
// return map[string]any 查询结果
func (c *CmdClient) DbSqlQueryMap(db, sql string) map[string]any {
	rows := c.DbSqlQuery(db, sql)
	res := map[string]any{}
	for _, row := range rows {
		for _, field := range row.Fields {
			res[field.Column] = field.Content
		}
	}
	return res
}

// 发送文本消息
// param msg string 要发送的消息，换行使用 `\\\\n` （单杠）；如果 @ 人的话，需要带上跟 `aters` 里数量相同的 @
// param receiver string 消息接收人，wxid 或者 roomid
// param aters string 要 @ 的 wxid，多个用逗号分隔；`@所有人` 只需要 `notify@all`
// return int32 0 为成功，其他失败
func (c *CmdClient) SendTxt(msg, receiver, aters string) int32 {
	req := genFunReq(Functions_FUNC_SEND_TXT)
	req.Msg = &Request_Txt{
		Txt: &TextMsg{
			Msg:      msg,
			Receiver: receiver,
			Aters:    aters,
		},
	}
	recv := c.call(req.build())
	return recv.GetStatus()
}

// 发送图片，非线程安全
// param path string 图片路径，如：`C:/Projs/WeChatRobot/TEQuant.jpeg`
// param receiver string 消息接收人，wxid 或者 roomid
// return int32 0 为成功，其他失败
func (c *CmdClient) SendImg(path, receiver string) int32 {
	req := genFunReq(Functions_FUNC_SEND_IMG)
	req.Msg = &Request_File{
		File: &PathMsg{
			Path:     path,
			Receiver: receiver,
		},
	}
	recv := c.call(req.build())
	return recv.GetStatus()
}

// 发送文件，非线程安全
// param path string 本地文件路径，如：`C:/Projs/WeChatRobot/README.MD`
// param receiver string 消息接收人，wxid 或者 roomid
// return int32 0 为成功，其他失败
func (c *CmdClient) SendFile(path, receiver string) int32 {
	req := genFunReq(Functions_FUNC_SEND_FILE)
	req.Msg = &Request_File{
		File: &PathMsg{
			Path:     path,
			Receiver: receiver,
		},
	}
	recv := c.call(req.build())
	return recv.GetStatus()
}

// 发送 XML
// param path string 封面图片路径
// param content string xml 内容
// param receiver string 消息接收人，wxid 或者 roomid
// param Type int32 xml 类型，如：0x21 为小程序
// return int32 0 为成功，其他失败
func (c *CmdClient) SendXml(path, content, receiver string, Type int32) int32 {
	req := genFunReq(Functions_FUNC_SEND_XML)
	req.Msg = &Request_Xml{
		Xml: &XmlMsg{
			Receiver: receiver,
			Content:  content,
			Path:     path,
			Type:     Type,
		},
	}
	recv := c.call(req.build())
	return recv.GetStatus()
}

// 发送表情
// param path string 本地表情路径，如：`C:/Projs/WeChatRobot/emo.gif`
// param receiver string 消息接收人，wxid 或者 roomid
// return int32 0 为成功，其他失败
func (c *CmdClient) SendEmotion(path, receiver string) int32 {
	req := genFunReq(Functions_FUNC_SEND_EMOTION)
	req.Msg = &Request_File{
		File: &PathMsg{
			Path:     path,
			Receiver: receiver,
		},
	}
	recv := c.call(req.build())
	return recv.GetStatus()
}

// 接受好友申请
// param v3 string 加密用户名 (好友申请消息里 v3 开头的字符串)
// param v4 string Ticket (好友申请消息里 v4 开头的字符串)
// param scene int32 申请方式 (好友申请消息里的 scene); 为了兼容旧接口，默认为扫码添加 (30)
// return int32 1 为成功，其他失败
func (c *CmdClient) AcceptNewFriend(v3, v4 string, scene int32) int32 {
	req := genFunReq(Functions_FUNC_ACCEPT_FRIEND)
	req.Msg = &Request_V{
		V: &Verification{
			V3:    v3,
			V4:    v4,
			Scene: scene,
		},
	}
	recv := c.call(req.build())
	return recv.GetStatus()
}

// 接收转账
// param wxid string 转账消息里的发送人 wxid
// param transferid string 转账消息里的 transferid
// param transactionid string 转账消息里的 transactionid
// return int32 1 为成功，其他失败
func (c *CmdClient) ReceiveTransfer(wxid, tfid, taid string) int32 {
	req := genFunReq(Functions_FUNC_RECV_TRANSFER)
	req.Msg = &Request_Tf{
		Tf: &Transfer{
			Wxid: wxid,
			Tfid: tfid,
			Taid: taid,
		},
	}
	recv := c.call(req.build())
	return recv.GetStatus()
}

// 刷新朋友圈
// param id int32 开始 id，0 为最新页
// return int32 1 为成功，其他失败
func (c *CmdClient) RefreshPyq(id uint64) int32 {
	req := genFunReq(Functions_FUNC_REFRESH_PYQ)
	req.Msg = &Request_Ui64{
		Ui64: id,
	}
	recv := c.call(req.build())
	return recv.GetStatus()
}

// 添加群成员
// param roomid string 待加群的 id
// param wxids string 要加到群里的 wxid，多个用逗号分隔
// return int32 1 为成功，其他失败
func (c *CmdClient) AddChatRoomMembers(roomId, wxIds string) int32 {
	req := genFunReq(Functions_FUNC_ADD_ROOM_MEMBERS)
	req.Msg = &Request_M{
		M: &AddMembers{
			Roomid: roomId,
			Wxids:  wxIds,
		},
	}
	recv := c.call(req.build())
	return recv.GetStatus()
}

// 删除群成员
// param roomid string 群的 id
// param wxids string 要删除成员的 wxid，多个用逗号分隔
// return int32 1 为成功，其他失败
func (c *CmdClient) DelChatRoomMembers(roomId, wxIds string) int32 {
	req := genFunReq(Functions_FUNC_DEL_ROOM_MEMBERS)
	req.Msg = &Request_M{
		M: &AddMembers{
			Roomid: roomId,
			Wxids:  wxIds,
		},
	}
	recv := c.call(req.build())
	return recv.GetStatus()
}

// 获取群成员列表
// param roomid string 群的 id
// return []*RpcContact 群成员列表
func (c *CmdClient) GetChatRoomMembers(roomId string) []*RpcContact {
	members := []*RpcContact{}
	// get user data
	userRds := c.DbSqlQuery("MicroMsg.db", "SELECT UserName, NickName FROM Contact;")
	userMap := map[string]string{}
	for _, user := range userRds {
		wxid := string(user.Fields[0].Content)
		userMap[wxid] = string(user.Fields[1].Content)
	}
	// get room data
	roomRds := c.DbSqlQuery("MicroMsg.db", "SELECT RoomData FROM ChatRoom WHERE ChatRoomName = '"+roomId+"';")
	if len(roomRds) == 0 || len(roomRds[0].Fields) == 0 {
		return members
	}
	roomData := &RoomData{}
	if err := proto.Unmarshal(roomRds[0].Fields[0].Content, roomData); err != nil {
		return members
	}
	// fix user name
	for _, member := range roomData.Members {
		if member.Name == "" {
			member.Name = userMap[member.Wxid]
		}
		members = append(members, &RpcContact{
			Wxid: member.Wxid,
			Name: member.Name,
		})
	}
	return members
}

// 获取群成员昵称
// param wxid string wxid
// param roomid string 群的 id
// return string 群成员昵称
func (c *CmdClient) GetAliasInChatRoom(wxid, roomId string) string {
	// get user data
	nickName := ""
	userRds := c.DbSqlQuery("MicroMsg.db", "SELECT NickName FROM Contact WHERE UserName = '"+wxid+"';")
	if len(userRds) > 0 && len(userRds[0].Fields) > 0 {
		nickName = string(userRds[0].Fields[0].Content)
	}
	// get room data
	roomRds := c.DbSqlQuery("MicroMsg.db", "SELECT RoomData FROM ChatRoom WHERE ChatRoomName = '"+roomId+"';")
	if len(roomRds) == 0 || len(roomRds[0].Fields) == 0 {
		return nickName
	}
	roomData := &RoomData{}
	if err := proto.Unmarshal(roomRds[0].Fields[0].Content, roomData); err != nil {
		return nickName
	}
	// fix user name
	for _, member := range roomData.Members {
		if member.Wxid == wxid {
			if member.Name != "" {
				nickName = member.Name
			}
			break
		}
	}
	return nickName
}

// 解密图片
// param src string 加密的图片路径
// param dst string 解密的图片路径
// return int32 1 为成功，其他失败
func (c *CmdClient) DecryptImage(src, dst string) int32 {
	req := genFunReq(Functions_FUNC_DECRYPT_IMAGE)
	req.Msg = &Request_Dec{
		Dec: &DecPath{
			Src: src,
			Dst: dst,
		},
	}
	recv := c.call(req.build())
	return recv.GetStatus()
}

// 开启消息服务器
// param pyq bool 是否接收朋友圈消息
// return int32 0 为成功，其他失败
func (c *CmdClient) EnableMsgServer(pyq bool) int32 {
	req := genFunReq(Functions_FUNC_ENABLE_RECV_TXT)
	req.Msg = &Request_Flag{
		Flag: pyq,
	}
	recv := c.call(req.build())
	return recv.GetStatus()
}

// 停止消息服务器
// return int32 0 为成功，其他失败
func (c *CmdClient) DisableMsgServer() int32 {
	req := genFunReq(Functions_FUNC_DISABLE_RECV_TXT)
	recv := c.call(req.build())
	return recv.GetStatus()
}
