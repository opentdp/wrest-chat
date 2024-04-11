import { httpRequest } from "./request";

export const WrestApi = {
    /**
     * @summary 接受好友请求
     * @param {WcfrestAcceptNewFriendRequest} body 接受好友参数
     * @param {*} [options] Override http request option.
     */
    acceptNewFriend(body: WcfrestAcceptNewFriendRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/accept_new_friend', options);
    },
    /**
     * @summary 添加群成员
     * @param {WcfrestChatroomMembersRequest} body 管理群成员参数
     * @param {*} [options] Override http request option.
     */
    addChatroomMembers(body: WcfrestChatroomMembersRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/add_chatroom_members', options);
    },
    /**
     * @summary 获取群成员昵称
     * @param {WcfrestGetAliasInChatRoomRequest} body 获取群成员昵称参数
     * @param {*} [options] Override http request option.
     */
    aliasInChatroom(body: WcfrestGetAliasInChatRoomRequest, options: RequestInit = {}): Promise<string> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/alias_in_chatroom', options);
    },
    /**
     * @summary 获取头像列表
     * @param {WcfrestGetAvatarsRequest} body 获取头像列表参数
     * @param {*} [options] Override http request option.
     */
    avatars(body: WcfrestGetAvatarsRequest, options: RequestInit = {}): Promise<Array<WcfrestAvatarPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/avatars', options);
    },
    /**
     * @summary 获取群成员列表
     * @param {WcfrestGetChatRoomMembersRequest} body 获取群成员列表参数
     * @param {*} [options] Override http request option.
     */
    chatroomMembers(body: WcfrestGetChatRoomMembersRequest, options: RequestInit = {}): Promise<Array<WcfrestContactPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/chatroom_members', options);
    },
    /**
     * @summary 获取群列表
     * @param {unknown} [options] body 获取群列表参数
     * @param {*} [options] Override http request option.
     */
    chatrooms(body?: unknown, options: RequestInit = {}): Promise<Array<WcfrestContactPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/chatrooms', options);
    },
    /**
     * @summary 获取完整通讯录
     * @param {unknown} [options] body 获取完整通讯录参数
     * @param {*} [options] Override http request option.
     */
    contacts(body?: unknown, options: RequestInit = {}): Promise<Array<WcfrestContactPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/contacts', options);
    },
    /**
     * @summary 获取数据库列表
     * @param {unknown} [options] body 获取数据库列表参数
     * @param {*} [options] Override http request option.
     */
    dbNames(body?: unknown, options: RequestInit = {}): Promise<Array<string>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/db_names', options);
    },
    /**
     * @summary 执行数据库查询
     * @param {WcfrestDbSqlQueryRequest} body 数据库查询参数
     * @param {*} [options] Override http request option.
     */
    dbQuerySql(body: WcfrestDbSqlQueryRequest, options: RequestInit = {}): Promise<Array<{ [key: string]: unknown; }>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/db_query_sql', options);
    },
    /**
     * @summary 获取数据库表列表
     * @param {WcfrestGetDbTablesRequest} body 获取数据库表列表参数
     * @param {*} [options] Override http request option.
     */
    dbTables(body: WcfrestGetDbTablesRequest, options: RequestInit = {}): Promise<Array<WcfrestDbTablePayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/db_tables', options);
    },
    /**
     * @summary 删除群成员
     * @param {WcfrestChatroomMembersRequest} body 管理群成员参数
     * @param {*} [options] Override http request option.
     */
    delChatroomMembers(body: WcfrestChatroomMembersRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/del_chatroom_members', options);
    },
    /**
     * @summary 关闭推送消息到URL
     * @param {WcfrestReceiverRequest} body 推送消息到URL参数
     * @param {*} [options] Override http request option.
     */
    disableReceiver(body: WcfrestReceiverRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/disable_receiver', options);
    },
    /**
     * @summary 下载附件
     * @param {WcfrestDownloadAttachRequest} body 下载附件参数
     * @param {*} [options] Override http request option.
     */
    downloadAttach(body: WcfrestDownloadAttachRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/download_attach', options);
    },
    /**
     * @summary 下载图片
     * @param {WcfrestDownloadImageRequest} body 下载图片参数
     * @param {*} [options] Override http request option.
     */
    downloadImage(body: WcfrestDownloadImageRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/download_image', options);
    },
    /**
     * @summary 开启推送消息到URL
     * @param {WcfrestReceiverRequest} body 推送消息到URL参数
     * @param {*} [options] Override http request option.
     */
    enableReceiver(body: WcfrestReceiverRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/enable_receiver', options);
    },
    /**
     * @summary 转发消息
     * @param {WcfrestForwardMsgRequest} body 转发消息参数
     * @param {*} [options] Override http request option.
     */
    forwardMsg(body: WcfrestForwardMsgRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/forward_msg', options);
    },
    /**
     * @summary 获取好友列表
     * @param {unknown} [options] body 获取好友列表参数
     * @param {*} [options] Override http request option.
     */
    friends(body?: unknown, options: RequestInit = {}): Promise<Array<WcfrestContactPayload>> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/friends', options);
    },
    /**
     * @summary 获取语音消息
     * @param {WcfrestGetAudioMsgRequest} body 获取语音消息参数
     * @param {*} [options] Override http request option.
     */
    getAudioMsg(body: WcfrestGetAudioMsgRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/get_audio_msg', options);
    },
    /**
     * @summary 获取OCR识别结果
     * @param {WcfrestGetOcrRequest} body 获取OCR识别结果参数
     * @param {*} [options] Override http request option.
     */
    getOcrResult(body: WcfrestGetOcrRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/get_ocr_result', options);
    },
    /**
     * @summary 邀请群成员
     * @param {WcfrestChatroomMembersRequest} body 管理群成员参数
     * @param {*} [options] Override http request option.
     */
    inviteChatroomMembers(body: WcfrestChatroomMembersRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/invite_chatroom_members', options);
    },
    /**
     * @summary 检查登录状态
     * @param {unknown} [options] body 检查登录状态参数
     * @param {*} [options] Override http request option.
     */
    isLogin(body?: unknown, options: RequestInit = {}): Promise<boolean> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/is_login', options);
    },
    /**
     * @summary 登录二维码
     * @param {unknown} [options] body 获取登录二维码参数
     * @param {*} [options] Override http request option.
     */
    loginQr(body?: unknown, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/login_qr', options);
    },
    /**
     * @summary 获取所有消息类型
     * @param {unknown} [options] body 获取所有消息类型参数
     * @param {*} [options] Override http request option.
     */
    msgTypes(body?: unknown, options: RequestInit = {}): Promise<{ [key: string]: string; }> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/msg_types', options);
    },
    /**
     * @summary 接受转账
     * @param {WcfrestReceiveTransferRequest} body 接受转账参数
     * @param {*} [options] Override http request option.
     */
    receiveTransfer(body: WcfrestReceiveTransferRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/receive_transfer', options);
    },
    /**
     * @summary 刷新朋友圈
     * @param {WcfrestRefreshPyqRequest} body 刷新朋友圈参数
     * @param {*} [options] Override http request option.
     */
    refreshPyq(body: WcfrestRefreshPyqRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/refresh_pyq', options);
    },
    /**
     * @summary 撤回消息
     * @param {WcfrestRevokeMsgRequest} body 撤回消息参数
     * @param {*} [options] Override http request option.
     */
    revokeMsg(body: WcfrestRevokeMsgRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/revoke_msg', options);
    },
    /**
     * @summary 获取登录账号个人信息
     * @param {unknown} [options] body 获取数据库列表参数
     * @param {*} [options] Override http request option.
     */
    selfInfo(body?: unknown, options: RequestInit = {}): Promise<WcfrestUserInfoPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/self_info', options);
    },
    /**
     * @summary 获取登录账号wxid
     * @param {unknown} [options] body 获取登录账号wxid参数
     * @param {*} [options] Override http request option.
     */
    selfWxid(body?: unknown, options: RequestInit = {}): Promise<string> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/self_wxid', options);
    },
    /**
     * @summary 发送文件消息
     * @param {WcfrestSendFileRequest} body 发送文件消息参数
     * @param {*} [options] Override http request option.
     */
    sendFile(body: WcfrestSendFileRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/send_file', options);
    },
    /**
     * @summary 发送图片消息
     * @param {WcfrestSendFileRequest} body 发送图片消息参数
     * @param {*} [options] Override http request option.
     */
    sendImg(body: WcfrestSendFileRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/send_img', options);
    },
    /**
     * @summary 拍一拍群友
     * @param {WcfrestSendPatMsgRequest} body 拍一拍群友参数
     * @param {*} [options] Override http request option.
     */
    sendPatMsg(body: WcfrestSendPatMsgRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/send_pat_msg', options);
    },
    /**
     * @summary 发送卡片消息
     * @param {WcfrestSendRichTextRequest} body 发送卡片消息参数
     * @param {*} [options] Override http request option.
     */
    sendRichText(body: WcfrestSendRichTextRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/send_rich_text', options);
    },
    /**
     * @summary 发送文本消息
     * @param {WcfrestSendTxtRequest} body 发送文本消息参数
     * @param {*} [options] Override http request option.
     */
    sendTxt(body: WcfrestSendTxtRequest, options: RequestInit = {}): Promise<WcfrestCommonPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/send_txt', options);
    },
    /**
     * @summary 根据wxid获取个人信息
     * @param {WcfrestGetInfoByWxidRequest} body 根据wxid获取个人信息参数
     * @param {*} [options] Override http request option.
     */
    userInfo(body: WcfrestGetInfoByWxidRequest, options: RequestInit = {}): Promise<WcfrestContactPayload> {
        options = { method: 'POST', body: JSON.stringify(body || {}), ...options };
        return httpRequest('/wcf/user_info', options);
    },
};

