package wcf

import (
	"strings"

	"github.com/opentdp/go-helper/logman"
	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol"
	"go.nanomsg.org/mangos/v3/protocol/pair1"
	_ "go.nanomsg.org/mangos/v3/transport/all"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	server string
	socket protocol.Socket

	IsReceivingMsg bool
}

func (c *Client) dial() error {
	socket, err := pair1.NewSocket()
	if err != nil {
		return err
	}
	err = socket.Dial(c.server)
	if err != nil {
		return err
	}
	c.socket = socket
	return err
}

func (c *Client) send(data []byte) error {
	return c.socket.Send(data)
}

func (c *Client) recv() (*Response, error) {
	resp := &Response{}
	recv, err := c.socket.Recv()
	if err == nil {
		err = proto.Unmarshal(recv, resp)
	}
	return resp, err
}

// 调用 RPC 接口
func (c *Client) Call(data []byte) *Response {
	if err := c.send(data); err != nil {
		logman.Error(err.Error())
	}
	recv, err := c.recv()
	if err != nil {
		logman.Error(err.Error())
	}
	return recv
}

// 关闭 RPC 连接
func (c *Client) Close() error {
	c.DisableReceivingMsg()
	return c.socket.Close()
}

// 检查登录状态
func (c *Client) IsLogin() bool {
	req := genFunReq(Functions_FUNC_IS_LOGIN)
	recv := c.Call(req.build())
	if recv.GetStatus() == 1 {
		return true
	}
	return false
}

// 获取登录账号wxid
func (c *Client) GetSelfWxid() string {
	req := genFunReq(Functions_FUNC_GET_SELF_WXID)
	recv := c.Call(req.build())
	return recv.GetStr()
}

// 获取登录账号个人信息
func (c *Client) GetUserInfo() *UserInfo {
	req := genFunReq(Functions_FUNC_GET_USER_INFO)
	recv := c.Call(req.build())
	return recv.GetUi()
}

// 获取所有消息类型
func (c *Client) GetMsgTypes() map[int32]string {
	req := genFunReq(Functions_FUNC_GET_MSG_TYPES)
	recv := c.Call(req.build())
	return recv.GetTypes().GetTypes()
}

// 获取完整通讯录
func (c *Client) GetContacts() []*RpcContact {
	req := genFunReq(Functions_FUNC_GET_CONTACTS)
	recv := c.Call(req.build())
	return recv.GetContacts().GetContacts()
}

