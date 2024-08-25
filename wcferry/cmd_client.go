package wcferry

import (
	"errors"
	"path"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
)

type CmdClient struct {
	*pbSocket // RPC 客户端
}

// 关闭 RPC 连接
// return error 错误信息
func (c *CmdClient) Destroy() error {
	return c.close()
}

// 检查登录状态
// return bool 是否已登录
func (c *CmdClient) IsLogin() bool {
	req := &Request{Func: Functions_FUNC_IS_LOGIN}
	recv := c.call(req)
	return recv.GetStatus() == 1
}

// 刷新登录二维码【not implemented】
// return string 登录二维码
func (c *CmdClient) RefreshQrcode() string {
	req := &Request{Func: Functions_FUNC_REFRESH_QRCODE}
	recv := c.call(req)
	return recv.GetStr()
}

// 获取登录账号wxid
// return string 登录账号wxid
func (c *CmdClient) GetSelfWxid() string {
	req := &Request{Func: Functions_FUNC_GET_SELF_WXID}
	recv := c.call(req)
	return recv.GetStr()
}

// 获取登录账号个人信息
// return *UserInfo 登录账号个人信息
func (c *CmdClient) GetSelfInfo() *UserInfo {
	req := &Request{Func: Functions_FUNC_GET_USER_INFO}
	recv := c.call(req)
	return recv.GetUi()
}

// 获取所有消息类型
// return map[int32]string 所有消息类型
func (c *CmdClient) GetMsgTypes() map[int32]string {
	req := &Request{Func: Functions_FUNC_GET_MSG_TYPES}
	recv := c.call(req)
	return recv.GetTypes().GetTypes()
}

// 获取所有数据库
// return []string 所有数据库名
func (c *CmdClient) GetDbNames() []string {
	req := &Request{Func: Functions_FUNC_GET_DB_NAMES}
	recv := c.call(req)
	return recv.GetDbs().Names
}

// 获取数据库中所有表
// param db string 数据库名
// return []*DbTable `db` 下的所有表名及对应建表语句
func (c *CmdClient) GetDbTables(db string) []*DbTable {
	req := &Request{Func: Functions_FUNC_GET_DB_TABLES}
	req.Msg = &Request_Str{
		Str: db,
	}
	recv := c.call(req)
	return recv.GetTables().GetTables()
}

