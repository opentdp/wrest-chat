import { Component } from '@angular/core';

import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';

import { WcferryContactComponent } from './contact';

@Component({
    selector: 'page-wcferry-chatroom',
    templateUrl: 'chatroom.html',
    styleUrls: ['chatroom.scss']
})
export class WcferryChatroomComponent extends WcferryContactComponent {

    public roomMembers: Record<string, Array<WcfrestContactPayload>> = {};

    public chat = {} as WcfrestContactPayload;

    public conactsFilter = '';

    public changeChat(item: WcfrestContactPayload) {
        if (item.wxid.indexOf('@chatroom') > 0) {
            this.getChatroom(item);
        }
        this.chat = item;
    }

    override getContacts() {
        return super.getContacts().then(() => {
            const c1 = this.contacts.filter((v) => '群聊'.includes(v.type));
            const c2 = this.contacts.filter((v) => '好友'.includes(v.type));
            this.contacts = [...c1, ...c2];
            this.chat = this.contacts[0];
        });
    }

    public getChatroom(room: WcfrestContactPayload) {
        if (this.roomMembers[room.wxid]) {
            return Promise.resolve(); //已获取
        }
        return WrestApi.chatroomMembers({ roomid: room.wxid }).then((data) => {
            this.roomMembers[room.wxid] = (data = data || []);
            this.getAvatars(data.map((v) => v.wxid));
        });
    }

}
