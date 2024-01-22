package args

type Member struct {
	AiArgot string // 唤醒词
	AiModel string // 会话模型
	Level   int    // 等级 [0:未注册, 1:已禁用 9:管理员]
	Remark  string // 备注信息
}

type ChatMember struct {
	Member // 继承用户基础信息
}

type ChatRoom struct {
	Argot   string                 // 群标记，用于生成加群指令
	Level   int                    // 等级 [0:未注册, 1:已禁用]
	Member  map[string]*ChatMember // 群成员列表
	Name    string                 // 群名称，在指令说明中使用
	Welcome string                 // 欢迎词
}

// Member Config

func GetMember(uid string) *Member {

	if _, ok := Usr.Member[uid]; !ok {
		Usr.Member[uid] = &Member{}
	}

	return Usr.Member[uid]

}

func (user *Member) GetModel() *LLModel {

	llmc := LLM.Models[user.AiModel]

	if llmc == nil {
		llmc = LLM.Models[LLM.Default]
	}

	if llmc == nil {
		for _, v := range LLM.Models {
			return v
		}
	}

	return llmc

}

// ChatRoom Config

func GetChatRoom(rid string) *ChatRoom {

	if _, ok := Usr.ChatRoom[rid]; !ok {
		Usr.ChatRoom[rid] = &ChatRoom{}
	}

	return Usr.ChatRoom[rid]

}

func (ChatRoom *ChatRoom) GetMember(uid string) *ChatMember {

	if _, ok := ChatRoom.Member[uid]; !ok {
		ChatRoom.Member[uid] = &ChatMember{*GetMember(uid)}
	}

	return ChatRoom.Member[uid]

}