export interface WcfrestAcceptNewFriendRequest {
    // 添加方式：17 名片，30 扫码
    scene: number;
    // 加密的用户名
    v3: string;
    // 验证信息 Ticket
    v4: string;
}

export interface WcfrestAvatarPayload {
    // 大头像 url
    big_head_img_url: string;
    // 小头像 url
    small_head_img_url: string;
    // 用户 id
    usr_name: string;
}

export interface WcfrestChatroomMembersRequest {
    // 群聊 id
    roomid: string;
    // 用户 id 列表
    wxids: Array<string>;
}

export interface WcfrestCommonPayload {
    // 错误信息
    error?: unknown;
    // 返回结果
    result?: string;
    // 是否成功
    success?: boolean;
}

export interface WcfrestContactPayload {
    // 城市
    city: string;
    // 微信号
    code: string;
    // 国家
    country: string;
    // 性别
    gender: number;
    // 昵称
    name: string;
    // 省/州
    province: string;
    // 备注
    remark: string;
    // 用户 id
    wxid: string;
}

export interface WcfrestDbSqlQueryRequest {
    // 数据库名称
    db: string;
    // 待执行的 SQL
    sql: string;
}

export interface WcfrestDbTablePayload {
    // 表名
    name: string;
    // 建表 SQL
    sql: string;
}

