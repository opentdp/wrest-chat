import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { RobotApi, TablesKeyword } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-keyword-create',
    templateUrl: 'create.html',
    styleUrls: ['create.scss']
})
export class KeywordCreateComponent {

    public chatrooms: Array<WcfrestContactPayload> = [];

    public keywords: Array<TablesKeyword> = [];

    public formdata = { level: 1, roomid: '-' } as TablesKeyword;

    constructor(private router: Router) {
        this.getChatrooms();
    }

    public getChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.chatrooms = data || [];
        });
    }

    public createKeyword() {
        RobotApi.keywordCreate(this.formdata).then(() => {
            this.router.navigate(['keyword/list']);
        });
    }

}
