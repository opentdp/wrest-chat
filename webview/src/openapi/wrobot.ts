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
     * 
     * @summary 添加计划任务
     * @param {CronjobCreateParam} body 添加计划任务参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    cronjobCreate(body: CronjobCreateParam, options: RequestInit = {}): Promise<number> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/cronjob/create', options);
    },
    /**
     * @summary 删除计划任务
     * @param {CronjobDeleteParam} body 删除计划任务参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    cronjobDelete(body: CronjobDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/cronjob/delete', options);
    },
    /**
     * @summary 获取计划任务
     * @param {CronjobFetchParam} body 获取计划任务参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    cronjobDetail(body: CronjobFetchParam, options: RequestInit = {}): Promise<TablesCronjob> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/cronjob/detail', options);
    },
    /**
     * @summary 计划任务列表
     * @param {CronjobFetchAllParam} body 获取计划任务列表参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    cronjobList(body: CronjobFetchAllParam, options: RequestInit = {}): Promise<TablesCronjob[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/cronjob/list', options);
    },
    /**
     * @summary 修改计划任务
     * @param {CronjobUpdateParam} body 修改计划任务参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    cronjobUpdate(body: CronjobUpdateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/cronjob/update', options);
    },
    /**
     * @summary 获取计划任务状态
     * @param {CronjobStatusParam} body 获取计划任务状态参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    cronjobStatus(body: CronjobStatusParam, options: RequestInit = {}): Promise<Record<number, CronjobStatusPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/cronjob/status', options);
    },
    /**
     * @summary 触发计划任务
     * @param {CronjobFetchParam} body 触发计划任务参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    cronjobExecute(body: CronjobFetchParam, options: RequestInit = {}): Promise<Record<number, CronjobStatusPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/cronjob/execute', options);
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
    /**
     * @summary 创建 webhook
     * @param {WebhookCreateWebhookParam} body 创建webhook参数
     * @param {*} [options] Override http request option.
     */
    webhookCreate(body: WebhookCreateWebhookParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/webhook/create', options);
    },
    /**
     * @summary webhook 详情
     * @param {WebhookFetchWebhookParam} body 参数
     * @param {*} [options] Override http request option.
     */
    webhookDetail(body: WebhookFetchWebhookParam, options: RequestInit = {}): Promise<Webhook> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/webhook/detail', options);
    },
    /**
     * @summary 删除 webhook
     * @param {WebhookDeleteWebhookParam} body 删除webhook参数
     * @param {*} [options] Override http request option.
     */
    webhookDelete(body: WebhookDeleteWebhookParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/bot/webhook/delete', options);
    },
    /**
     * @summary Webhook 列表
     * @param options
     * @param {*} [options] Override http request option.
     */
    webhookList(options: RequestInit = {}): Promise<Webhook[]> {
        options = { method: 'POST', ...options };
        return httpRequest('/bot/webhook/list', options);
    },
};

