import { Component } from '@angular/core';

import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';

import { WcferryContactComponent } from './contact';

@Component({
    selector: 'page-wcferry-chatroom',
    templateUrl: 'chatroom.html',
    styleUrls: ['chatroom.scss']
})
export class WcferryChatroomComponent extends WcferryContactComponent {

    public avatars: Record<string, string> = {};
    public roomMembers: Record<string, Array<WcfrestContactPayload>> = {};

    public chat = {} as WcfrestContactPayload;

    public conactsFilter = '';

    override getContacts() {
        return super.getContacts().then(() => {
            this.contacts = this.contacts.filter((v) => '群聊,好友'.includes(v.type));
            this.getAvatars(this.contacts.map((v) => v.wxid));
            this.chat = this.contacts[0];
        });
    }

    public changeChat(item: WcfrestContactPayload) {
        this.chat = item;
        item.wxid = item.wxid.trim();
        if (item.wxid.indexOf('@chatroom') > 0) {
            this.getChatroom(item);
        }
    }

    public getChatroom(room: WcfrestContactPayload) {
        if (this.roomMembers[room.wxid]) {
            return; // 已获取
        }
        return WrestApi.chatroomMembers({ roomid: room.wxid }).then((data) => {
            this.roomMembers[room.wxid] = (data = data || []);
            this.getAvatars(data.map((v) => v.wxid));
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