// 执行 SQL 查询，如果数据量大注意分页
// param db string 要查询的数据库
// param sql string 要执行的 SQL
// return []map[string]any 查询结果
func (c *CmdClient) DbSqlQuery(db, sql string) []map[string]any {
	req := &Request{Func: Functions_FUNC_EXEC_DB_QUERY}
	req.Msg = &Request_Query{
		Query: &DbQuery{
			Db:  db,
			Sql: sql,
		},
	}
	recv := c.call(req)
	rows := []map[string]any{}
	for _, row := range recv.GetRows().GetRows() {
		fields := map[string]any{}
		for _, field := range row.Fields {
			fields[field.Column] = ParseDbField(field)
		}
		rows = append(rows, fields)
	}
	return rows
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

// 获取群成员列表
// param roomid string 群的 id
// return []*RpcContact 群成员列表
func (c *CmdClient) GetChatRoomMembers(roomid string) []*RpcContact {
	members := []*RpcContact{}
	// get room data
	roomList := c.DbSqlQuery("MicroMsg.db", "SELECT RoomData FROM ChatRoom WHERE ChatRoomName = '"+roomid+"';")
	if len(roomList) == 0 || len(roomList[0]) == 0 || roomList[0]["RoomData"] == nil {
		return members
	}
	roomData := &RoomData{}
	if err := proto.Unmarshal(roomList[0]["RoomData"].([]byte), roomData); err != nil {
		return members
	}
	// get user data
	userList := c.DbSqlQuery("MicroMsg.db", "SELECT UserName, NickName FROM Contact;")
	userMap := map[string]string{}
	for _, user := range userList {
		wxid := user["UserName"].(string)
		userMap[wxid] = user["NickName"].(string)
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
func (c *CmdClient) GetAliasInChatRoom(wxid, roomid string) string {
	// get user data
	nickName := ""
	userList := c.DbSqlQuery("MicroMsg.db", "SELECT NickName FROM Contact WHERE UserName = '"+wxid+"';")
	if len(userList) > 0 && len(userList[0]) > 0 {
		if userList[0]["NickName"] != nil {
			nickName = userList[0]["NickName"].(string)
		}
	}
	// get room data
	roomList := c.DbSqlQuery("MicroMsg.db", "SELECT RoomData FROM ChatRoom WHERE ChatRoomName = '"+roomid+"';")
	if len(roomList) == 0 || len(roomList[0]) == 0 || roomList[0]["RoomData"] == nil {
		return nickName
	}
	roomData := &RoomData{}
	if err := proto.Unmarshal(roomList[0]["RoomData"].([]byte), roomData); err != nil {
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

// 邀请群成员
// param roomid string 群的 id
// param wxids string 要邀请成员的 wxid, 多个用逗号`,`分隔
// return int32 1 为成功，其他失败
func (c *CmdClient) InviteChatroomMembers(roomid, wxids string) int32 {
	req := &Request{Func: Functions_FUNC_INV_ROOM_MEMBERS}
	req.Msg = &Request_M{
		M: &MemberMgmt{
			Roomid: roomid,
			Wxids:  strings.ReplaceAll(wxids, " ", ""),
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 添加群成员
// param roomid string 待加群的 id
// param wxids string 要加到群里的 wxid，多个用逗号分隔
// return int32 1 为成功，其他失败
func (c *CmdClient) AddChatRoomMembers(roomid, wxIds string) int32 {
	req := &Request{Func: Functions_FUNC_ADD_ROOM_MEMBERS}
	req.Msg = &Request_M{
		M: &MemberMgmt{
			Roomid: roomid,
			Wxids:  wxIds,
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 删除群成员
// param roomid string 群的 id
// param wxids string 要删除成员的 wxid，多个用逗号分隔
// return int32 1 为成功，其他失败
func (c *CmdClient) DelChatRoomMembers(roomid, wxIds string) int32 {
	req := &Request{Func: Functions_FUNC_DEL_ROOM_MEMBERS}
	req.Msg = &Request_M{
		M: &MemberMgmt{
			Roomid: roomid,
			Wxids:  wxIds,
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 撤回消息
// param msgid (uint64): 消息 id
// return int: 1 为成功，其他失败
func (c *CmdClient) RevokeMsg(msgid uint64) int32 {
	req := &Request{Func: Functions_FUNC_REVOKE_MSG}
	req.Msg = &Request_Ui64{
		Ui64: msgid,
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 转发消息
// param msgid (uint64): 消息 id
// param receiver string 消息接收人，wxid 或者 roomid
// return int: 1 为成功，其他失败
func (c *CmdClient) ForwardMsg(msgid uint64, receiver string) int32 {
	req := &Request{Func: Functions_FUNC_FORWARD_MSG}
	req.Msg = &Request_Fm{
		Fm: &ForwardMsg{
			Id:       msgid,
			Receiver: receiver,
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 发送文本消息
// param msg string 要发送的消息，\n使用 `\\\\n` （单杠）；如果 @ 人的话，需要带上跟 `aters` 里数量相同的 @
// param receiver string 消息接收人，wxid 或者 roomid
// param aters string 要 @ 的 wxid，多个用逗号分隔；`@所有人` 只需要 `notify@all`
// return int32 0 为成功，其他失败
func (c *CmdClient) SendTxt(msg, receiver, aters string) int32 {
	req := &Request{Func: Functions_FUNC_SEND_TXT}
	req.Msg = &Request_Txt{
		Txt: &TextMsg{
			Msg:      msg,
			Receiver: receiver,
			Aters:    aters,
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 发送图片，非线程安全
// param path string 图片路径，如：`C:/Projs/WeChatRobot/TEQuant.jpeg`
// param receiver string 消息接收人，wxid 或者 roomid
// return int32 0 为成功，其他失败
func (c *CmdClient) SendImg(path, receiver string) int32 {
	if tmp := DownloadFile(path); tmp != "" {
		path = tmp
	}
	req := &Request{Func: Functions_FUNC_SEND_IMG}
	req.Msg = &Request_File{
		File: &PathMsg{
			Path:     path,
			Receiver: receiver,
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 发送文件，非线程安全
// param path string 本地文件路径，如：`C:/Projs/WeChatRobot/README.MD`
// param receiver string 消息接收人，wxid 或者 roomid
// return int32 0 为成功，其他失败
func (c *CmdClient) SendFile(path, receiver string) int32 {
	if tmp := DownloadFile(path); tmp != "" {
		path = tmp
	}
	req := &Request{Func: Functions_FUNC_SEND_FILE}
	req.Msg = &Request_File{
		File: &PathMsg{
			Path:     path,
			Receiver: receiver,
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 发送 XML（暂不支持）
// param path string 封面图片路径
// param content string xml 内容
// param receiver string 消息接收人，wxid 或者 roomid
// param Type int32 xml 类型，如：0x21 为小程序
// return int32 0 为成功，其他失败
func (c *CmdClient) SendXml(path, content, receiver string, Type int32) int32 {
	if tmp := DownloadFile(path); tmp != "" {
		path = tmp
	}
	req := &Request{Func: Functions_FUNC_SEND_XML}
	req.Msg = &Request_Xml{
		Xml: &XmlMsg{
			Receiver: receiver,
			Content:  content,
			Path:     path,
			Type:     Type,
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 发送表情（暂不支持）
// param path string 本地表情路径，如：`C:/Projs/WeChatRobot/emo.gif`
// param receiver string 消息接收人，wxid 或者 roomid
// return int32 0 为成功，其他失败
func (c *CmdClient) SendEmotion(path, receiver string) int32 {
	if tmp := DownloadFile(path); tmp != "" {
		path = tmp
	}
	req := &Request{Func: Functions_FUNC_SEND_EMOTION}
	req.Msg = &Request_File{
		File: &PathMsg{
			Path:     path,
			Receiver: receiver,
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 发送富文本消息
// 卡片样式：
// |-------------------------------------|
// |title, 最长两行                       |
// |(长标题, 标题短的话这行没有)           |
// |digest, 最多三行，会占位     |--------|
// |digest, 最多三行，会占位     |thumburl|
// |digest, 最多三行，会占位     |--------|
// |(account logo) name                  |
// |-------------------------------------|
// param name string 左下显示的名字
// param account string 填公众号 id 可以显示对应的头像（gh_ 开头的）
// param title string 标题，最多两行
// param digest string 摘要，三行
// param url string 点击后跳转的链接
// param thumburl string 缩略图的链接
// param receiver string 接收人, wxid 或者 roomid
// return int32 0 为成功，其他失败
func (c *CmdClient) SendRichText(name, account, title, digest, url, thumburl, receiver string) int32 {
	req := &Request{Func: Functions_FUNC_SEND_RICH_TXT}
	req.Msg = &Request_Rt{
		Rt: &RichText{
			Name:     name,
			Account:  account,
			Title:    title,
			Digest:   digest,
			Url:      url,
			Thumburl: thumburl,
			Receiver: receiver,
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 拍一拍群友
// param roomid string 群 id
// param wxid string 要拍的群友的 wxid
// return int32 1 为成功，其他失败
func (c *CmdClient) SendPatMsg(roomid, wxid string) int32 {
	req := &Request{Func: Functions_FUNC_SEND_PAT_MSG}
	req.Msg = &Request_Pm{
		Pm: &PatMsg{
			Roomid: roomid,
			Wxid:   wxid,
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 获取语音消息并转成 MP3
// param msgid 语音消息 id
// param dir MP3 保存目录（目录不存在会出错）
// return string 成功返回存储路径；空字符串为失败
func (c *CmdClient) GetAudioMsg(msgid uint64, dir string) string {
	req := &Request{Func: Functions_FUNC_GET_AUDIO_MSG}
	req.Msg = &Request_Am{
		Am: &AudioMsg{
			Id:  msgid,
			Dir: dir,
		},
	}
	recv := c.call(req)
	return recv.GetStr()
}

// 获取语音消息并转成 MP3
// param msgid 语音消息 id
// param dir MP3 保存目录（目录不存在会出错）
// param timeout 超时重试次数（每次重试间隔1秒）
// return string 成功返回存储路径；空字符串为失败
func (c *CmdClient) GetAudioMsgTimeout(msgid uint64, dir string, timeout int) (string, error) {
	cnt := 0
	for cnt <= timeout {
		if path := c.GetAudioMsg(msgid, dir); path != "" {
			return path, nil
		}
		time.Sleep(1 * time.Second)
		cnt++
	}
	// 超时
	return "", errors.New("timeout")
}

// 获取 OCR 结果
// 鸡肋，需要图片能自动下载；通过下载接口下载的图片无法识别
// param extra string 待识别的图片路径，消息里的 extra
// return string OCR 结果
// return int32 状态码，0 为成功，其他失败
func (c *CmdClient) GetOcrResult(extra string) (string, int32) {
	req := &Request{Func: Functions_FUNC_EXEC_OCR}
	req.Msg = &Request_Str{
		Str: extra,
	}
	recv := c.call(req)
	ocr := recv.GetOcr()
	return ocr.GetResult(), ocr.GetStatus()
}

// 获取 OCR 结果，带有超时时间
// param extra string 待识别的图片路径，消息里的 extra
// param timeout int 超时时间
// return string OCR 结果
func (c *CmdClient) GetOcrResultTimeout(extra string, timeout int) (string, error) {
	cnt := 0
	for cnt <= timeout {
		result, status := c.GetOcrResult(extra)
		if status == 0 {
			return result, nil
		}
		time.Sleep(1 * time.Second)
		cnt++
	}
	// 超时
	return "", errors.New("timeout")
}

// 下载图片
// param msgid uint64 消息 id
// param extra string 消息中的 extra
// param dir string 存放图片的目录（目录不存在会出错）
// param timeout int 超时重试次数（每次重试间隔1秒）
// return string 成功返回存储路径
func (c *CmdClient) DownloadImage(msgid uint64, extra, dir string, timeout int) (string, error) {
	if c.DownloadAttach(msgid, "", extra) != 0 {
		time.Sleep(1 * time.Second)
		if c.DownloadAttach(msgid, "", extra) != 0 {
			return "", errors.New("download failed")
		}
	}
	cnt := 0
	for cnt <= timeout {
		if path := c.DecryptImage(extra, dir); path != "" {
			return path, nil
		}
		time.Sleep(1 * time.Second)
		cnt++
	}
	// 超时
	return "", errors.New("timeout")
}

// 下载附件
// param msgid string 消息 id
// param thumb string 消息中的 thumb
// param extra string 消息中的 extra
// return int32 0 为成功，其他失败
func (c *CmdClient) DownloadAttach(msgid uint64, thumb, extra string) int32 {
	req := &Request{Func: Functions_FUNC_DOWNLOAD_ATTACH}
	req.Msg = &Request_Att{
		Att: &AttachMsg{
			Id:    msgid,
			Thumb: thumb,
			Extra: extra,
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 解密图片
// 此方法别直接调用，下载图片使用 `DownloadImage` 方法
// param src string 加密的图片路径
// param dir string 保存图片的目录
// return str 解密图片的保存路径
func (c *CmdClient) DecryptImage(src, dir string) string {
	if dir == "" {
		dir = path.Dir(src)
	}
	req := &Request{Func: Functions_FUNC_DECRYPT_IMAGE}
	req.Msg = &Request_Dec{
		Dec: &DecPath{
			Src: src,
			Dst: dir,
		},
	}
	recv := c.call(req)
	return recv.GetStr()
}

// 获取完整通讯录
// return []*RpcContact 完整通讯录
func (c *CmdClient) GetContacts() []*RpcContact {
	req := &Request{Func: Functions_FUNC_GET_CONTACTS}
	recv := c.call(req)
	return recv.GetContacts().GetContacts()
}

// 获取好友列表
// return []*RpcContact 好友列表
func (c *CmdClient) GetFriends() []*RpcContact {
	friends := []*RpcContact{}
	for _, cnt := range c.GetContacts() {
		if ContactType(cnt.Wxid) == "好友" {
			friends = append(friends, cnt)
		}
	}
	return friends
}

// 通过 wxid 查询微信号昵称等信息【not implemented】
// param wxid (str): 联系人 wxid
// return *RpcContact
func (c *CmdClient) GetInfoByWxid(wxid string) *RpcContact {
	req := &Request{Func: Functions_FUNC_GET_CONTACT_INFO}
	req.Msg = &Request_Str{
		Str: wxid,
	}
	recv := c.call(req)
	contacts := recv.GetContacts()
	if contacts != nil {
		contacts := contacts.GetContacts()
		if contacts != nil {
			return contacts[0]
		}
	}
	return nil
}

// 刷新朋友圈
// param id int32 开始 id，0 为最新页
// return int32 1 为成功，其他失败
func (c *CmdClient) RefreshPyq(id uint64) int32 {
	req := &Request{Func: Functions_FUNC_REFRESH_PYQ}
	req.Msg = &Request_Ui64{
		Ui64: id,
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 接受好友申请【not implemented】
// param v3 string 加密用户名 (好友申请消息里 v3 开头的字符串)
// param v4 string Ticket (好友申请消息里 v4 开头的字符串)
// param scene int32 申请方式 (好友申请消息里的 scene); 为了兼容旧接口，默认为扫码添加 (30)
// return int32 1 为成功，其他失败
func (c *CmdClient) AcceptNewFriend(v3, v4 string, scene int32) int32 {
	req := &Request{Func: Functions_FUNC_ACCEPT_FRIEND}
	req.Msg = &Request_V{
		V: &Verification{
			V3:    v3,
			V4:    v4,
			Scene: scene,
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 接收好友转账
// param wxid string 转账消息里的发送人 wxid
// param transferid string 转账消息里的 transferid
// param transactionid string 转账消息里的 transactionid
// return int32 1 为成功，其他失败
func (c *CmdClient) ReceiveTransfer(wxid, tfid, taid string) int32 {
	req := &Request{Func: Functions_FUNC_RECV_TRANSFER}
	req.Msg = &Request_Tf{
		Tf: &Transfer{
			Wxid: wxid,
			Tfid: tfid,
			Taid: taid,
		},
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 开启消息接收服务
// param pyq bool 是否接收朋友圈消息
// return int32 0 为成功，其他失败
func (c *CmdClient) EnableMsgReciver(pyq bool) int32 {
	req := &Request{Func: Functions_FUNC_ENABLE_RECV_TXT}
	req.Msg = &Request_Flag{
		Flag: pyq,
	}
	recv := c.call(req)
	return recv.GetStatus()
}

// 停止消息接收服务
// return int32 0 为成功，其他失败
func (c *CmdClient) DisableMsgReciver() int32 {
	req := &Request{Func: Functions_FUNC_DISABLE_RECV_TXT}
	recv := c.call(req)
	return recv.GetStatus()
}
