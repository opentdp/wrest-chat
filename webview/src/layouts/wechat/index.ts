import { Component, OnDestroy, Input, ViewChild, ElementRef } from '@angular/core';

import { WrestApi, WcfrestContactPayload, WcfrestUserInfoPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'layout-wechat',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class LayoutWechatComponent implements OnDestroy {

    @ViewChild('scrollLayout')
    private scrollLayout!: ElementRef;

    public wss!: WebSocket;
    public wsMsg: Array<IMessage> = [];

    public messages: Array<IMessage> = [];
    public self = {} as WcfrestUserInfoPayload;

    public chat!: WcfrestContactPayload;
    public memberMap: Record<string, WcfrestContactPayload> = {};
    public memberCount: number = 0;

    public showPanel = false;
    public isGroup = false;
    public content = '';

    @Input()
    public avatars: Record<string, string> = {};

    @Input()
    public set value(val: WcfrestContactPayload) {
        if (!val || !val.wxid) {
            return;
        }
        this.chat = val;
        this.memberCount = 0;
        this.memberMap[val.wxid] = val;
        this.isGroup = val.wxid.includes('@chatroom');
        this.isGroup ? this.getChatroom(val.wxid) : this.getAvatars([val.wxid]);
        this.messages = this.wsMsg.filter(v => this.pickupMsg(v));
    }

    public constructor() {
        this.restoreMsg();
        this.avatars['system'] = '/assets/icon.png';
        WrestApi.selfInfo().then((data) => {
            this.self = data;
            this.startSocket();
        });
    }

    public ngOnDestroy() {
        this.wss && this.wss.close();
        this.messages = [];
    }

    public sendTxt() {
        const rq = {
            msg: this.content,
            receiver: this.chat.wxid,
        };
        return WrestApi.sendTxt(rq).then(() => {
            const msg: IMessage = {
                id: 0,
                type: 1,
                ts: Date.now(),
                sender: this.self.wxid,
                is_group: this.isGroup,
                roomid: this.isGroup ? this.chat.wxid : '',
                receiver: this.chat.wxid,
                content: this.content,
                sign: '',
            };
            this.storeMsg(msg);
            this.content = '';
        });
    }

    public storeMsg(msg: IMessage) {
        this.wsMsg.push(msg);
        if (this.pickupMsg(msg)) {
            this.messages.push(msg);
        }
        sessionStorage.setItem('wx::message', JSON.stringify(this.wsMsg));
    }

    public restoreMsg() {
        const str = sessionStorage.getItem('wx::message') || '[]';
        this.wsMsg = JSON.parse(str) as IMessage[];
        this.getAvatars(this.wsMsg.map(v => v.sender));
    }

    public pickupMsg(msg: IMessage) {
        if (!msg || !this.chat) {
            return false;
        }
        // 滚动
        setTimeout(() => {
            const el = this.scrollLayout.nativeElement;
            el.scrollTop = el.scrollHeight;
        }, 100);
        // 头像
        if (this.avatars[msg.sender] === undefined) {
            this.avatars[msg.sender] = '/assets/icon.png';
            this.getAvatars([msg.sender]);
        }
        // 群聊
        if (this.isGroup) {
            return this.chat.wxid === msg.roomid;
        }
        // 私聊
        return !msg.is_group && (
            msg.sender == 'system' ||
            msg.sender == this.chat.wxid ||
            msg.receiver == this.chat.wxid
        );
    }

    public startSocket() {
        const token = sessionStorage.getItem('token');
        const url = location.origin.replace(/^http/, 'ws') + '/wcf/socket_receiver';
        const wss = new WebSocket(url + (token ? '?token=' + token : ''));
        // 接收消息
        wss.onmessage = (event) => {
            const msg = JSON.parse(event.data) as IMessage;
            if (msg && msg.ts > 0) {
                msg.ts = msg.ts * 1000;
                this.storeMsg(msg);
            }
        };
        // 连接成功
        wss.onopen = () => {
            this.wss = wss;
            const msg = {
                ts: Date.now(),
                sender: 'system',
                is_group: this.isGroup,
                roomid: this.isGroup ? this.chat.wxid : '',
                content: 'websocket is connected',
            };
            this.storeMsg(msg as IMessage);
        };
        // 自动重连
        wss.onclose = () => {
            const msg = {
                ts: Date.now(),
                sender: 'system',
                is_group: this.isGroup,
                roomid: this.isGroup ? this.chat.wxid : '',
                content: 'websocket is closed, retry in 5s',
            };
            this.storeMsg(msg as IMessage);
            setTimeout(() => this.startSocket(), 5 * 1000);
        };
        // 捕获错误
        wss.onerror = (event) => {
            const msg = {
                ts: Date.now(),
                sender: 'system',
                is_group: this.isGroup,
                roomid: this.isGroup ? this.chat.wxid : '',
                content: 'websocket error, details to console',
            };
            console.log(event);
            this.storeMsg(msg as IMessage);
            setTimeout(() => this.startSocket(), 5 * 1000);
        };
    }

    public getAvatars(ids: string[]) {
        const wxids = [...new Set(ids)];
        return WrestApi.avatars({ wxids }).then((data) => {
            data && data.forEach((item) => {
                this.avatars[item.usr_name] = item.small_head_img_url;
            });
        });
    }

    public getChatroom(roomid: string) {
        return WrestApi.chatroomMembers({ roomid }).then((data) => {
            data.map((v) => this.memberMap[v.wxid] = v);
            this.memberCount = data.length;
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
    receiver: string;
    sign: string;
    content: string;
    xml?: {
        msgsource: {
            atuserlist: string;
        };
    };
}
