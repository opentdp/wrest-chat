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
     * @param {WcferryVerification} body 接受好友参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    acceptNewFriend(body: WcferryVerification, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/accept_new_friend', options);
    },
    /**
     * @summary 添加群成员
     * @param {WcferryMemberMgmt} body 管理群成员参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    addChatroomMembers(body: WcferryMemberMgmt, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
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
     * @param {unknown} [options] body 获取头像列表参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    avatarsPost(body?: unknown, options: RequestInit = {}): Promise<WcferryContactHeadImgUrlTable> {
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
    chatroomMembers(body: WcfrestGetChatRoomMembersRequest, options: RequestInit = {}): Promise<Array<WcferryRpcContact>> {
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
    chatrooms(body?: unknown, options: RequestInit = {}): Promise<Array<WcferryRpcContact>> {
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
    contacts(body?: unknown, options: RequestInit = {}): Promise<Array<WcferryRpcContact>> {
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
    dbTables(body: WcfrestGetDbTablesRequest, options: RequestInit = {}): Promise<Array<WcferryDbTable>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/db_tables', options);
    },
    /**
     * @summary 删除群成员
     * @param {WcferryMemberMgmt} body 管理群成员参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    delChatroomMembers(body: WcferryMemberMgmt, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
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
     * @param {WcferryForwardMsg} body 转发消息参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    forwardMsg(body: WcferryForwardMsg, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
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
    friends(body?: unknown, options: RequestInit = {}): Promise<Array<WcferryRpcContact>> {
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
     * @param {WcferryMemberMgmt} body 管理群成员参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    inviteChatroomMembers(body: WcferryMemberMgmt, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
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
     * @param {WcferryTransfer} body 接受转账参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    receiveTransfer(body: WcferryTransfer, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
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
    selfInfo(body?: unknown, options: RequestInit = {}): Promise<WcferryUserInfo> {
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
     * @param {WcferryPathMsg} body 发送文件消息参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendFile(body: WcferryPathMsg, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/send_file', options);
    },
    /**
     * @summary 发送图片消息
     * @param {WcferryPathMsg} body 发送图片消息参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendImg(body: WcferryPathMsg, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/send_img', options);
    },
    /**
     * @summary 拍一拍群友
     * @param {WcferryPatMsg} body 拍一拍群友参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendPatMsg(body: WcferryPatMsg, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/send_pat_msg', options);
    },
    /**
     * @summary 发送卡片消息
     * @param {WcferryRichText} body 发送卡片消息参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendRichText(body: WcferryRichText, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/send_rich_text', options);
    },
    /**
     * @summary 发送文本消息
     * @param {WcferryTextMsg} body 发送文本消息参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendTxt(body: WcferryTextMsg, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
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
    userInfo(body: WcfrestGetInfoByWxidRequest, options: RequestInit = {}): Promise<WcferryRpcContact> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        return httpRequest('/user_info', options);
    },
};

/**
 * @export
 * @interface WcferryContactHeadImgUrlTable
 */
export interface WcferryContactHeadImgUrlTable {
    /**
     * @type {string}
     * @memberof WcferryContactHeadImgUrlTable
     */
    bigHeadImgUrl?: string;
    /**
     * @type {string}
     * @memberof WcferryContactHeadImgUrlTable
     */
    headImgMd5?: string;
    /**
     * @type {number}
     * @memberof WcferryContactHeadImgUrlTable
     */
    reverse0?: number;
    /**
     * @type {unknown}
     * @memberof WcferryContactHeadImgUrlTable
     */
    reverse1?: unknown;
    /**
     * @type {string}
     * @memberof WcferryContactHeadImgUrlTable
     */
    smallHeadImgUrl?: string;
    /**
     * @type {string}
     * @memberof WcferryContactHeadImgUrlTable
     */
    usrName?: string;
}

/**
 * @export
 * @interface WcferryDbTable
 */
export interface WcferryDbTable {
    /**
     * 表名
     * @type {string}
     * @memberof WcferryDbTable
     */
    name?: string;
    /**
     * 建表 SQL
     * @type {string}
     * @memberof WcferryDbTable
     */
    sql?: string;
}

/**
 * @export
 * @interface WcferryForwardMsg
 */
export interface WcferryForwardMsg {
    /**
     * 待转发消息 ID
     * @type {number}
     * @memberof WcferryForwardMsg
     */
    id?: number;
    /**
     * 转发接收目标，群为 roomId，个人为 wxid
     * @type {string}
     * @memberof WcferryForwardMsg
     */
    receiver?: string;
}

/**
 * @export
 * @interface WcferryMemberMgmt
 */
