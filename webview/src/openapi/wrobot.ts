import { httpRequest } from "./request";

export class RobotApi {
    /**
     * @summary 添加群聊
     * @param {ChatroomCreateParam} body 添加群聊参数
     * @param {*} [options] Override http request option.
     */
    public chatroomCreate(body: ChatroomCreateParam, options: RequestInit = {}): Promise<number> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/create', options);
    }
    /**
     * @summary 删除群聊
     * @param {ChatroomDeleteParam} body 删除群聊参数
     * @param {*} [options] Override http request option.
     */
    public chatroomDelete(body: ChatroomDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/delete', options);
    }
    /**
     * @summary 获取群聊
     * @param {ChatroomFetchParam} body 获取群聊参数
     * @param {*} [options] Override http request option.
     */
    public chatroomDetail(body: ChatroomFetchParam, options: RequestInit = {}): Promise<TablesChatroom> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/detail', options);
    }
    /**
     * @summary 群聊列表
     * @param {ChatroomFetchAllParam} body 获取群聊列表参数
     * @param {*} [options] Override http request option.
     */
    public chatroomList(body: ChatroomFetchAllParam, options: RequestInit = {}): Promise<TablesChatroom[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/list', options);
    }
    /**
     * @summary 修改群聊
     * @param {ChatroomUpdateParam} body 修改群聊参数
     * @param {*} [options] Override http request option.
     */
    public chatroomUpdate(body: ChatroomUpdateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/update', options);
    }
    /**
     * @summary 添加关键字
     * @param {KeywordCreateParam} body 添加关键字参数
     * @param {*} [options] Override http request option.
     */
    public keywordCreate(body: KeywordCreateParam, options: RequestInit = {}): Promise<number> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/create', options);
    }
    /**
     * @summary 删除关键字
     * @param {KeywordDeleteParam} body 删除关键字参数
     * @param {*} [options] Override http request option.
     */
    public keywordDelete(body: KeywordDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/delete', options);
    }
    /**
     * @summary 获取关键字
     * @param {KeywordFetchParam} body 获取关键字参数
     * @param {*} [options] Override http request option.
     */
    public keywordDetail(body: KeywordFetchParam, options: RequestInit = {}): Promise<TablesKeyword> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/detail', options);
    }
    /**
     * @summary 关键字列表
     * @param {KeywordFetchAllParam} body 获取关键字列表参数
     * @param {*} [options] Override http request option.
     */
    public keywordList(body: KeywordFetchAllParam, options: RequestInit = {}): Promise<TablesKeyword[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/list', options);
    }
    /**
     * @summary 修改关键字
     * @param {KeywordUpdateParam} body 修改关键字参数
     * @param {*} [options] Override http request option.
     */
    public keywordUpdate(body: KeywordUpdateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/update', options);
    }
    /**
     * @summary 添加模型
     * @param {LlmodelCreateParam} body 添加模型参数
     * @param {*} [options] Override http request option.
     */
    public llmodelCreate(body: LlmodelCreateParam, options: RequestInit = {}): Promise<number> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/llmodel/create', options);
    }
    /**
     * @summary 删除模型
     * @param {LlmodelDeleteParam} body 删除模型参数
     * @param {*} [options] Override http request option.
     */
    public llmodelDelete(body: LlmodelDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/llmodel/delete', options);
    }
    /**
     * @summary 获取模型
     * @param {LlmodelFetchParam} body 获取模型参数
     * @param {*} [options] Override http request option.
     */
    public llmodelDetail(body: LlmodelFetchParam, options: RequestInit = {}): Promise<TablesLLModel> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/llmodel/detail', options);
    }
    /**
     * @summary 模型列表
     * @param {LlmodelFetchAllParam} body 获取模型列表参数
     * @param {*} [options] Override http request option.
     */
    public llmodelList(body: LlmodelFetchAllParam, options: RequestInit = {}): Promise<TablesLLModel[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/llmodel/list', options);
    }
    /**
     * @summary 修改模型
     * @param {LlmodelUpdateParam} body 修改模型参数
     * @param {*} [options] Override http request option.
     */
    public llmodelUpdate(body: LlmodelUpdateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/llmodel/update', options);
    }
    /**
     * @summary 添加配置
     * @param {ProfileCreateParam} body 添加配置参数
     * @param {*} [options] Override http request option.
     */
    public profileCreate(body: ProfileCreateParam, options: RequestInit = {}): Promise<number> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/create', options);
    }
    /**
     * @summary 删除配置
     * @param {ProfileDeleteParam} body 删除配置参数
     * @param {*} [options] Override http request option.
     */
    public profileDelete(body: ProfileDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/delete', options);
    }
    /**
     * @summary 获取配置
     * @param {ProfileFetchParam} body 获取配置参数
     * @param {*} [options] Override http request option.
     */
    public profileDetail(body: ProfileFetchParam, options: RequestInit = {}): Promise<TablesProfile> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/detail', options);
    }
    /**
     * @summary 配置列表
     * @param {ProfileFetchAllParam} body 获取配置列表参数
     * @param {*} [options] Override http request option.
     */
    public profileList(body: ProfileFetchAllParam, options: RequestInit = {}): Promise<TablesProfile[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/list', options);
    }
    /**
     * @summary 修改配置
     * @param {ProfileUpdateParam} body 修改配置参数
     * @param {*} [options] Override http request option.
     */
    public profileUpdate(body: ProfileUpdateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/update', options);
    }
}

