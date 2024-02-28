import { httpRequest } from "./request";

export const RobotApi = {
    /**
     * @summary 添加群聊
     * @param {ChatroomCreateParam} body 添加群聊参数
     * @param {*} [options] Override http request option.
     */
    chatroomCreate(body: ChatroomCreateParam, options: RequestInit = {}): Promise<number> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/create', options);
    },
    /**
     * @summary 删除群聊
     * @param {ChatroomDeleteParam} body 删除群聊参数
     * @param {*} [options] Override http request option.
     */
    chatroomDelete(body: ChatroomDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/delete', options);
    },
    /**
     * @summary 获取群聊
     * @param {ChatroomFetchParam} body 获取群聊参数
     * @param {*} [options] Override http request option.
     */
    chatroomDetail(body: ChatroomFetchParam, options: RequestInit = {}): Promise<TablesChatroom> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/detail', options);
    },
    /**
     * @summary 群聊列表
     * @param {ChatroomFetchAllParam} body 获取群聊列表参数
     * @param {*} [options] Override http request option.
     */
    chatroomList(body: ChatroomFetchAllParam, options: RequestInit = {}): Promise<TablesChatroom[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/list', options);
    },
    /**
     * @summary 修改群聊
     * @param {ChatroomUpdateParam} body 修改群聊参数
     * @param {*} [options] Override http request option.
     */
    chatroomUpdate(body: ChatroomUpdateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/update', options);
    },
    /**
     * @summary 添加关键字
     * @param {KeywordCreateParam} body 添加关键字参数
     * @param {*} [options] Override http request option.
     */
    keywordCreate(body: KeywordCreateParam, options: RequestInit = {}): Promise<number> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/create', options);
    },
    /**
     * @summary 删除关键字
     * @param {KeywordDeleteParam} body 删除关键字参数
     * @param {*} [options] Override http request option.
     */
    keywordDelete(body: KeywordDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/delete', options);
    },
    /**
     * @summary 获取关键字
     * @param {KeywordFetchParam} body 获取关键字参数
     * @param {*} [options] Override http request option.
     */
    keywordDetail(body: KeywordFetchParam, options: RequestInit = {}): Promise<TablesKeyword> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/detail', options);
    },
    /**
     * @summary 关键字列表
     * @param {KeywordFetchAllParam} body 获取关键字列表参数
     * @param {*} [options] Override http request option.
     */
    keywordList(body: KeywordFetchAllParam, options: RequestInit = {}): Promise<TablesKeyword[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/list', options);
    },
    /**
     * @summary 修改关键字
     * @param {KeywordUpdateParam} body 修改关键字参数
     * @param {*} [options] Override http request option.
     */
    keywordUpdate(body: KeywordUpdateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/update', options);
    },
    /**
     * @summary 添加模型
     * @param {LlmodelCreateParam} body 添加模型参数
     * @param {*} [options] Override http request option.
     */
    llmodelCreate(body: LlmodelCreateParam, options: RequestInit = {}): Promise<number> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/llmodel/create', options);
    },
    /**
     * @summary 删除模型
     * @param {LlmodelDeleteParam} body 删除模型参数
     * @param {*} [options] Override http request option.
     */
    llmodelDelete(body: LlmodelDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/llmodel/delete', options);
    },
    /**
     * @summary 获取模型
     * @param {LlmodelFetchParam} body 获取模型参数
     * @param {*} [options] Override http request option.
     */
    llmodelDetail(body: LlmodelFetchParam, options: RequestInit = {}): Promise<TablesLLModel> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/llmodel/detail', options);
    },
    /**
     * @summary 模型列表
     * @param {LlmodelFetchAllParam} body 获取模型列表参数
     * @param {*} [options] Override http request option.
     */
    llmodelList(body: LlmodelFetchAllParam, options: RequestInit = {}): Promise<TablesLLModel[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/llmodel/list', options);
    },
    /**
     * @summary 修改模型
     * @param {LlmodelUpdateParam} body 修改模型参数
     * @param {*} [options] Override http request option.
     */
    llmodelUpdate(body: LlmodelUpdateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/llmodel/update', options);
    },
    /**
     * @summary 添加配置
     * @param {ProfileCreateParam} body 添加配置参数
     * @param {*} [options] Override http request option.
     */
    profileCreate(body: ProfileCreateParam, options: RequestInit = {}): Promise<number> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/create', options);
    },
    /**
     * @summary 删除配置
     * @param {ProfileDeleteParam} body 删除配置参数
     * @param {*} [options] Override http request option.
     */
    profileDelete(body: ProfileDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/delete', options);
    },
    /**
     * @summary 获取配置
     * @param {ProfileFetchParam} body 获取配置参数
     * @param {*} [options] Override http request option.
     */
    profileDetail(body: ProfileFetchParam, options: RequestInit = {}): Promise<TablesProfile> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/detail', options);
    },
    /**
     * @summary 配置列表
     * @param {ProfileFetchAllParam} body 获取配置列表参数
     * @param {*} [options] Override http request option.
     */
    profileList(body: ProfileFetchAllParam, options: RequestInit = {}): Promise<TablesProfile[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/list', options);
    },
    /**
     * @summary 修改配置
     * @param {ProfileUpdateParam} body 修改配置参数
     * @param {*} [options] Override http request option.
     */
    profileUpdate(body: ProfileUpdateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/update', options);
    },
};

