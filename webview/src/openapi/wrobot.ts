import { httpRequest } from "./request";

export const RobotApi = {
    /**
     * @summary 添加群聊配置
     * @param {ChatroomCreateParam} body 添加群聊配置参数
     * @param {*} [options] Override http request option.
     */
    chatroomCreate(body: ChatroomCreateParam, options: RequestInit = {}): Promise<number> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/create', options);
    },
    /**
     * @summary 删除群聊配置
     * @param {ChatroomDeleteParam} body 删除群聊配置参数
     * @param {*} [options] Override http request option.
     */
    chatroomDelete(body: ChatroomDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/delete', options);
    },
    /**
     * @summary 获取群聊配置
     * @param {ChatroomFetchParam} body 获取群聊配置参数
     * @param {*} [options] Override http request option.
     */
    chatroomDetail(body: ChatroomFetchParam, options: RequestInit = {}): Promise<TablesChatroom> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/detail', options);
    },
    /**
     * @summary 群聊配置列表
     * @param {ChatroomFetchAllParam} body 获取群聊配置列表参数
     * @param {*} [options] Override http request option.
     */
    chatroomList(body: ChatroomFetchAllParam, options: RequestInit = {}): Promise<TablesChatroom[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/list', options);
    },
    /**
     * @summary 修改群聊配置
     * @param {ChatroomUpdateParam} body 修改群聊配置参数
     * @param {*} [options] Override http request option.
     */
    chatroomUpdate(body: ChatroomUpdateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/chatroom/update', options);
    },
    /**
     * @summary 添加关键词
     * @param {KeywordCreateParam} body 添加关键词参数
     * @param {*} [options] Override http request option.
     */
    keywordCreate(body: KeywordCreateParam, options: RequestInit = {}): Promise<number> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/create', options);
    },
    /**
     * @summary 删除关键词
     * @param {KeywordDeleteParam} body 删除关键词参数
     * @param {*} [options] Override http request option.
     */
    keywordDelete(body: KeywordDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/delete', options);
    },
    /**
     * @summary 获取关键词
     * @param {KeywordFetchParam} body 获取关键词参数
     * @param {*} [options] Override http request option.
     */
    keywordDetail(body: KeywordFetchParam, options: RequestInit = {}): Promise<TablesKeyword> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/detail', options);
    },
    /**
     * @summary 关键词列表
     * @param {KeywordFetchAllParam} body 获取关键词列表参数
     * @param {*} [options] Override http request option.
     */
    keywordList(body: KeywordFetchAllParam, options: RequestInit = {}): Promise<TablesKeyword[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/keyword/list', options);
    },
    /**
     * @summary 修改关键词
     * @param {KeywordUpdateParam} body 修改关键词参数
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
     * @summary 添加用户配置
     * @param {ProfileCreateParam} body 添加用户配置参数
     * @param {*} [options] Override http request option.
     */
    profileCreate(body: ProfileCreateParam, options: RequestInit = {}): Promise<number> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/create', options);
    },
    /**
     * @summary 删除用户配置
     * @param {ProfileDeleteParam} body 删除用户配置参数
     * @param {*} [options] Override http request option.
     */
    profileDelete(body: ProfileDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/delete', options);
    },
    /**
     * @summary 获取用户配置
     * @param {ProfileFetchParam} body 获取用户配置参数
     * @param {*} [options] Override http request option.
     */
    profileDetail(body: ProfileFetchParam, options: RequestInit = {}): Promise<TablesProfile> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/detail', options);
    },
    /**
     * @summary 用户配置列表
     * @param {ProfileFetchAllParam} body 获取用户配置列表参数
     * @param {*} [options] Override http request option.
     */
    profileList(body: ProfileFetchAllParam, options: RequestInit = {}): Promise<TablesProfile[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/list', options);
    },
    /**
     * @summary 修改用户配置
     * @param {ProfileUpdateParam} body 修改用户配置参数
     * @param {*} [options] Override http request option.
     */
    profileUpdate(body: ProfileUpdateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/profile/update', options);
    },
    /**
     * @summary 创建全局设置
     * @param {SettingCreateParam} body 创建全局设置参数
     * @param {*} [options] Override http request option.
     */
    settingCreate(body: SettingCreateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/setting/create', options);
    },
    /**
     * @summary 删除全局设置
     * @param {SettingDeleteParam} body 删除全局设置参数
     * @param {*} [options] Override http request option.
     */
    settingDelete(body: SettingDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/setting/delete', options);
    },
    /**
     * @summary 获取全局设置
     * @param {SettingFetchParam} body 获取全局设置参数
     * @param {*} [options] Override http request option.
     */
    settingDetail(body: SettingFetchParam, options: RequestInit = {}): Promise<TablesSetting> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/setting/detail', options);
    },
    /**
     * @summary 全局设置列表
     * @param {SettingFetchAllParam} body 获取全局设置列表参数
     * @param {*} [options] Override http request option.
     */
    settingList(body: SettingFetchAllParam, options: RequestInit = {}): Promise<TablesSetting[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/setting/list', options);
    },
    /**
     * @summary 修改全局设置
     * @param {SettingUpdateParam} body 修改全局设置参数
     * @param {*} [options] Override http request option.
     */
    settingUpdate(body: SettingUpdateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/setting/update', options);
    },
};

