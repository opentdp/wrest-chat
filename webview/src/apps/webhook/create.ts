import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { UserLevels } from '../../openapi/const';
import {RobotApi, WebhookCreateWebHookParam} from '../../openapi/wrobot';
import {WcfrestContactPayload, WrestApi} from "../../openapi/wcfrest";


@Component({
    selector: 'page-webhook-create',
    templateUrl: 'create.html'
})
export class WebhookCreateComponent {

    public wcfFriends: Array<WcfrestContactPayload> = [];
    public wcfChatrooms: Array<WcfrestContactPayload> = [];
    public filter = '';

    public formdata: WebhookCreateWebHookParam = {
        remark: "",
        targetId: "",
        target: "friend",
    }

    constructor(private router: Router) {
        this.getWcfFriends();
        this.getWcfChatrooms();
    }

    public createWebhook() {
        return RobotApi.webhookCreate(this.formdata).then(() => {
            this.router.navigate(['webhook/list']);
        })
    }

    public getWcfFriends() {
        return WrestApi.friends().then((data) => {
            this.wcfFriends = data || [];
        });
    }

    public getWcfChatrooms() {
        return WrestApi.chatrooms().then((data) => {
            this.wcfChatrooms = data || [];
        });
    }

}