export interface ChatroomCreateParam {
    joinArgot: string;
    level: number;
    name: string;
    remark: string;
    revokeMsg: string;
    roomid: string;
    welcomeMsg: string;
}

export interface ChatroomDeleteParam {
    roomid: string;
}

export interface ChatroomFetchAllParam {
    level: number;
}

export interface ChatroomFetchParam {
    roomid: string;
}

export interface ChatroomUpdateParam {
    joinArgot: string;
    level: number;
    name: string;
    remark: string;
    revokeMsg: string;
    roomid: string;
    welcomeMsg: string;
}

export interface KeywordCreateParam {
    level: number;
    phrase: string;
    roomid: string;
}

export interface KeywordDeleteParam {
    phrase: string;
    roomid: string;
}

export interface KeywordFetchAllParam {
    level: number;
    roomid: string;
}

export interface KeywordFetchParam {
    phrase: string;
    roomid: string;
}

export interface KeywordUpdateParam {
    level: number;
    phrase: string;
    rd: number;
    roomid: string;
}

export interface LlmodelCreateParam {
    endpoint: string;
    family: string;
    mid: string;
    model: string;
    provider: string;
    secret: string;
}

export interface LlmodelDeleteParam {
    mid: string;
}

export interface LlmodelFetchAllParam {
    family: string;
    model: string;
    provider: string;
}

export interface LlmodelFetchParam {
    mid: string;
}

export interface LlmodelUpdateParam {
    endpoint: string;
    family: string;
    mid: string;
    model: string;
    provider: string;
    secret: string;
}

export interface ProfileCreateParam {
    aiArgot: string;
    aiModel: string;
    level: number;
    remark: string;
    roomid: string;
    wxid: string;
}

export interface ProfileDeleteParam {
    roomid: string;
    wxid: string;
}

export interface ProfileFetchAllParam {
    roomid: string;
    wxid: string;
}

export interface ProfileFetchParam {
    roomid: string;
    wxid: string;
}

export interface ProfileUpdateParam {
    aiArgot: string;
    aiModel: string;
    level: number;
    remark: string;
    roomid: string;
    wxid: string;
}

export interface TablesChatroom {
    // 创建时间戳
    createdAt: number;
    // 入群口令
    joinArgot: string;
    // 等级
    level: number;
    // 群聊名称
    name: string;
    // 主键
    rd: number;
    // 备注
    remark: string;
    // 防撤回消息
    revokeMsg: string;
    // 群聊 id
    roomid: string;
    // 最后更新时间戳
    updatedAt: number;
    // 欢迎消息
    welcomeMsg: string;
}

export interface TablesKeyword {
    // 创建时间戳
    createdAt: number;
    // 优先级等级
    level: number;
    // 词语或短语
    phrase: string;
    // 主键
    rd: number;
    // 群聊 id
    roomid: string;
    // 最后更新时间戳
    updatedAt: number;
}

export interface TablesLLModel {
    // 创建时间戳
    createdAt: number;
    // 仅 google 和 openai 支持自定义，留空则使用官方接口
    endpoint: string;
    // 模型家族，用于生成模型切换指令
    family: string;
    // 模型 Id
    mid: string;
    // 模型，必须和服务商提供的值对应
    model: string;
    // 服务商 [google, openai, xunfei]
    provider: string;
    // 主键
    rd: number;
    // 密钥，google 和 openai 填写 KEY，xunfei 填写 APP-ID,API-KEY,API-SECRET
    secret: string;
    // 最后更新时间戳
    updatedAt: number;
}

export interface TablesProfile {
    // 唤醒词
    aiArgot: string;
    // 会话模型
    aiModel: string;
    // 创建时间戳
    createdAt: number;
    // 等级
    level: number;
    // 主键
    rd: number;
    // 备注
    remark: string;
    // 群聊 id
    roomid: string;
    // 最后更新时间戳
    updatedAt: number;
    // 微信 id
    wxid: string;
}
