
import { httpRequest } from "./request";
import { CronjobUpdateParam, KeywordUpdateParam } from "./wrobot";

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