export interface WcfrestDownloadAttachRequest {
    // 消息中的 extra 字段
    extra: string;
    // 消息 id
    msgid: number;
    // 消息中的 thumb 字段
    thumb?: string;
}

export interface WcfrestDownloadImageRequest {
    // 存储路径
    dir: string;
    // 消息中的 extra 字段
    extra: string;
    // 消息 id
    msgid: number;
    // 超时重试次数
    timeout?: number;
}

export interface WcfrestForwardMsgRequest {
    // 待转发消息 id
    id: number;
    // 转发接收人或群的 id 列表
    receiver: Array<string>;
}

export interface WcfrestGetAliasInChatRoomRequest {
    // 群聊 id
    roomid: string;
    // 用户 id
    wxid: string;
}

export interface WcfrestGetAudioMsgRequest {
    // 消息 id
    msgid: number;
    // 存储路径
    path: string;
    // 超时重试次数
    timeout?: number;
}

export interface WcfrestGetAvatarsRequest {
    // 用户 id 列表
    wxids: Array<string>;
}

export interface WcfrestGetChatRoomMembersRequest {
    // 群聊 id
    roomid: string;
}

export interface WcfrestGetDbTablesRequest {
    // 数据库名称
    db: string;
}

export interface WcfrestGetInfoByWxidRequest {
    // 用户 id
    wxid: string;
}

export interface WcfrestGetOcrRequest {
    // 消息中的 extra 字段
    extra: string;
    // 超时重试次数
    timeout?: number;
}

export interface WcfrestReceiveTransferRequest {
    // Transaction id
    taid: string;
    // 转账id transferid
    tfid: string;
    // 转账人
    wxid: string;
}

export interface WcfrestReceiverRequest {
    // 接收推送消息的 url
    url: string;
}

export interface WcfrestRefreshPyqRequest {
    // 分页 id
    id: number;
}

export interface WcfrestRevokeMsgRequest {
    // 消息 id
    msgid: number;
}

export interface WcfrestSendFileRequest {
    // 文件 base64 数据
    base64?: string;
    // 文件路径，若提供 base64 则写入此路径
    path: string;
    // 接收人或群的 id
    receiver: string;
}

export interface WcfrestSendPatMsgRequest {
    // 群 id
    roomid: string;
    // 用户 id
    wxid: string;
}

export interface WcfrestSendRichTextRequest {
    // 填公众号 id 可以显示对应的头像（gh_ 开头的）
    account: string;
    // 摘要，三行
    digest: string;
    // 左下显示的名字
    name: string;
    // 接收人或群的 id
    receiver: string;
    // 缩略图的链接
    thumburl: string;
    // 标题，最多两行
    title: string;
    // 点击后跳转的链接
    url: string;
}

export interface WcfrestSendTxtRequest {
    // 需要 At 的用户 id 列表
    aters?: Array<string>;
    // 消息内容
    msg: string;
    // 接收人或群的 id
    receiver: string;
}

export interface WcfrestUserInfoPayload {
    // 文件/图片等父路径
    home: string;
    // 手机号
    mobile: string;
    // 昵称
    name: string;
    // 用户 id
    wxid: string;
}
