import { Component } from '@angular/core';

import { WrestApi, WcferryRpcContact } from '../../openapi/wcfrest';


@Component({
    selector: 'page-chats',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class ChatsComponent {

    public chatrooms: Array<WcferryRpcContact> = [];

    constructor() {
        WrestApi.chatrooms().then((chatrooms) => {
            this.chatrooms = chatrooms;
        });
    }

}
