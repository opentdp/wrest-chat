import { Component, OnDestroy, Input } from '@angular/core';

import { WrestApi, WcfrestContactPayload, WcfrestUserInfoPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'layout-wechat',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class LayoutWechatComponent implements OnDestroy {

    public wss!: WebSocket;

    public messages: Array<IMessage> = [];
    public avatars: Record<string, string> = {};
    public self = {} as WcfrestUserInfoPayload;

    public chatroom!: WcfrestContactPayload;
    public memberMap: Record<string, WcfrestContactPayload> = {};
    public memberCount: number = 0;

    public subPanel = false;
    public talkId = '';

    public content = '';

    @Input()
    public set chat(val: WcfrestContactPayload) {
        this.chatroom = val;
        if (this.chatroom) {
            this.getChatroom(this.chatroom.wxid);
        }
    }

    public constructor() {
        this.startSocket();
        WrestApi.selfInfo().then((data) => {
            this.self = data;
        });
    }

    public ngOnDestroy() {
        this.wss && this.wss.close();
        this.messages = [];
    }

    public sendTxt() {
        const rq = {
            msg: this.content,
            receiver: this.chatroom.wxid,
        };
        return WrestApi.sendTxt(rq).then(() => {
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

    public startSocket() {
        // 消息本地缓存
        const sid = 'wx::message'
        const msg = sessionStorage.getItem(sid) || '[]';
        this.messages = JSON.parse(msg) as IMessage[];
        // 注册消息接收
        const token = sessionStorage.getItem('token');
        const url = location.origin.replace(/^http/, 'ws') + '/wcf/socket_receiver';
        const wss = new WebSocket(url + (token ? '?token=' + token : ''));
        // 连接成功
        wss.onopen = () => {
            this.wss = wss;
            const data = {
                ts: Date.now(),
                sender: 'system',
                content: 'websocket is connected'
            };
            this.messages.push(data as IMessage);
        };
        // 自动重连
        wss.onclose = () => {
            const data = {
                ts: Date.now(),
                sender: 'system',
                content: 'websocket is closed, retry in 5s'
            };
            this.messages.push(data as IMessage);
            setTimeout(() => this.startSocket(), 5 * 1000);
        };
        // 捕获错误
        wss.onerror = (event) => {
            const data = {
                ts: Date.now(),
                sender: 'system',
                content: 'websocket error, details to console'
            };
            this.messages.push(data as IMessage);
            console.log(event);
        };
        // 接收消息
        wss.onmessage = (event) => {
            const data = JSON.parse(event.data) as IMessage;
            if (data && data.ts > 0) {
                data.ts = data.ts * 1000;
                this.messages.push(data);
                // 获取头像
                this.getAvatar(data.sender);
                // 消息本地缓存
                sessionStorage.setItem(sid, JSON.stringify(this.messages));
            }
        };
    }

    public getAvatar(id: string) {
        if (this.avatars[id]) {
            return Promise.resolve(this.avatars[id]);
        }
        return WrestApi.avatars({ wxids: [id] }).then((data) => {
            data && data.forEach((item) => {
                this.avatars[item.usr_name] = item.small_head_img_url;
            });
            return this.avatars[id] || '/assets/icon.png';
        });
    }

    public getChatroom(roomid: string) {
        return WrestApi.chatroomMembers({ roomid }).then((data) => {
            this.memberCount = data.length;
            data.map((item) => {
                this.memberMap[item.wxid] = item;
            });
        });
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
