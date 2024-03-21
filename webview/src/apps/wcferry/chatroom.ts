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

    constructor() {
        this.getChatrooms();
    }

    public getChatrooms() {
        return WrestApi.chatrooms().then((data) => {
            this.chatrooms = data || [];
            this.getAvatars(this.chatrooms.map((v) => v.wxid));
        });
    }

    public getChatroom(room: WcfrestContactPayload) {
        this.chatroom = room;
        if (this.roomMembers[room.wxid]) {
            this.members = this.roomMembers[room.wxid];
            return; // 已获取
        }
        return WrestApi.chatroomMembers({ roomid: room.wxid }).then((data) => {
            this.roomMembers[room.wxid] = data || [];
            this.members = data || [];
            // 批量获取头像
            this.getAvatars(this.members.map((v) => v.wxid));
        });
    }

    public getAvatars(ids: string[]) {
        return WrestApi.avatars({ wxids: [...new Set(ids)] }).then((data) => {
            data && data.forEach((item) => {
                this.avatars[item.usr_name] = item.small_head_img_url;
            });
        });
    }

}