export interface ChatroomCreateParam {
    // 入群口令
    join_argot: string;
    // 等级
    level: number;
    // 群聊名称
    name: string;
    // 备注
    remark: string;
    // 防撤回消息
    revoke_msg: string;
    // 群聊 id
    roomid: string;
    // 欢迎消息
    welcome_msg: string;
}

export interface ChatroomDeleteParam {
    // 群聊 id
    roomid: string;
}

export interface ChatroomFetchAllParam {
    // 等级
    level: number;
}

export interface ChatroomFetchParam {
    // 群聊 id
    roomid: string;
}

export interface ChatroomUpdateParam {
    // 入群口令
    join_argot: string;
    // 等级
    level: number;
    // 群聊名称
    name: string;
    // 备注
    remark: string;
    // 防撤回消息
    revoke_msg: string;
    // 群聊 id
    roomid: string;
    // 欢迎消息
    welcome_msg: string;
}

export interface KeywordCreateParam {
    // 等级
    level: number;
    // 词语或短语
    phrase: string;
    // 群聊 id
    roomid: string;
}

export interface KeywordDeleteParam {
    // 词语或短语
    phrase: string;
    // 群聊 id
    roomid: string;
}

export interface KeywordFetchAllParam {
    // 等级
    level: number;
    // 群聊 id
    roomid: string;
}

export interface KeywordFetchParam {
    // 词语或短语
    phrase: string;
    // 群聊 id
    roomid: string;
}

export interface KeywordUpdateParam {
    // 等级
    level: number;
    // 词语或短语
    phrase: string;
    // 主键
    rd: number;
    // 群聊 id
    roomid: string;
}

export interface LlmodelCreateParam {
    // 接口地址
    endpoint: string;
    // 模型家族
    family: string;
    // 模型 Id
    mid: string;
    // 模型名称
    model: string;
    // 服务商名称
    provider: string;
    // 密钥
    secret: string;
}

export interface LlmodelDeleteParam {
    // 模型 Id
    mid: string;
}

export interface LlmodelFetchAllParam {
    // 模型家族
    family: string;
    // 模型名称
    model: string;
    // 服务商名称
    provider: string;
}

export interface LlmodelFetchParam {
    // 模型 Id
    mid: string;
}

export interface LlmodelUpdateParam {
    // 接口地址
    endpoint: string;
    // 模型家族
    family: string;
    // 模型 Id
    mid: string;
    // 模型名称
    model: string;
    // 服务商名称
    provider: string;
    // 密钥
    secret: string;
}

export interface ProfileCreateParam {
    // 唤醒词
    ai_argot: string;
    // 会话模型
    ai_model: string;
    // 等级
    level: number;
    // 备注
    remark: string;
    // 群聊 id
    roomid: string;
    // 微信 id
    wxid: string;
}

export interface ProfileDeleteParam {
    // 群聊 id
    roomid: string;
    // 微信 id
    wxid: string;
}

export interface ProfileFetchAllParam {
    // 群聊 id
    roomid: string;
    // 微信 id
    wxid: string;
}

export interface ProfileFetchParam {
    // 群聊 id
    roomid: string;
    // 微信 id
    wxid: string;
}

export interface ProfileUpdateParam {
    // 唤醒词
    ai_argot: string;
    // 会话模型
    ai_model: string;
    // 等级
    level: number;
    // 备注
    remark: string;
    // 群聊 id
    roomid: string;
    // 微信 id
    wxid: string;
}

export interface TablesChatroom {
    // 创建时间戳
    created_at: number;
    // 入群口令
    join_argot: string;
    // 等级
    level: number;
    // 群聊名称
    name: string;
    // 主键
    rd: number;
    // 备注
    remark: string;
    // 防撤回消息
    revoke_msg: string;
    // 群聊 id
    roomid: string;
    // 最后更新时间戳
    updated_at: number;
    // 欢迎消息
    welcome_msg: string;
}

export interface TablesKeyword {
    // 创建时间戳
    created_at: number;
    // 优先级等级
    level: number;
    // 词语或短语
    phrase: string;
    // 主键
    rd: number;
    // 群聊 id
    roomid: string;
    // 最后更新时间戳
    updated_at: number;
}

export interface TablesLLModel {
    // 创建时间戳
    created_at: number;
    // 接口地址，仅 google 和 openai 支持自定义
    endpoint: string;
    // 模型家族，用于生成模型切换指令
    family: string;
    // 模型 Id
    mid: string;
    // 模型，必须和服务商提供的值对应
    model: string;
    // 服务商，支持 google, openai, xunfei
    provider: string;
    // 主键
    rd: number;
    // 密钥，google 和 openai 填写 KEY，xunfei 填写 APP-ID,API-KEY,API-SECRET
    secret: string;
    // 最后更新时间戳
    updated_at: number;
}

export interface TablesProfile {
    // 唤醒词
    ai_argot: string;
    // 会话模型
    ai_model: string;
    // 拉黑截止时间
    ban_expire: number;
    // 创建时间戳
    created_at: number;
    // 等级
    level: number;
    // 主键
    rd: number;
    // 备注
    remark: string;
    // 群聊 id
    roomid: string;
    // 最后更新时间戳
    updated_at: number;
    // 微信 id
    wxid: string;
}
