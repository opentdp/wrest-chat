import { Component } from '@angular/core';

import { WrestApi, WcferryRpcContact } from '../../openapi/wcfrest';


@Component({
    selector: 'page-contacts',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class ContactsComponent {

    public contacts: Array<WcferryRpcContact> = [];

    public contactTypes: Record<string, RegExp> = {
        '公众平台助手': /^mphelper$/,
        '朋友推荐消息': /^fmessage$/,
        '语音记事本': /^medianote$/,
        '漂流瓶': /^floatbottle$/,
        '文件传输助手': /^filehelper$/,
        '新闻': /^newsapp$/,
        '公众号': /^gh_/,
        '群聊': /@chatroom$/,
        '企微好友': /@openim$/,
    };

    constructor() {
        WrestApi.contacts().then((contacts) => {
            this.contacts = contacts;
        });
    }

    public getContactType(contact: WcferryRpcContact) {
        for (const type in this.contactTypes) {
            const regex = this.contactTypes[type];
            if (regex && regex.test(String(contact.wxid))) {
                return type;
            }
        }
        return '好友';
    }

}
