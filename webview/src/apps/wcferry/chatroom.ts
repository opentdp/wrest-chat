import { Component, OnDestroy } from '@angular/core';

import { WrestApi, WcfrestContactPayload, WcfrestUserInfoPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-wcferry-chatroom',
    templateUrl: 'chatroom.html',
    styleUrls: ['chatroom.scss']
})
export class WcferryChatroomComponent implements OnDestroy {

    public avatars: Record<string, string> = {};
    public roomMembers: Record<string, Array<WcfrestContactPayload>> = {};

    public chatrooms: Array<WcfrestContactPayload> = [];
    public chatroom!: WcfrestContactPayload;

    public memberMap: Record<string, WcfrestContactPayload> = {};
    public members: Array<WcfrestContactPayload> = [];
    public member!: WcfrestContactPayload;

    constructor() {
        this.getChatrooms();
        this.startSocket();
    }

    public ngOnDestroy() {
        this.stopSocket();
    }

    public getChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.chatrooms = data || [];
            // 批量获取头像
            const ids = this.chatrooms.map((item) => item.wxid);
            this.getAvatars(ids);
        });
    }

    public getChatroom(room: WcfrestContactPayload) {
        this.chatroom = room;
        if (this.roomMembers[room.wxid]) {
            this.members = this.roomMembers[room.wxid];
            return; // 已获取
        }
        WrestApi.chatroomMembers({ roomid: room.wxid }).then((data) => {
            this.roomMembers[room.wxid] = data || [];
            this.members = data || [];
            // 更新会员列表
            const ids = this.members.map((item) => {
                this.memberMap[item.wxid] = item;
                return item.wxid;
            });
            // 批量获取头像
            this.getAvatars(ids);
        });
    }

    public getAvatars(ids: string[]) {
        WrestApi.avatars({ wxids: [...new Set(ids)] }).then((data) => {
            data && data.forEach((item) => {
                this.avatars[item.usr_name] = item.small_head_img_url;
            });
        });
    }

    // 聊天记录

    public wss!: WebSocket;
    public messages: Array<IMessage> = [];
    public self = {} as WcfrestUserInfoPayload;

    public showMember = false;
    public content = '';

    public sendTxt() {
        const rq = {
            msg: this.content,
            receiver: this.chatroom.wxid,
        };
        WrestApi.sendTxt(rq).then(() => {
            const msg: IMessage = {
                ts: Date.now(),
                roomid: this.chatroom.wxid,
                sender: this.self.wxid,
                content: this.content,
                is_group: true,
                id: 0,
                type: 1,
                sign: '',
            };
            this.messages.push(msg);
            this.content = '';
        });
    }

    public stopSocket() {
        this.wss && this.wss.close();
        this.messages = [];
    }

    public startSocket() {
        WrestApi.selfInfo().then((data) => {
            this.self = data;
        });
        // 注册消息接收
        const token = sessionStorage.getItem('token');
        const url = location.origin.replace(/^http/, 'ws') + '/wcf/socket_receiver';
        const wss = new WebSocket(url + (token ? '?token=' + token : ''));
        wss.onopen = () => {
            const data = {
                ts: Date.now(),
                sender: 'system',
                content: 'websocket is connected'
            };
            this.messages.push(data as IMessage);
            this.wss = wss;
        };
        wss.onclose = () => {
            const data = {
                ts: Date.now(),
                sender: 'system',
                content: 'websocket is closed, retry in 5s'
            };
            this.messages.push(data as IMessage);
            setTimeout(() => this.startSocket(), 5 * 1000);
        };
        wss.onerror = (event) => {
            const data = {
                ts: Date.now(),
                sender: 'system',
                content: 'websocket error, details to console'
            };
            this.messages.push(data as IMessage);
            console.log(event);
        };
        wss.onmessage = (event) => {
            const data = JSON.parse(event.data) as IMessage;
            data.ts = data.ts * 1000;
            this.messages.push(data);
        };
    }

}

interface IMessage {
    is_group: boolean;
    id: number;
    type: number;
    ts: number;
    roomid: string;
    sender: string;
    sign: string;
    content: string;
    xml?: {
        msgsource: {
            atuserlist: string;
        };
    };
}
