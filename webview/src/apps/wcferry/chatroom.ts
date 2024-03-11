import { Component } from '@angular/core';

import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-wcferry-chatroom',
    templateUrl: 'chatroom.html',
    styleUrls: ['chatroom.scss']
})
export class WcferryChatroomComponent {

    public avatars: Record<string, string> = {};
    public roomMembers: Record<string, Array<WcfrestContactPayload>> = {};

    public chatrooms: Array<WcfrestContactPayload> = [];
    public chatroom!: WcfrestContactPayload;

    public members: Array<WcfrestContactPayload> = [];
    public member!: WcfrestContactPayload;

    constructor() {
        this.getChatrooms();
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
            // 批量获取头像
            const ids = this.members.map((item) => item.wxid);
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

}
