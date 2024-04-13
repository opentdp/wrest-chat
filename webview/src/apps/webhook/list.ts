import { Component } from '@angular/core';

import { RobotApi, Webhook } from '../../openapi/wrobot';
import { WcfrestContactPayload, WrestApi } from "../../openapi/wcfrest";


@Component({
    selector: 'page-webhook-list',
    templateUrl: 'list.html'
})
export class WebhookListComponent {
    public list: Array<Webhook> = [];

    public wcfContacts: Record<string, WcfrestContactPayload> = {};
    public wcfChatrooms: Record<string, WcfrestContactPayload> = {};
    public wcfRoomMembers: Record<string, Record<string, WcfrestContactPayload>> = {};

    constructor() {
        this.getWebhookList();
        this.getWcfContacts();
        this.getWcfChatrooms();
    }

    public getWebhookList() {
        return RobotApi.webhookList().then((data) => {
            this.list = data || [];
        });
    }

    public deleteWebhook(item: Webhook) {
        return RobotApi.webhookDelete({ rd: item.rd }).then(() => {
            this.getWebhookList();
        });
    }

    public getWcfContacts() {
        return WrestApi.contacts().then((data) => {
            data.forEach((item) => this.wcfContacts[item.wxid] = item);
        });
    }

    public getWcfChatrooms() {
        return WrestApi.chatrooms().then((data) => {
            data.forEach((item) => this.wcfChatrooms[item.wxid] = item);
        });
    }

    public getWcfRoomMembers(id: string) {
        if (this.wcfRoomMembers[id]) {
            return Promise.resolve(); //已获取
        }
        this.wcfRoomMembers[id] = {}; //初始化
        return WrestApi.chatroomMembers({ roomid: id }).then((data) => {
            data && data.forEach((item) => {
                this.wcfRoomMembers[id][item.wxid] = item;
            });
        });
    }

}
