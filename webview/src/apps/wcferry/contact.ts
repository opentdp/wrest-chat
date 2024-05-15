import { Component } from '@angular/core';

import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-wcferry-contact',
    templateUrl: 'contact.html'
})
export class WcferryContactComponent {

    public avatars: Record<string, string> = {};
    public contacts: Array<WcfrestContactPayload & { type: string }> = [];

    public contactTypes: Record<string, RegExp> = {
        '群聊': /@chatroom$/,
        '公众号': /^gh_/,
        '企业微信': /@openim$/,
        '内置服务': /^fmessage|filehelper|floatbottle|medianote|mphelper|newsapp$/,
        '可能异常': /[\n\r]$/,
    };

    public contactType = '好友';

    constructor() {
        this.getContacts();
    }

    public getContacts() {
        return WrestApi.contacts().then((data) => {
            this.contacts = data.map(item => ({
                ...item, type: this.getContactType(item)
            }));
            this.getAvatars(this.contacts.map((v) => v.wxid));
        });
    }

    public getContactType(contact: WcfrestContactPayload) {
        for (const type in this.contactTypes) {
            const regex = this.contactTypes[type];
            if (regex && regex.test(String(contact.wxid))) {
                return type;
            }
        }
        return '好友';
    }

    public getAvatars(ids: string[]) {
        const wxids = [...new Set(ids)];
        return WrestApi.avatars({ wxids }).then((data) => {
            data && data.forEach((item) => {
                this.avatars[item.usr_name] = item.small_head_img_url;
            });
        });
    }

}