export interface WcferryMemberMgmt {
    /**
     * 要加的群ID
     * @type {string}
     * @memberof WcferryMemberMgmt
     */
    roomid?: string;
    /**
     * 要加群的人列表，逗号分隔
     * @type {string}
     * @memberof WcferryMemberMgmt
     */
    wxids?: string;
}

/**
 * @export
 * @interface WcferryPatMsg
 */
export interface WcferryPatMsg {
    /**
     * 群 id
     * @type {string}
     * @memberof WcferryPatMsg
     */
    roomid?: string;
    /**
     * wxid
     * @type {string}
     * @memberof WcferryPatMsg
     */
    wxid?: string;
}

/**
 * @export
 * @interface WcferryPathMsg
 */
export interface WcferryPathMsg {
    /**
     * 要发送的图片的路径
     * @type {string}
     * @memberof WcferryPathMsg
     */
    path?: string;
    /**
     * 消息接收人
     * @type {string}
     * @memberof WcferryPathMsg
     */
    receiver?: string;
}

/**
 * @export
 * @interface WcferryRichText
 */
export interface WcferryRichText {
    /**
     * 公众号 id
     * @type {string}
     * @memberof WcferryRichText
     */
    account?: string;
    /**
     * 摘要
     * @type {string}
     * @memberof WcferryRichText
     */
    digest?: string;
    /**
     * 显示名字
     * @type {string}
     * @memberof WcferryRichText
     */
    name?: string;
    /**
     * 接收人
     * @type {string}
     * @memberof WcferryRichText
     */
    receiver?: string;
    /**
     * 缩略图
     * @type {string}
     * @memberof WcferryRichText
     */
    thumburl?: string;
    /**
     * 标题
     * @type {string}
     * @memberof WcferryRichText
     */
    title?: string;
    /**
     * 链接
     * @type {string}
     * @memberof WcferryRichText
     */
    url?: string;
}

/**
 * @export
 * @interface WcferryRpcContact
 */
export interface WcferryRpcContact {
    /**
     * 城市
     * @type {string}
     * @memberof WcferryRpcContact
     */
    city?: string;
    /**
     * 微信号
     * @type {string}
     * @memberof WcferryRpcContact
     */
    code?: string;
    /**
     * 国家
     * @type {string}
     * @memberof WcferryRpcContact
     */
    country?: string;
    /**
     * 性别
     * @type {number}
     * @memberof WcferryRpcContact
     */
    gender?: number;
    /**
     * 微信昵称
     * @type {string}
     * @memberof WcferryRpcContact
     */
    name?: string;
    /**
     * 省/州
     * @type {string}
     * @memberof WcferryRpcContact
     */
    province?: string;
    /**
     * 备注
     * @type {string}
     * @memberof WcferryRpcContact
     */
    remark?: string;
    /**
     * 微信 id
     * @type {string}
     * @memberof WcferryRpcContact
     */
    wxid?: string;
}

/**
 * @export
 * @interface WcferryTextMsg
 */
export interface WcferryTextMsg {
    /**
     * 要@的人列表，逗号分隔
     * @type {string}
     * @memberof WcferryTextMsg
     */
    aters?: string;
    /**
     * 要发送的消息内容
     * @type {string}
     * @memberof WcferryTextMsg
     */
    msg?: string;
    /**
     * 消息接收人，当为群时可@
     * @type {string}
     * @memberof WcferryTextMsg
     */
    receiver?: string;
}

/**
 * @export
 * @interface WcferryTransfer
 */
export interface WcferryTransfer {
    /**
     * Transaction id
     * @type {string}
     * @memberof WcferryTransfer
     */
    taid?: string;
    /**
     * 转账id transferid
     * @type {string}
     * @memberof WcferryTransfer
     */
    tfid?: string;
    /**
     * 转账人
     * @type {string}
     * @memberof WcferryTransfer
     */
    wxid?: string;
}

/**
 * @export
 * @interface WcferryUserInfo
 */
export interface WcferryUserInfo {
    /**
     * 文件/图片等父路径
     * @type {string}
     * @memberof WcferryUserInfo
     */
    home?: string;
    /**
     * 手机号
     * @type {string}
     * @memberof WcferryUserInfo
     */
    mobile?: string;
    /**
     * 昵称
     * @type {string}
     * @memberof WcferryUserInfo
     */
    name?: string;
    /**
     * 微信ID
     * @type {string}
     * @memberof WcferryUserInfo
     */
    wxid?: string;
}

/**
 * @export
 * @interface WcferryVerification
 */
