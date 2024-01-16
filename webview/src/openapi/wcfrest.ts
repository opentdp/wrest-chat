export class RequiredError extends Error {
    override name = 'RequiredError';
    constructor(public field: string, msg?: string) {
        super(msg);
    }
}

export function httpRequest(input: string, options: RequestInit = {}) {
    return fetch('/api' + input, options).then(response => {
        if (response.status >= 200 && response.status < 300) {
            return response.json().then(data => {
                if (data.Error) {
                    throw data.Error;
                }
                return data.Payload;
            });
        } else {
            throw response;
        }
    });
}

export const WrestApi = {
    /**
     * @summary 接受好友请求
     * @param {WcfrestAcceptNewFriendRequest} body 接受好友参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    acceptNewFriend(body: WcfrestAcceptNewFriendRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/accept_new_friend', options);
    },
    /**
     * @summary 添加群成员
     * @param {WcfrestChatroomMembersRequest} body 管理群成员参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    addChatroomMembers(body: WcfrestChatroomMembersRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/add_chatroom_members', options);
    },
    /**
     * @summary 获取群成员昵称
     * @param {WcfrestGetAliasInChatRoomRequest} body 获取群成员昵称参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    aliasInChatroom(body: WcfrestGetAliasInChatRoomRequest, options: RequestInit = {}): Promise<string> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/alias_in_chatroom', options);
    },
    /**
     * @summary 获取头像列表
     * @param {WcfrestGetAvatarsRequest} body 获取头像列表参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    avatars(body: WcfrestGetAvatarsRequest, options: RequestInit = {}): Promise<Array<WcfrestAvatarPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/avatars', options);
    },
    /**
     * @summary 获取群成员列表
     * @param {WcfrestGetChatRoomMembersRequest} body 获取群成员列表参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    chatroomMembers(body: WcfrestGetChatRoomMembersRequest, options: RequestInit = {}): Promise<Array<WcfrestContactPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/chatroom_members', options);
    },
    /**
     * @summary 获取群列表
     * @param {unknown} [options] body 获取群列表参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    chatrooms(body?: unknown, options: RequestInit = {}): Promise<Array<WcfrestContactPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/chatrooms', options);
    },
    /**
     * @summary 获取完整通讯录
     * @param {unknown} [options] body 获取完整通讯录参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    contacts(body?: unknown, options: RequestInit = {}): Promise<Array<WcfrestContactPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/contacts', options);
    },
    /**
     * @summary 获取数据库列表
     * @param {unknown} [options] body 获取数据库列表参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    dbNames(body?: unknown, options: RequestInit = {}): Promise<Array<string>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/db_names', options);
    },
    /**
     * @summary 执行数据库查询
     * @param {WcfrestDbSqlQueryRequest} body 数据库查询参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    dbQuerySql(body: WcfrestDbSqlQueryRequest, options: RequestInit = {}): Promise<Array<{ [key: string]: unknown; }>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/db_query_sql', options);
    },
    /**
     * @summary 获取数据库表列表
     * @param {WcfrestGetDbTablesRequest} body 获取数据库表列表参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    dbTables(body: WcfrestGetDbTablesRequest, options: RequestInit = {}): Promise<Array<WcfrestDbTablePayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/db_tables', options);
    },
    /**
     * @summary 删除群成员
     * @param {WcfrestChatroomMembersRequest} body 管理群成员参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    delChatroomMembers(body: WcfrestChatroomMembersRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/del_chatroom_members', options);
    },
    /**
     * @summary 关闭推送消息到URL
     * @param {WcfrestReceiverRequest} body 推送消息到URL参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    disableReceiver(body: WcfrestReceiverRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/disable_receiver', options);
    },
    /**
     * @summary 下载附件
     * @param {WcfrestDownloadAttachRequest} body 下载附件参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    downloadAttach(body: WcfrestDownloadAttachRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/download_attach', options);
    },
    /**
     * @summary 下载图片
     * @param {WcfrestDownloadImageRequest} body 下载图片参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    downloadImage(body: WcfrestDownloadImageRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/download_image', options);
    },
    /**
     * @summary 开启推送消息到URL
     * @param {WcfrestReceiverRequest} body 推送消息到URL参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    enableReceiver(body: WcfrestReceiverRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/enable_receiver', options);
    },
    /**
     * @summary 转发消息
     * @param {WcfrestForwardMsgRequest} body 转发消息参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    forwardMsg(body: WcfrestForwardMsgRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/forward_msg', options);
    },
    /**
     * @summary 获取好友列表
     * @param {unknown} [options] body 获取好友列表参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    friends(body?: unknown, options: RequestInit = {}): Promise<Array<WcfrestContactPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/friends', options);
    },
    /**
     * @summary 获取语音消息
     * @param {WcfrestGetAudioMsgRequest} body 获取语音消息参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    getAudioMsg(body: WcfrestGetAudioMsgRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/get_audio_msg', options);
    },
    /**
     * @summary 获取OCR识别结果
     * @param {WcfrestGetOcrRequest} body 获取OCR识别结果参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    getOcrResult(body: WcfrestGetOcrRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/get_ocr_result', options);
    },
    /**
     * @summary 邀请群成员
     * @param {WcfrestChatroomMembersRequest} body 管理群成员参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    inviteChatroomMembers(body: WcfrestChatroomMembersRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/invite_chatroom_members', options);
    },
    /**
     * @summary 检查登录状态
     * @param {unknown} [options] body 检查登录状态参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    isLogin(body?: unknown, options: RequestInit = {}): Promise<boolean> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/is_login', options);
    },
    /**
     * @summary 获取所有消息类型
     * @param {unknown} [options] body 获取所有消息类型参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    msgTypes(body?: unknown, options: RequestInit = {}): Promise<{ [key: string]: string; }> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/msg_types', options);
    },
    /**
     * @summary 接受转账
     * @param {WcfrestReceiveTransferRequest} body 接受转账参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    receiveTransfer(body: WcfrestReceiveTransferRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/receive_transfer', options);
    },
    /**
     * @summary 刷新朋友圈
     * @param {WcfrestRefreshPyqRequest} body 刷新朋友圈参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    refreshPyq(body: WcfrestRefreshPyqRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/refresh_pyq', options);
    },
    /**
     * @summary 撤回消息
     * @param {WcfrestRevokeMsgRequest} body 撤回消息参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    revokeMsg(body: WcfrestRevokeMsgRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/revoke_msg', options);
    },
    /**
     * @summary 获取登录账号个人信息
     * @param {unknown} [options] body 获取数据库列表参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    selfInfo(body?: unknown, options: RequestInit = {}): Promise<WcfrestUserInfoPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/self_info', options);
    },
    /**
     * @summary 获取登录账号wxid
     * @param {unknown} [options] body 获取登录账号wxid参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    selfWxid(body?: unknown, options: RequestInit = {}): Promise<string> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/self_wxid', options);
    },
    /**
     * @summary 发送文件消息
     * @param {WcfrestSendFileRequest} body 发送文件消息参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendFile(body: WcfrestSendFileRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/send_file', options);
    },
    /**
     * @summary 发送图片消息
     * @param {WcfrestSendImgRequest} body 发送图片消息参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendImg(body: WcfrestSendImgRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/send_img', options);
    },
    /**
     * @summary 拍一拍群友
     * @param {WcfrestSendPatMsgRequest} body 拍一拍群友参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendPatMsg(body: WcfrestSendPatMsgRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/send_pat_msg', options);
    },
    /**
     * @summary 发送卡片消息
     * @param {WcfrestSendRichTextRequest} body 发送卡片消息参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendRichText(body: WcfrestSendRichTextRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/send_rich_text', options);
    },
    /**
     * @summary 发送文本消息
     * @param {WcfrestSendTxtRequest} body 发送文本消息参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendTxt(body: WcfrestSendTxtRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/send_txt', options);
    },
    /**
     * @summary 根据wxid获取个人信息
     * @param {WcfrestGetInfoByWxidRequest} body 根据wxid获取个人信息参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userInfo(body: WcfrestGetInfoByWxidRequest, options: RequestInit = {}): Promise<WcfrestContactPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/user_info', options);
    },
};

/**
 * @export
 * @interface WcfrestAcceptNewFriendRequest
 */
export interface WcfrestAcceptNewFriendRequest {
    /**
     * 添加方式：17 名片，30 扫码
     * @type {number}
     * @memberof WcfrestAcceptNewFriendRequest
     */
    scene: number;
    /**
     * 加密的用户名
     * @type {string}
     * @memberof WcfrestAcceptNewFriendRequest
     */
    v3: string;
    /**
     * 验证信息 Ticket
     * @type {string}
     * @memberof WcfrestAcceptNewFriendRequest
     */
    v4: string;
}

/**
 * @export
 * @interface WcfrestAvatarPayload
 */
export interface WcfrestAvatarPayload {
    /**
     * 大头像 url
     * @type {string}
     * @memberof WcfrestAvatarPayload
     */
    big_head_img_url: string;
    /**
     * 小头像 url
     * @type {string}
     * @memberof WcfrestAvatarPayload
     */
    small_head_img_url: string;
    /**
     * 用户 id
     * @type {string}
     * @memberof WcfrestAvatarPayload
     */
    usr_name: string;
}

/**
 * @export
 * @interface WcfrestChatroomMembersRequest
 */
export interface WcfrestChatroomMembersRequest {
    /**
     * 群聊 id
     * @type {string}
     * @memberof WcfrestChatroomMembersRequest
     */
    roomid: string;
    /**
     * 用户 id 列表
     * @type {Array<string>}
     * @memberof WcfrestChatroomMembersRequest
     */
    wxids: Array<string>;
}

/**
 * @export
 * @interface WcfrestCommonPayload
 */
export interface WcfrestCommonPayload {
    /**
     * 错误信息
     * @type {unknown}
     * @memberof WcfrestCommonPayload
     */
    error?: unknown;
    /**
     * 返回结果
     * @type {string}
     * @memberof WcfrestCommonPayload
     */
    result?: string;
    /**
     * 是否成功
     * @type {boolean}
     * @memberof WcfrestCommonPayload
     */
    success?: boolean;
}

/**
 * @export
 * @interface WcfrestContactPayload
 */
export interface WcfrestContactPayload {
    /**
     * 城市
     * @type {string}
     * @memberof WcfrestContactPayload
     */
    city: string;
    /**
     * 微信号
     * @type {string}
     * @memberof WcfrestContactPayload
     */
    code: string;
    /**
     * 国家
     * @type {string}
     * @memberof WcfrestContactPayload
     */
    country: string;
    /**
     * 性别
     * @type {number}
     * @memberof WcfrestContactPayload
     */
    gender: number;
    /**
     * 昵称
     * @type {string}
     * @memberof WcfrestContactPayload
     */
    name: string;
    /**
     * 省/州
     * @type {string}
     * @memberof WcfrestContactPayload
     */
    province: string;
    /**
     * 备注
     * @type {string}
     * @memberof WcfrestContactPayload
     */
    remark: string;
    /**
     * 用户 id
     * @type {string}
     * @memberof WcfrestContactPayload
     */
    wxid: string;
}

/**
 * @export
 * @interface WcfrestDbSqlQueryRequest
 */
export interface WcfrestDbSqlQueryRequest {
    /**
     * 数据库名称
     * @type {string}
     * @memberof WcfrestDbSqlQueryRequest
     */
    db: string;
    /**
     * 待执行的 SQL
     * @type {string}
     * @memberof WcfrestDbSqlQueryRequest
     */
    sql: string;
}

/**
 * @export
 * @interface WcfrestDbTablePayload
 */
export interface WcfrestDbTablePayload {
    /**
     * 表名
     * @type {string}
     * @memberof WcfrestDbTablePayload
     */
    name: string;
    /**
     * 建表 SQL
     * @type {string}
     * @memberof WcfrestDbTablePayload
     */
    sql: string;
}

/**
 * @export
 * @interface WcfrestDownloadAttachRequest
 */
export interface WcfrestDownloadAttachRequest {
    /**
     * 消息中的 extra 字段
     * @type {string}
     * @memberof WcfrestDownloadAttachRequest
     */
    extra: string;
    /**
     * 消息 id
     * @type {number}
     * @memberof WcfrestDownloadAttachRequest
     */
    msgid: number;
    /**
     * 消息中的 thumb 字段
     * @type {string}
     * @memberof WcfrestDownloadAttachRequest
     */
    thumb?: string;
}

/**
 * @export
 * @interface WcfrestDownloadImageRequest
 */
export interface WcfrestDownloadImageRequest {
    /**
     * 存储路径
     * @type {string}
     * @memberof WcfrestDownloadImageRequest
     */
    dir: string;
    /**
     * 消息中的 extra 字段
     * @type {string}
     * @memberof WcfrestDownloadImageRequest
     */
    extra: string;
    /**
     * 消息 id
     * @type {number}
     * @memberof WcfrestDownloadImageRequest
     */
    msgid: number;
    /**
     * 超时重试次数
     * @type {number}
     * @memberof WcfrestDownloadImageRequest
     */
    timeout?: number;
}

/**
 * @export
 * @interface WcfrestForwardMsgRequest
 */
export interface WcfrestForwardMsgRequest {
    /**
     * 待转发消息 id
     * @type {number}
     * @memberof WcfrestForwardMsgRequest
     */
    id: number;
    /**
     * 转发接收人或群的 id 列表
     * @type {Array<string>}
     * @memberof WcfrestForwardMsgRequest
     */
    receiver: Array<string>;
}

/**
 * @export
 * @interface WcfrestGetAliasInChatRoomRequest
 */
export interface WcfrestGetAliasInChatRoomRequest {
    /**
     * 群聊 id
     * @type {string}
     * @memberof WcfrestGetAliasInChatRoomRequest
     */
    roomid: string;
    /**
     * 用户 id
     * @type {string}
     * @memberof WcfrestGetAliasInChatRoomRequest
     */
    wxid: string;
}

/**
 * @export
 * @interface WcfrestGetAudioMsgRequest
 */
export interface WcfrestGetAudioMsgRequest {
    /**
     * 消息 id
     * @type {number}
     * @memberof WcfrestGetAudioMsgRequest
     */
    msgid: number;
    /**
     * 存储路径
     * @type {string}
     * @memberof WcfrestGetAudioMsgRequest
     */
    path: string;
    /**
     * 超时重试次数
     * @type {number}
     * @memberof WcfrestGetAudioMsgRequest
     */
    timeout?: number;
}

/**
 * @export
 * @interface WcfrestGetAvatarsRequest
 */
export interface WcfrestGetAvatarsRequest {
    /**
     * 用户 id 列表
     * @type {Array<string>}
     * @memberof WcfrestGetAvatarsRequest
     */
    wxids: Array<string>;
}

/**
 * @export
 * @interface WcfrestGetChatRoomMembersRequest
 */
export interface WcfrestGetChatRoomMembersRequest {
    /**
     * 群聊 id
     * @type {string}
     * @memberof WcfrestGetChatRoomMembersRequest
     */
    roomid: string;
}

/**
 * @export
 * @interface WcfrestGetDbTablesRequest
 */
export interface WcfrestGetDbTablesRequest {
    /**
     * 数据库名称
     * @type {string}
     * @memberof WcfrestGetDbTablesRequest
     */
    db: string;
}

/**
 * @export
 * @interface WcfrestGetInfoByWxidRequest
 */
export interface WcfrestGetInfoByWxidRequest {
    /**
     * 用户 id
     * @type {string}
     * @memberof WcfrestGetInfoByWxidRequest
     */
    wxid: string;
}

/**
 * @export
 * @interface WcfrestGetOcrRequest
 */
export interface WcfrestGetOcrRequest {
    /**
     * 消息中的 extra 字段
     * @type {string}
     * @memberof WcfrestGetOcrRequest
     */
    extra: string;
    /**
     * 超时重试次数
     * @type {number}
     * @memberof WcfrestGetOcrRequest
     */
    timeout?: number;
}

/**
 * @export
 * @interface WcfrestReceiveTransferRequest
 */
export interface WcfrestReceiveTransferRequest {
    /**
     * Transaction id
     * @type {string}
     * @memberof WcfrestReceiveTransferRequest
     */
    taid: string;
    /**
     * 转账id transferid
     * @type {string}
     * @memberof WcfrestReceiveTransferRequest
     */
    tfid: string;
    /**
     * 转账人
     * @type {string}
     * @memberof WcfrestReceiveTransferRequest
     */
    wxid: string;
}

/**
 * @export
 * @interface WcfrestReceiverRequest
 */
export interface WcfrestReceiverRequest {
    /**
     * 接收推送消息的 url
     * @type {string}
     * @memberof WcfrestReceiverRequest
     */
    url: string;
}

/**
 * @export
 * @interface WcfrestRefreshPyqRequest
 */
export interface WcfrestRefreshPyqRequest {
    /**
     * 分页 id
     * @type {number}
     * @memberof WcfrestRefreshPyqRequest
     */
    id: number;
}

/**
 * @export
 * @interface WcfrestRevokeMsgRequest
 */
export interface WcfrestRevokeMsgRequest {
    /**
     * 消息 id
     * @type {number}
     * @memberof WcfrestRevokeMsgRequest
     */
    msgid: number;
}

/**
 * @export
 * @interface WcfrestSendFileRequest
 */
export interface WcfrestSendFileRequest {
    /**
     * 文件路径
     * @type {string}
     * @memberof WcfrestSendFileRequest
     */
    path: string;
    /**
     * 接收人或群的 id
     * @type {string}
     * @memberof WcfrestSendFileRequest
     */
    receiver: string;
}

/**
 * @export
 * @interface WcfrestSendImgRequest
 */
export interface WcfrestSendImgRequest {
    /**
     * 图片路径
     * @type {string}
     * @memberof WcfrestSendImgRequest
     */
    path: string;
    /**
     * 接收人或群的 id
     * @type {string}
     * @memberof WcfrestSendImgRequest
     */
    receiver: string;
}

/**
 * @export
 * @interface WcfrestSendPatMsgRequest
 */
export interface WcfrestSendPatMsgRequest {
    /**
     * 群 id
     * @type {string}
     * @memberof WcfrestSendPatMsgRequest
     */
    roomid: string;
    /**
     * 用户 id
     * @type {string}
     * @memberof WcfrestSendPatMsgRequest
     */
    wxid: string;
}

/**
 * @export
 * @interface WcfrestSendRichTextRequest
 */
export interface WcfrestSendRichTextRequest {
    /**
     * 填公众号 id 可以显示对应的头像（gh_ 开头的）
     * @type {string}
     * @memberof WcfrestSendRichTextRequest
     */
    account: string;
    /**
     * 摘要，三行
     * @type {string}
     * @memberof WcfrestSendRichTextRequest
     */
    digest: string;
    /**
     * 左下显示的名字
     * @type {string}
     * @memberof WcfrestSendRichTextRequest
     */
    name: string;
    /**
     * 接收人或群的 id
     * @type {string}
     * @memberof WcfrestSendRichTextRequest
     */
    receiver: string;
    /**
     * 缩略图的链接
     * @type {string}
     * @memberof WcfrestSendRichTextRequest
     */
    thumburl: string;
    /**
     * 标题，最多两行
     * @type {string}
     * @memberof WcfrestSendRichTextRequest
     */
    title: string;
    /**
     * 点击后跳转的链接
     * @type {string}
     * @memberof WcfrestSendRichTextRequest
     */
    url: string;
}

/**
 * @export
 * @interface WcfrestSendTxtRequest
 */
export interface WcfrestSendTxtRequest {
    /**
     * 需要 At 的用户 id 列表
     * @type {Array<string>}
     * @memberof WcfrestSendTxtRequest
     */
    aters?: Array<string>;
    /**
     * 消息内容
     * @type {string}
     * @memberof WcfrestSendTxtRequest
     */
    msg: string;
    /**
     * 接收人或群的 id
     * @type {string}
     * @memberof WcfrestSendTxtRequest
     */
    receiver: string;
}

/**
 * @export
 * @interface WcfrestUserInfoPayload
 */
export interface WcfrestUserInfoPayload {
    /**
     * 文件/图片等父路径
     * @type {string}
     * @memberof WcfrestUserInfoPayload
     */
    home: string;
    /**
     * 手机号
     * @type {string}
     * @memberof WcfrestUserInfoPayload
     */
    mobile: string;
    /**
     * 昵称
     * @type {string}
     * @memberof WcfrestUserInfoPayload
     */
    name: string;
    /**
     * 用户 id
     * @type {string}
     * @memberof WcfrestUserInfoPayload
     */
    wxid: string;
}