// 获取好友列表
func (c *Client) GetFriends() []*RpcContact {
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

// 获取所有数据库
func (c *Client) GetDbNames() []string {
	req := genFunReq(Functions_FUNC_GET_DB_NAMES)
	recv := c.Call(req.build())
	return recv.GetDbs().Names
}

// 获取数据库中所有表
// Args:
//
//	db string: 数据库名
//
// Returns:
//
//	[]*DbTable: `db` 下的所有表名及对应建表语句
func (c *Client) GetDbTables(db string) []*DbTable {
	req := genFunReq(Functions_FUNC_GET_DB_TABLES)
	req.Msg = &Request_Str{
		Str: db,
	}
	recv := c.Call(req.build())
	return recv.GetTables().GetTables()
}

// 执行 SQL 查询，如果数据量大注意分页
// Args:
//
//	db string: 要查询的数据库
//	sql string: 要执行的 SQL
//
// Returns:
//
//	[]*DbRow: 查询结果
func (c *Client) ExecDbQuery(db, sql string) []*DbRow {
	req := genFunReq(Functions_FUNC_EXEC_DB_QUERY)
	req.Msg = &Request_Query{
		Query: &DbQuery{
			Db:  db,
			Sql: sql,
		},
	}
	recv := c.Call(req.build())
	return recv.GetRows().GetRows()
}

// 发送文本消息
//
// Args:
//
//	msg (str): 要发送的消息，换行使用 `\\\\n` （单杠）；如果 @ 人的话，需要带上跟 `aters` 里数量相同的 @
//	receiver (str): 消息接收人，wxid 或者 roomid
//	aters (str): 要 @ 的 wxid，多个用逗号分隔；`@所有人` 只需要 `notify@all`
//
// Returns:
//
//	int: 0 为成功，其他失败
func (c *Client) SendTxt(msg string, receiver string, ates []string) int32 {
	req := genFunReq(Functions_FUNC_SEND_TXT)
	req.Msg = &Request_Txt{
		Txt: &TextMsg{
			Msg:      msg,
			Receiver: receiver,
			Aters:    strings.Join(ates, ","),
		},
	}
	recv := c.Call(req.build())
	return recv.GetStatus()
}

// 发送图片，非线程安全
//
// Args:
//
//	path (str): 图片路径，如：`C:/Projs/WeChatRobot/TEQuant.jpeg` 或 `https://raw.githubusercontent.com/lich0821/WeChatFerry/master/assets/TEQuant.jpg`
//	receiver (str): 消息接收人，wxid 或者 roomid
//
// Returns:
//
//	int: 0 为成功，其他失败
func (c *Client) SendIMG(path string, receiver string) int32 {
	req := genFunReq(Functions_FUNC_SEND_IMG)
	req.Msg = &Request_File{
		File: &PathMsg{
			Path:     path,
			Receiver: receiver,
		},
	}
	recv := c.Call(req.build())
	return recv.GetStatus()
}

// 发送文件，非线程安全
//
// Args:
//
//	path (str): 本地文件路径，如：`C:/Projs/WeChatRobot/README.MD` 或 `https://raw.githubusercontent.com/lich0821/WeChatFerry/master/README.MD`
//	receiver (str): 消息接收人，wxid 或者 roomid
//
// Returns:
//
//	int: 0 为成功，其他失败
func (c *Client) SendFile(path string, receiver string) int32 {
	req := genFunReq(Functions_FUNC_SEND_FILE)
	req.Msg = &Request_File{
		File: &PathMsg{
			Path:     path,
			Receiver: receiver,
		},
	}
	recv := c.Call(req.build())
	return recv.GetStatus()
}

// 发送 XML
//
// Args:
//
//	receiver (str): 消息接收人，wxid 或者 roomid
//	xml (str): xml 内容
//	type (int): xml 类型，如：0x21 为小程序
//	path (str): 封面图片路径
//
// Returns:
//
//	int: 0 为成功，其他失败
func (c *Client) SendXml(path, content, receiver string, Type int32) int32 {
	req := genFunReq(Functions_FUNC_SEND_XML)
	req.Msg = &Request_Xml{
		Xml: &XmlMsg{
			Receiver: receiver,
			Content:  content,
			Path:     path,
			Type:     Type,
		},
	}
	recv := c.Call(req.build())
	return recv.GetStatus()
}

// 发送表情
//
// Args:
//
//	path (str): 本地表情路径，如：`C:/Projs/WeChatRobot/emo.gif`
//	receiver (str): 消息接收人，wxid 或者 roomid
//
// Returns:
//
//	int: 0 为成功，其他失败
func (c *Client) SendEmotion(path, receiver string) int32 {
	req := genFunReq(Functions_FUNC_SEND_EMOTION)
	req.Msg = &Request_File{
		File: &PathMsg{
			Path:     path,
			Receiver: receiver,
		},
	}
	recv := c.Call(req.build())
	return recv.GetStatus()
}

// 接受好友申请
//
// Args:
//
//	v3 (str): 加密用户名 (好友申请消息里 v3 开头的字符串)
//	v4 (str): Ticket (好友申请消息里 v4 开头的字符串)
//	scene: 申请方式 (好友申请消息里的 scene); 为了兼容旧接口，默认为扫码添加 (30)
//
// Returns:
//
//	int: 1 为成功，其他失败
func (c *Client) AcceptNewFriend(v3, v4 string, scene int32) int32 {
	req := genFunReq(Functions_FUNC_ACCEPT_FRIEND)
	req.Msg = &Request_V{
		V: &Verification{
			V3:    v3,
			V4:    v4,
			Scene: scene,
		},
	}
	recv := c.Call(req.build())
	return recv.GetStatus()
}

// 接收转账
//
// Args:
//
//	wxid (str): 转账消息里的发送人 wxid
//	transferid (str): 转账消息里的 transferid
//	transactionid (str): 转账消息里的 transactionid
//
// Returns:
//
//	int: 1 为成功，其他失败
func (c *Client) ReceiveTransfer(wxid, tfid, taid string) int32 {
	req := genFunReq(Functions_FUNC_RECV_TRANSFER)
	req.Msg = &Request_Tf{
		Tf: &Transfer{
			Wxid: wxid,
			Tfid: tfid,
			Taid: taid,
		},
	}
	recv := c.Call(req.build())
	return recv.GetStatus()
}

// 刷新朋友圈
//
// Args:
//
//	id (int): 开始 id，0 为最新页
//
// Returns:
//
//	int: 1 为成功，其他失败
func (c *Client) RefreshPyq(id uint64) int32 {
	req := genFunReq(Functions_FUNC_REFRESH_PYQ)
	req.Msg = &Request_Ui64{
		Ui64: id,
	}
	recv := c.Call(req.build())
	return recv.GetStatus()
}

// 添加群成员
//
// Args:
//
//	roomid (str): 待加群的 id
//	wxids (str): 要加到群里的 wxid，多个用逗号分隔
//
// Returns:
//
//	int: 1 为成功，其他失败
func (c *Client) AddChatRoomMembers(roomId string, wxIds []string) int32 {
	req := genFunReq(Functions_FUNC_ADD_ROOM_MEMBERS)
	req.Msg = &Request_M{
		M: &AddMembers{
			Roomid: roomId,
			Wxids:  strings.Join(wxIds, ","),
		},
	}
	recv := c.Call(req.build())
	return recv.GetStatus()
}

// 删除群成员
//
// Args:
//
//	roomid (str): 群的 id
//	wxids (str): 要删除成员的 wxid，多个用逗号分隔
//
// Returns:
//
//	int: 1 为成功，其他失败
func (c *Client) DelChatRoomMembers(roomId string, wxIds []string) int32 {
	req := genFunReq(Functions_FUNC_DEL_ROOM_MEMBERS)
	req.Msg = &Request_M{
		M: &AddMembers{
			Roomid: roomId,
			Wxids:  strings.Join(wxIds, ","),
		},
	}
	recv := c.Call(req.build())
	return recv.GetStatus()
}

// 获取群成员列表
//
// Args:
//
//	roomid (str): 群的 id
//
// Returns:
//
//	Dict: 群成员列表: {wxid1: 昵称1, wxid2: 昵称2, ...}
func (c *Client) GetChatRoomMembers(roomId string) []*RpcContact {
	members := []*RpcContact{}
	// get user data
	userRds := c.ExecDbQuery("MicroMsg.db", "SELECT UserName, NickName FROM Contact;")
	userMap := map[string]string{}
	for _, user := range userRds {
		wxid := string(user.Fields[0].Content)
		userMap[wxid] = string(user.Fields[1].Content)
	}
	// get room data
	roomRds := c.ExecDbQuery("MicroMsg.db", "SELECT RoomData FROM ChatRoom WHERE ChatRoomName = '"+roomId+"';")
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
//
// Args:
//
//	wxid (str): wxid
//	roomid (str): 群的 id
//
// Returns:
//
//	str: 群成员昵称
func (c *Client) GetNickNameInChatRoom(wxid, roomId string) string {
	// get user data
	nickName := ""
	userRds := c.ExecDbQuery("MicroMsg.db", "SELECT NickName FROM Contact WHERE UserName = '"+wxid+"';")
	if len(userRds) > 0 && len(userRds[0].Fields) > 0 {
		nickName = string(userRds[0].Fields[0].Content)
	}
	// get room data
	roomRds := c.ExecDbQuery("MicroMsg.db", "SELECT RoomData FROM ChatRoom WHERE ChatRoomName = '"+roomId+"';")
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
//
// Args:
//
//	src (str): 加密的图片路径
//	dst (str): 解密的图片路径
//
// Returns:
//
//	bool: 是否成功
func (c *Client) DecryptImage(src, dst string) int32 {
	req := genFunReq(Functions_FUNC_DECRYPT_IMAGE)
	req.Msg = &Request_Dec{
		Dec: &DecPath{
			Src: src,
			Dst: dst,
		},
	}
	recv := c.Call(req.build())
	return recv.GetStatus()
}

// 异步处理接的消息
func (c *Client) OnReceivingMsg(addr string, f func(msg *WxMsg)) error {
	socket, err := pair1.NewSocket()
	if err != nil {
		return err
	}
	socket.SetOption(mangos.OptionRecvDeadline, 2000)
	socket.SetOption(mangos.OptionSendDeadline, 2000)
	if err = socket.Dial(addr); err != nil {
		return err
	}
	defer socket.Close()
	for c.IsReceivingMsg {
		if recv, err := socket.Recv(); err == nil {
			resp := &Response{}
			proto.Unmarshal(recv, resp)
			go f(resp.GetWxmsg())
		} else {
			return err
		}
	}
	return err
}

// 允许接收消息
func (c *Client) EnableReceivingMsg() int32 {
	req := genFunReq(Functions_FUNC_ENABLE_RECV_TXT)
	req.Msg = &Request_Flag{
		Flag: true,
	}
	recv := c.Call(req.build())
	c.IsReceivingMsg = true
	return recv.GetStatus()
}

// 停止接收消息
func (c *Client) DisableReceivingMsg() int32 {
	req := genFunReq(Functions_FUNC_DISABLE_RECV_TXT)
	recv := c.Call(req.build())
	c.IsReceivingMsg = false
	return recv.GetStatus()
}