export interface ChatroomCreateParam {
    // 加群指令
    join_argot?: string;
    // 等级
    level?: number;
    // 群聊名称
    name?: string;
    // 回应拍拍我
    pat_return?: string;
    // 备注
    remark?: string;
    // 防撤回消息
    revoke_msg?: string;
    // 群聊 id
    roomid: string;
    // 欢迎消息
    welcome_msg?: string;
}

export interface ChatroomDeleteParam {
    // 主键
    rd: number;
    // 群聊 id
    roomid?: string;
}

export interface ChatroomFetchAllParam {
    // 等级
    level?: number;
}

export interface ChatroomFetchParam {
    // 主键
    rd: number;
    // 群聊 id
    roomid?: string;
}

export interface ChatroomUpdateParam {
    // 加群指令
    join_argot: string;
    // 等级
    level: number;
    // 群聊名称
    name: string;
    // 回应拍拍我
    pat_return: string;
    // 备注
    remark: string;
    // 防撤回消息
    revoke_msg: string;
    // 主键
    rd: number;
    // 群聊 id
    roomid: string;
    // 欢迎消息
    welcome_msg: string;
}

export interface KeywordCreateParam {
    // 等级
    level?: number;
    // 词语或短语
    phrase: string;
    // 群聊 id
    roomid: string;
}

export interface KeywordDeleteParam {
    // 词语或短语
    phrase?: string;
    // 主键
    rd: number;
    // 群聊 id
    roomid?: string;
}

export interface KeywordFetchAllParam {
    // 等级
    level?: number;
    // 群聊 id
    roomid?: string;
}

export interface KeywordFetchParam {
    // 词语或短语
    phrase?: string;
    // 主键
    rd: number;
    // 群聊 id
    roomid?: string;
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
    endpoint?: string;
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
    mid?: string;
    // 主键
    rd: number;
}

export interface LlmodelFetchAllParam {
    // 模型家族
    family?: string;
    // 模型名称
    model?: string;
    // 服务商名称
    provider?: string;
}

export interface LlmodelFetchParam {
    // 模型 Id
    mid?: string;
    // 主键
    rd: number;
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
    // 主键
    rd: number;
    // 密钥
    secret: string;
}

export interface ProfileCreateParam {
    // 唤醒词
    ai_argot?: string;
    // 会话模型
    ai_model?: string;
    // 封禁期限
    ban_expire?: number;
    // 等级
    level?: number;
    // 备注
    remark?: string;
    // 群聊 id
    roomid: string;
    // 微信 id
    wxid: string;
}

export interface ProfileDeleteParam {
    // 主键
    rd: number;
    // 群聊 id
    roomid?: string;
    // 微信 id
    wxid?: string;
}

export interface ProfileFetchAllParam {
    // 群聊 id
    roomid?: string;
    // 微信 id
    wxid?: string;
}

export interface ProfileFetchParam {
    // 主键
    rd: number;
    // 群聊 id
    roomid?: string;
    // 微信 id
    wxid?: string;
}

export interface ProfileUpdateParam {
    // 唤醒词
    ai_argot: string;
    // 会话模型
    ai_model: string;
    // 封禁期限
    ban_expire: number;
    // 等级
    level: number;
    // 备注
    remark: string;
    // 主键
    rd: number;
    // 群聊 id
    roomid: string;
    // 微信 id
    wxid: string;
}

export interface SettingCreateParam {
    // 分组
    group?: string;
    // 键名
    name: string;
    // 备注
    remark?: string;
    // 标题
    title?: string;
    // 类型
    type?: string;
    // 键值
    value?: string;
}

export interface SettingDeleteParam {
    // 键名
    name?: string;
    // 主键
    rd: number;
}

export interface SettingFetchAllParam {
    // 分组
    group?: string;
}

export interface SettingFetchParam {
    // 键名
    name?: string;
    // 主键
    rd: number;
}

export interface SettingUpdateParam {
    // 分组
    group: string;
    // 键名
    name: string;
    // 主键
    rd: number;
    // 备注
    remark: string;
    // 标题
    title: string;
    // 类型
    type: string;
    // 键值
    value: string;
}

export interface TablesChatroom {
    // 创建时间戳
    created_at: number;
    // 加群指令
    join_argot: string;
    // 等级
    level: number;
    // 群聊名称
    name: string;
    // 回应拍拍我
    pat_return: string;
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

export interface TablesSetting {
    // 创建时间戳
    createdAt: number;
    // 分组
    group: string;
    // 键名
    name: string;
    // 主键
    rd: number;
    // 备注
    remark: string;
    // 标题
    title: string;
    // 类型
    type: string;
    // 最后更新时间戳
    updatedAt: number;
    // 键值
    value: string;
}
