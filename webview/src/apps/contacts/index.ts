import { Component } from '@angular/core';

import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-contacts',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class ContactsComponent {

    public contacts: Array<WcfrestContactPayload & { type: string }> = [];

    public contactTypes: Record<string, RegExp> = {
        '公众平台助手': /^mphelper$/,
        '朋友推荐消息': /^fmessage$/,
        '语音记事本': /^medianote$/,
        '漂流瓶': /^floatbottle$/,
        '文件传输助手': /^filehelper$/,
        '新闻': /^newsapp$/,
        '公众号': /^gh_/,
        '群聊': /@chatroom$/,
        '企业微信': /@openim$/,
    };

    constructor() {
        this.getContacts();
    }

    public getContacts() {
        WrestApi.contacts().then((contacts) => {
            this.contacts = contacts.map(contact => ({
                ...contact, type: this.getContactType(contact)
            }));
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

}