export interface WcferryVerification {
    /**
     * 添加方式：17 名片，30 扫码
     * @type {number}
     * @memberof WcferryVerification
     */
    scene?: number;
    /**
     * 加密的用户名
     * @type {string}
     * @memberof WcferryVerification
     */
    v3?: string;
    /**
     * Ticket
     * @type {string}
     * @memberof WcferryVerification
     */
    v4?: string;
}

/**
 * @export
 * @interface WcfrestCommonPayload
 */
export interface WcfrestCommonPayload {
    /**
     * @type {unknown}
     * @memberof WcfrestCommonPayload
     */
    error?: unknown;
    /**
     * @type {string}
     * @memberof WcfrestCommonPayload
     */
    result?: string;
    /**
     * @type {boolean}
     * @memberof WcfrestCommonPayload
     */
    success?: boolean;
}

/**
 * @export
 * @interface WcfrestDbSqlQueryRequest
 */
export interface WcfrestDbSqlQueryRequest {
    /**
     * @type {string}
     * @memberof WcfrestDbSqlQueryRequest
     */
    db?: string;
    /**
     * @type {string}
     * @memberof WcfrestDbSqlQueryRequest
     */
    sql?: string;
}

/**
 * @export
 * @interface WcfrestDownloadAttachRequest
 */
export interface WcfrestDownloadAttachRequest {
    /**
     * @type {string}
     * @memberof WcfrestDownloadAttachRequest
     */
    extra?: string;
    /**
     * @type {number}
     * @memberof WcfrestDownloadAttachRequest
     */
    msgid?: number;
    /**
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
     * @type {string}
     * @memberof WcfrestDownloadImageRequest
     */
    dir?: string;
    /**
     * @type {string}
     * @memberof WcfrestDownloadImageRequest
     */
    extra?: string;
    /**
     * @type {number}
     * @memberof WcfrestDownloadImageRequest
     */
    msgid?: number;
    /**
     * @type {number}
     * @memberof WcfrestDownloadImageRequest
     */
    timeout?: number;
}

/**
 * @export
 * @interface WcfrestGetAliasInChatRoomRequest
 */
export interface WcfrestGetAliasInChatRoomRequest {
    /**
     * @type {string}
     * @memberof WcfrestGetAliasInChatRoomRequest
     */
    roomid?: string;
    /**
     * @type {string}
     * @memberof WcfrestGetAliasInChatRoomRequest
     */
    wxid?: string;
}

/**
 * @export
 * @interface WcfrestGetAudioMsgRequest
 */
export interface WcfrestGetAudioMsgRequest {
    /**
     * @type {number}
     * @memberof WcfrestGetAudioMsgRequest
     */
    msgid?: number;
    /**
     * @type {string}
     * @memberof WcfrestGetAudioMsgRequest
     */
    path?: string;
    /**
     * @type {number}
     * @memberof WcfrestGetAudioMsgRequest
     */
    timeout?: number;
}

/**
 * @export
 * @interface WcfrestGetChatRoomMembersRequest
 */
export interface WcfrestGetChatRoomMembersRequest {
    /**
     * @type {string}
     * @memberof WcfrestGetChatRoomMembersRequest
     */
    roomid?: string;
}

/**
 * @export
 * @interface WcfrestGetDbTablesRequest
 */
export interface WcfrestGetDbTablesRequest {
    /**
     * @type {string}
     * @memberof WcfrestGetDbTablesRequest
     */
    db?: string;
}

/**
 * @export
 * @interface WcfrestGetInfoByWxidRequest
 */
export interface WcfrestGetInfoByWxidRequest {
    /**
     * @type {string}
     * @memberof WcfrestGetInfoByWxidRequest
     */
    wxid?: string;
}

/**
 * @export
 * @interface WcfrestGetOcrRequest
 */
export interface WcfrestGetOcrRequest {
    /**
     * @type {string}
     * @memberof WcfrestGetOcrRequest
     */
    extra?: string;
    /**
     * @type {number}
     * @memberof WcfrestGetOcrRequest
     */
    timeout?: number;
}

/**
 * @export
 * @interface WcfrestReceiverRequest
 */
export interface WcfrestReceiverRequest {
    /**
     * @type {string}
     * @memberof WcfrestReceiverRequest
     */
    url?: string;
}

/**
 * @export
 * @interface WcfrestRefreshPyqRequest
 */
export interface WcfrestRefreshPyqRequest {
    /**
     * @type {number}
     * @memberof WcfrestRefreshPyqRequest
     */
    id?: number;
}

/**
 * @export
 * @interface WcfrestRevokeMsgRequest
 */
export interface WcfrestRevokeMsgRequest {
    /**
     * @type {number}
     * @memberof WcfrestRevokeMsgRequest
     */
    msgid?: number;
}
