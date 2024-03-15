
import { httpRequest } from "./request";
import { KeywordCreateParam } from "./wrobot";

export const SundryApi = {
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
     * @summary 获取指令列表
     * @param {HandlerListParam} body 获取指令列表参数
     * @param {*} [options] Override http request option.
     */
    handlerList(body: HandlerListParam, options: RequestInit = {}): Promise<Handler[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/handler/list', options);
    },
    /**
     * @summary 获取计划任务插件
     * @param {*} body 获取计划任务插件参数
     * @param {*} [options] Override http request option.
     */
    pluginCronjobs(body = {}, options: RequestInit = {}): Promise<Handler[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/plugin/cronjobs', options);
    },
    /**
     * @summary 获取外部指令插件
     * @param {*} body 获取外部指令插件参数
     * @param {*} [options] Override http request option.
     */
    pluginKeywords(body = {}, options: RequestInit = {}): Promise<Handler[]> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/plugin/keywords', options);
    },
};

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
    entryId?: number;
    // 要删除的计划的 ID
    rd?: number;
}

export interface CronjobFetchAllParam {
    // 要获取的计划的类型
    type?: string;
}

export interface CronjobFetchParam {
    // 要获取的计划的 ID
    entryId?: number;
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
    entryId?: number;
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

export interface HandlerListParam {
    // 重装指令
    reset?: boolean;
}

export interface CronjobPlugin {
    // 插件信息
    config: CronjobCreateParam;
    // 错误信息
    error: string;
    // 插件文件
    file: string;
}

export interface KeywordPlugin {
    // 插件信息
    config: KeywordCreateParam;
    // 错误信息
    error: string;
    // 插件文件
    file: string;
}