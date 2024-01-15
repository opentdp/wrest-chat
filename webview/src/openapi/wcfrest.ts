export class RequiredError extends Error {
    override name = 'RequiredError';
    constructor(public field: string, msg?: string) {
        super(msg);
    }
}

export function httpRequest(input: string, options?: RequestInit) {
    return fetch('/api' + input, options).then((response) => {
        if (response.status >= 200 && response.status < 300) {
            return response.json();
        } else {
            throw response;
        }
    });
}

export const WrestApi = {
    /**
     * 
     * @summary 接受好友请求
     * @param {WcferryVerification} body 接受好友请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    acceptNewFriend(body: WcferryVerification, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling acceptNewFriendPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/accept_new_friend`, options);
    },
    /**
     * 
     * @summary 添加群成员
     * @param {WcferryMemberMgmt} body 增删群成员请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    addChatroomMembers(body: WcferryMemberMgmt, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling addChatroomMembersPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/add_chatroom_members`, options);
    },
    /**
     * 
     * @summary 获取群成员昵称
     * @param {string} wxid wxid
     * @param {string} roomid 群id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    aliasInChatroomWxidRoomid(wxid: string, roomid: string, options: RequestInit = {}): Promise<string> {
        if (wxid === null || wxid === undefined) {
            throw new RequiredError('wxid', 'Required parameter wxid was null or undefined when calling aliasInChatroomWxidRoomidGet.');
        }
        if (roomid === null || roomid === undefined) {
            throw new RequiredError('roomid', 'Required parameter roomid was null or undefined when calling aliasInChatroomWxidRoomidGet.');
        }

        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/alias_in_chatroom/${wxid}/${roomid}`, options);
    },
    /**
     * 
     * @summary 获取群成员列表
     * @param {string} roomid 群id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    chatroomMembersRoomid(roomid: string, options: RequestInit = {}): Promise<Array<WcferryRpcContact>> {
        if (roomid === null || roomid === undefined) {
            throw new RequiredError('roomid', 'Required parameter roomid was null or undefined when calling chatroomMembersRoomidGet.');
        }

        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/chatroom_members/${roomid}`, options);
    },
    /**
     * 
     * @summary 获取群列表
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    chatrooms(options: RequestInit = {}): Promise<Array<WcferryRpcContact>> {
        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/chatrooms`, options);
    },
    /**
     * 
     * @summary 获取完整通讯录
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    contacts(options: RequestInit = {}): Promise<Array<WcferryRpcContact>> {
        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/contacts`, options);
    },
    /**
     * 
     * @summary 获取数据库列表
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    dbNames(options: RequestInit = {}): Promise<Array<string>> {
        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/db_names`, options);
    },
    /**
     * 
     * @summary 执行数据库查询
     * @param {WcfrestDbSqlQueryRequest} body 数据库查询请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    dbQuerySql(body: WcfrestDbSqlQueryRequest, options: RequestInit = {}): Promise<Array<{ [key: string]: unknown; }>> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling dbQuerySqlPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/db_query_sql`, options);
    },
    /**
     * 
     * @summary 获取数据库表列表
     * @param {string} db 数据库名
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    dbTablesDb(db: string, options: RequestInit = {}): Promise<Array<WcferryDbTable>> {
        if (db === null || db === undefined) {
            throw new RequiredError('db', 'Required parameter db was null or undefined when calling dbTablesDbGet.');
        }

        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/db_tables/${db}`, options);
    },
    /**
     * 
     * @summary 删除群成员
     * @param {WcferryMemberMgmt} body 增删群成员请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    delChatroomMembers(body: WcferryMemberMgmt, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling delChatroomMembersPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/del_chatroom_members`, options);
    },
    /**
     * 
     * @summary 关闭推送消息到URL
     * @param {WcfrestReceiverRequest} body 消息推送请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    disableReceiver(body: WcfrestReceiverRequest, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling disableReceiverPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/disable_receiver`, options);
    },
    /**
     * 
     * @summary 下载附件
     * @param {WcfrestDownloadAttachRequest} body 下载附件参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    downloadAttach(body: WcfrestDownloadAttachRequest, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling downloadAttachPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/download_attach`, options);
    },
    /**
     * 
     * @summary 下载图片
     * @param {WcfrestDownloadImageRequest} body 下载图片参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    downloadImage(body: WcfrestDownloadImageRequest, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling downloadImagePost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/download_image`, options);
    },
    /**
     * 
     * @summary 开启推送消息到URL
     * @param {WcfrestReceiverRequest} body 消息推送请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    enableReceiver(body: WcfrestReceiverRequest, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling enableReceiverPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/enable_receiver`, options);
    },
    /**
     * 
     * @summary 转发消息
     * @param {WcferryForwardMsg} body 转发消息请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    forwardMsg(body: WcferryForwardMsg, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling forwardMsgPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/forward_msg`, options);
    },
    /**
     * 
     * @summary 获取好友列表
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    friends(options: RequestInit = {}): Promise<Array<WcferryRpcContact>> {
        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/friends`, options);
    },
    /**
     * 
     * @summary 获取语音消息
     * @param {WcfrestGetAudioMsgRequest} body 语音消息请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    getAudioMsg(body: WcfrestGetAudioMsgRequest, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling getAudioMsgPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/get_audio_msg`, options);
    },
    /**
     * 
     * @summary 获取OCR识别结果
     * @param {WcfrestGetOcrRequest} body 文本请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    getOcrResult(body: WcfrestGetOcrRequest, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling getOcrResultPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/get_ocr_result`, options);
    },
    /**
     * 
     * @summary 邀请群成员
     * @param {WcferryMemberMgmt} body 增删群成员请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    inviteChatroomMembers(body: WcferryMemberMgmt, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling inviteChatroomMembersPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/invite_chatroom_members`, options);
    },
    /**
     * 
     * @summary 检查登录状态
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    isLogin(options: RequestInit = {}): Promise<boolean> {
        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/is_login`, options);
    },
    /**
     * 
     * @summary 获取所有消息类型
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    msgTypes(options: RequestInit = {}): Promise<{ [key: string]: string; }> {
        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/msg_types`, options);
    },
    /**
     * 
     * @summary 接受转账
     * @param {WcferryTransfer} body 接受转账请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    receiveTransfer(body: WcferryTransfer, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling receiveTransferPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/receive_transfer`, options);
    },
    /**
     * 
     * @summary 刷新朋友圈
     * @param {number} id 朋友圈id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    refreshPyqId(id: number, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (id === null || id === undefined) {
            throw new RequiredError('id', 'Required parameter id was null or undefined when calling refreshPyqIdGet.');
        }

        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/refresh_pyq/${id}`, options);
    },
    /**
     * 
     * @summary 撤回消息
     * @param {number} msgid 消息id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    revokeMsgMsgid(msgid: number, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (msgid === null || msgid === undefined) {
            throw new RequiredError('msgid', 'Required parameter msgid was null or undefined when calling revokeMsgMsgidGet.');
        }

        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/revoke_msg/${msgid}`, options);
    },
    /**
     * 
     * @summary 获取登录账号个人信息
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    selfInfo(options: RequestInit = {}): Promise<WcferryUserInfo> {
        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/self_info`, options);
    },
    /**
     * 
     * @summary 获取登录账号wxid
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    selfWxid(options: RequestInit = {}): Promise<string> {
        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/self_wxid`, options);
    },
    /**
     * 
     * @summary 发送文件消息
     * @param {WcferryPathMsg} body 文件消息请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendFile(body: WcferryPathMsg, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling sendFilePost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/send_file`, options);
    },
    /**
     * 
     * @summary 发送图片消息
     * @param {WcferryPathMsg} body 图片消息请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendImg(body: WcferryPathMsg, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling sendImgPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/send_img`, options);
    },
    /**
     * 
     * @summary 拍一拍群友
     * @param {WcferryPatMsg} body 拍一拍请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendPatMsg(body: WcferryPatMsg, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling sendPatMsgPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/send_pat_msg`, options);
    },
    /**
     * 
     * @summary 发送卡片消息
     * @param {WcferryRichText} body 卡片消息请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendRichText(body: WcferryRichText, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling sendRichTextPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/send_rich_text`, options);
    },
    /**
     * 
     * @summary 发送文本消息
     * @param {WcferryTextMsg} body 文本消息请求参数
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    sendTxt(body: WcferryTextMsg, options: RequestInit = {}): Promise<WcfrestRespPayload> {
        if (body === null || body === undefined) {
            throw new RequiredError('body', 'Required parameter body was null or undefined when calling sendTxtPost.');
        }

        options = Object.assign({ method: 'POST' }, options);
        options.headers = Object.assign({ 'Content-Type': 'application/json' }, options.headers);
        options.body = JSON.stringify(body || {});

        return httpRequest(`/send_txt`, options);
    },
    /**
     * 
     * @summary 根据wxid获取个人信息
     * @param {string} wxid wxid
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userInfoWxid(wxid: string, options: RequestInit = {}): Promise<WcferryRpcContact> {
        if (wxid === null || wxid === undefined) {
            throw new RequiredError('wxid', 'Required parameter wxid was null or undefined when calling userInfoWxidGet.');
        }

        options = Object.assign({ method: 'GET' }, options);
        options.headers = Object.assign({}, options.headers);

        return httpRequest(`/user_info/${wxid}`, options);
    },
};

/**
 * 
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
 * 
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
 * 
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
 * 
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
 * 
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
 * 
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
 * 
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
 * 
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
 * 
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
 * 
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
 * 
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
 * 
 * @export
 * @interface WcfrestDbSqlQueryRequest
 */
export interface WcfrestDbSqlQueryRequest {
    /**
     * 
     * @type {string}
     * @memberof WcfrestDbSqlQueryRequest
     */
    db?: string;
    /**
     * 
     * @type {string}
     * @memberof WcfrestDbSqlQueryRequest
     */
    sql?: string;
}

/**
 * 
 * @export
 * @interface WcfrestDownloadAttachRequest
 */
export interface WcfrestDownloadAttachRequest {
    /**
     * 
     * @type {string}
     * @memberof WcfrestDownloadAttachRequest
     */
    extra?: string;
    /**
     * 
     * @type {number}
     * @memberof WcfrestDownloadAttachRequest
     */
    msgid?: number;
    /**
     * 
     * @type {string}
     * @memberof WcfrestDownloadAttachRequest
     */
    thumb?: string;
}

/**
 * 
 * @export
 * @interface WcfrestDownloadImageRequest
 */
export interface WcfrestDownloadImageRequest {
    /**
     * 
     * @type {string}
     * @memberof WcfrestDownloadImageRequest
     */
    dir?: string;
    /**
     * 
     * @type {string}
     * @memberof WcfrestDownloadImageRequest
     */
    extra?: string;
    /**
     * 
     * @type {number}
     * @memberof WcfrestDownloadImageRequest
     */
    msgid?: number;
    /**
     * 
     * @type {number}
     * @memberof WcfrestDownloadImageRequest
     */
    timeout?: number;
}

/**
 * 
 * @export
 * @interface WcfrestGetAudioMsgRequest
 */
export interface WcfrestGetAudioMsgRequest {
    /**
     * 
     * @type {number}
     * @memberof WcfrestGetAudioMsgRequest
     */
    msgid?: number;
    /**
     * 
     * @type {string}
     * @memberof WcfrestGetAudioMsgRequest
     */
    path?: string;
    /**
     * 
     * @type {number}
     * @memberof WcfrestGetAudioMsgRequest
     */
    timeout?: number;
}

/**
 * 
 * @export
 * @interface WcfrestGetOcrRequest
 */
export interface WcfrestGetOcrRequest {
    /**
     * 
     * @type {string}
     * @memberof WcfrestGetOcrRequest
     */
    extra?: string;
    /**
     * 
     * @type {number}
     * @memberof WcfrestGetOcrRequest
     */
    timeout?: number;
}

/**
 * 
 * @export
 * @interface WcfrestReceiverRequest
 */
export interface WcfrestReceiverRequest {
    /**
     * 
     * @type {string}
     * @memberof WcfrestReceiverRequest
     */
    url?: string;
}

/**
 * 
 * @export
 * @interface WcfrestRespPayload
 */
export interface WcfrestRespPayload {
    /**
     * 
     * @type {unknown}
     * @memberof WcfrestRespPayload
     */
    error?: unknown;
    /**
     * 
     * @type {string}
     * @memberof WcfrestRespPayload
     */
    result?: string;
    /**
     * 
     * @type {boolean}
     * @memberof WcfrestRespPayload
     */
    success?: boolean;
}
