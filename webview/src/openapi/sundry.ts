
import { httpRequest } from "./request";
import { KeywordUpdateParam } from "./wrobot";

export const SundryApi = {
    /**
     * 
     * @summary 获取模型配置
     * @param {AiChatParam} body 获取模型配置参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    aichatConfig(body: AiChatParam, options: RequestInit = {}): Promise<AiChatUserConfig> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/aichat/config', options);
    },
    /**
     * 
     * @summary 发起文本聊天
     * @param {AiChatParam} body 发起文本聊天参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    aichatText(body: AiChatParam, options: RequestInit = {}): Promise<string> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/aichat/text', options);
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
        return httpRequest('/api/cronjob/create', options);
    },
    /**
     * @summary 删除计划任务
     * @param {CronjobDeleteParam} body 删除计划任务参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    cronjobDelete(body: CronjobDeleteParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/cronjob/delete', options);
    },
    /**
     * @summary 获取计划任务
     * @param {CronjobFetchParam} body 获取计划任务参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    cronjobDetail(body: CronjobFetchParam, options: RequestInit = {}): Promise<TablesCronjob> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/cronjob/detail', options);
    },
    /**
     * @summary 计划任务列表
     * @param {CronjobFetchAllParam} body 获取计划任务列表参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    cronjobList(body: CronjobFetchAllParam, options: RequestInit = {}): Promise<TablesCronjob[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/cronjob/list', options);
    },
    /**
     * @summary 修改计划任务
     * @param {CronjobUpdateParam} body 修改计划任务参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    cronjobUpdate(body: CronjobUpdateParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/cronjob/update', options);
    },
    /**
     * @summary 获取计划任务状态
     * @param {CronjobStatusParam} body 获取计划任务状态参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    cronjobStatus(body: CronjobStatusParam, options: RequestInit = {}): Promise<Record<number, CronjobStatusPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/cronjob/status', options);
    },
    /**
     * @summary 触发计划任务
     * @param {CronjobFetchParam} body 触发计划任务参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    cronjobExecute(body: CronjobFetchParam, options: RequestInit = {}): Promise<Record<number, CronjobStatusPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/cronjob/execute', options);
    },
    /**
     * @summary 获取计划任务插件
     * @param {*} body 获取计划任务插件参数
     * @param {*} [options] Override http request option.
     */
    pluginCronjobs(body = {}, options: RequestInit = {}): Promise<CronjobPlugin[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/plugin/cronjobs', options);
    },
    /**
     * @summary 获取外部指令插件
     * @param {*} body 获取外部指令插件参数
     * @param {*} [options] Override http request option.
     */
    pluginKeywords(body = {}, options: RequestInit = {}): Promise<KeywordPlugin[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/plugin/keywords', options);
    },
    /**
     * @summary 获取系统版本
     * @param {*} body 获取外部指令插件参数
     * @param {*} [options] Override http request option.
     */
    systemVersion(body = {}, options: RequestInit = {}): Promise<SystemVersion> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/system/version', options);
    },
    /**
     * @summary 获取指令列表
     * @param {HandlersParam} body 获取指令列表参数
     * @param {*} [options] Override http request option.
     */
    systemHandlers(body: HandlersParam, options: RequestInit = {}): Promise<Handler[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/system/handlers', options);
    },
};

export interface AiChatParam {
    // 微信 ID
    wxid: string;
    // 聊天内容
    message: string;
}

export interface AiChatMsgHistory {
    // 角色, 'user' | 'assistant'
    role: string;
    // 消息内容
    content: string;
}

export interface AiChatUserConfig {
    // 族类描述
    family: string;
    // 供应商
    provider: string;
    // 接口地址
    endpoint: string;
    // 模型
    model: string;
    // 角色设定
    role_context: string;
    // 消息历史记录
    msg_historys: AiChatMsgHistory[];
    // 消息记录最大条数
    msg_history_max: number;
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

export interface CronjobPlugin {
    // 插件信息
    config: CronjobUpdateParam;
    // 错误信息
    error: string;
    // 插件文件
    file: string;
}

export interface KeywordPlugin {
    // 插件信息
    config: KeywordUpdateParam;
    // 错误信息
    error: string;
    // 插件文件
    file: string;
}

export interface SystemVersion {
    // 系统版本
    version: string;
    // 系统编译版本
    build_version: string;
    // wcferry 版本
    wcf_version: string;
    // wechat 版本
    wechat_version: string;
}

export interface Handler {
    // 0:不限制 7:群管理 9:创始人
    level: number;
    // 排序，越小越靠前
    order: number;
    // 群聊 id
    roomid: string;
    // 指令
    command: string;
    // 指令的描述信息
    describe: string;
}

export interface HandlersParam {
    // 重装指令
    reset?: boolean;
}