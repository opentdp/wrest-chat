
import { httpRequest } from "./request";

export const CronjobApi = {
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
    cronjobStatus(body: CronjobStatusParam, options: RequestInit = {}): Promise<unknown> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/api/cronjob/status', options);
    },
};

export interface CronjobCreateParam {
    // 要执行的命令内容
    content: string;
    // 每月执行计划的日期
    dayOfMonth: string;
    // 每周执行计划的星期
    dayOfWeek: string;
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
    // 计划的 ID (如果已存在)
    rd?: number;
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
    dayOfMonth: string;
    // 每周执行计划的星期
    dayOfWeek: string;
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
    // 计划的 ID (如果已存在)
    rd?: number;
    // 执行计划的秒数
    second: string;
    // 命令执行的超时时间 (秒)
    timeout: number;
    // 命令类型 (CMD, POWERSHELL, SHELL)
    type: string;
}

export interface CronjobStatusParam {
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
    dayOfMonth: string;
    // 每周执行计划的星期
    dayOfWeek: string;
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
    name: string;
    // 计划的 ID
    rd?: number;
    // 执行计划的秒数
    second: string;
    // 命令执行的超时时间 (秒)
    timeout: number;
    // 命令类型 (CMD, POWERSHELL, SHELL)
    type: string;
    // 最后更新时间戳
    updated_at: number;
}
