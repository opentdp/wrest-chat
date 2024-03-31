import { Component } from '@angular/core';

import {RobotApi, Webhook} from '../../openapi/wrobot';


@Component({
    selector: 'page-webhook-list',
    templateUrl: 'list.html'
})
export class WebhookListComponent {
    public list: Array<Webhook> = [];

    constructor() {
        this.getWebhookList();
    }

    public getWebhookList() {
        return RobotApi.webhookList().then((data) => {
            this.list = data || [];
        });
    }

    public deleteWebhook(item: Webhook) {
        return RobotApi.webhookDelete({ rd: item.rd }).then(() => {
            this.getWebhookList();
        })
    }

}