export interface ChatroomCreateParam {
    // 加群指令
    join_argot?: string;
    // 等级
    level?: number;
    //  定义默认模型
    model_default?: string;
    // 定义模型扮演的身份
    model_context?: string;
    // 历史消息数量
    model_history?: number;
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
    //  定义默认模型
    model_default: string;
    // 定义模型扮演的身份
    model_context: string;
    // 历史消息数量
    model_history: number;
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

export interface CronjobCreateParam {
    // 要执行的命令内容
    content: string;
    // 每月执行计划的日期
    day_of_month: string;
    // 每周执行计划的星期
    day_of_week: string;
    // 执行结果交付方式
    deliver: string;
    // 命令执行的工作目录
    directory: string;
    // 执行计划的小时
    hour: string;
    // 执行计划的分钟
    minute: string;
    // 执行计划的月份
    month: string;
    // 计划的名称
    name: string;
    // 执行计划的秒数
    second: string;
    // 命令执行的超时时间 (秒)
    timeout: number;
    // 命令类型 (CMD, POWERSHELL, SHELL)
    type: string;
}

export interface CronjobDeleteParam {
    // 要删除的计划的 ID
    entry_id?: number;
    // 要删除的计划的 ID
    rd?: number;
}

export interface CronjobFetchAllParam {
    // 要获取的计划的类型
    type?: string;
}

export interface CronjobFetchParam {
    // 要获取的计划的 ID
    entry_id?: number;
    // 要获取的计划的 ID
    rd?: number;
}

export interface CronjobUpdateParam {
    // 要执行的命令内容
    content: string;
    // 每月执行计划的日期
    day_of_month: string;
    // 每周执行计划的星期
    day_of_week: string;
    // 执行结果交付方式
    deliver: string;
    // 命令执行的工作目录
    directory: string;
    // 当前计划的 ID
    entry_id?: number;
    // 执行计划的小时
    hour: string;
    // 执行计划的分钟
    minute: string;
    // 执行计划的月份
    month: string;
    // 计划的名称
    name: string;
    // 计划的 ID
    rd: number;
    // 执行计划的秒数
    second: string;
    // 命令执行的超时时间 (秒)
    timeout: number;
    // 命令类型 (CMD, POWERSHELL, SHELL)
    type: string;
}

export interface CronjobStatusParam {

}

export interface CronjobStatusPayload {
    // 当前任务 ID
    entry_id: number;
    // 下次执行时间
    next_time: number;
    // 上次执行时间
    prev_time: number;
}

export interface KeywordCreateParam {
    // 分组
    group: string;
    // 等级
    level?: number;
    // 短语
    phrase: string;
    // 群聊 id
    roomid: string;
    // 目标
    target?: string;
    // 备注
    remark?: string;
}

export interface KeywordDeleteParam {
    // 分组
    group?: string;
    // 短语
    phrase?: string;
    // 主键
    rd: number;
    // 群聊 id
    roomid?: string;
}

export interface KeywordFetchAllParam {
    // 分组
    group?: string;
    // 等级
    level?: number;
    // 群聊 id
    roomid?: string;
    // 目标
    target?: string;
}

export interface KeywordFetchParam {
    // 分组
    group?: string;
    // 短语
    phrase?: string;
    // 主键
    rd: number;
    // 群聊 id
    roomid?: string;
    // 目标
    target?: string;
}

export interface KeywordUpdateParam {
    // 分组
    group: string;
    // 等级
    level: number;
    // 短语
    phrase: string;
    // 主键
    rd: number;
    // 群聊 id
    roomid: string;
    // 目标
    target: string;
    // 备注
    remark: string;
}

export interface LlmodelCreateParam {
    // 接口地址
    endpoint?: string;
    // 模型家族
    family: string;
    // 等级
    level?: number;
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
    // 等级
    level?: number;
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
    // 等级
    level: number;
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
    // 等级
    level?: number;
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
    //  定义默认模型
    model_default: string;
    // 定义模型扮演的身份
    model_context: string;
    // 历史消息数量
    model_history: number;
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

export interface TablesCronjob {
    // 要执行的命令内容
    content: string;
    // 创建时间戳
    created_at: number;
    // 每月执行计划的日期
    day_of_month: string;
    // 每周执行计划的星期
    day_of_week: string;
    // 执行结果交付方式
    deliver: string;
    // 命令执行的工作目录
    directory: string;
    // 计划的 ID
    entry_id: number;
    // 执行计划的小时
    hour: string;
    // 执行计划的分钟
    minute: string;
    // 执行计划的月份
    month: string;
    // 计划的名称
    name: string; //
    // 计划的 ID
    rd: number; //
    // 执行计划的秒数
    second: string;
    // 命令执行的超时时间 (秒)
    timeout: number;
    // 命令类型 (CMD, POWERSHELL, SHELL)
    type: string;
    // 最后更新时间戳
    updated_at: number;
}

export interface TablesKeyword {
    // 创建时间戳
    created_at: number;
    // 分组
    group: string;
    // 优先级等级
    level: number;
    // 短语
    phrase: string;
    // 主键
    rd: number;
    // 群聊 id
    roomid: string;
    // 目标
    target: string;
    // 备注
    remark: string;
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
    // 等级，用于限制用户访问
    level: number;
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

export interface Webhook {
    // 主键
    rd: number;
    // 目标ID， 用户或者是群聊ID
    target_id: string;
    // token
    token: string;
    // 备注
    remark: string;
    // 创建时间戳
    created_at: number;
    // 最后更新时间戳
    updated_at: number;
}

export interface WebhookCreateWebhookParam {
    // 备注
    remark?: string;
    // 目标ID：用户或者是群聊ID
    target_id: string;
    // 目标类型
    target: string;
}

export interface WebhookFetchWebhookParam {
    // 主键
    rd: number;
}

export interface WebhookDeleteWebhookParam {
    // 主键
    rd: number;
}
