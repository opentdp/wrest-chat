import { Component } from '@angular/core';

import { KeywordLevels } from 'src/openapi/const';
import { RobotApi, TablesKeyword } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-keyword-list',
    templateUrl: 'list.html',
    styleUrls: ['list.scss']
})
export class KeywordListComponent {

    public keywordLevels = KeywordLevels;

    public wcfChatrooms: Record<string, WcfrestContactPayload> = {};

    public keywords: Array<TablesKeyword> = [];

    constructor() {
        this.getKeywords();
        this.getWcfChatrooms();
    }

    public getKeywords() {
        RobotApi.keywordList({}).then((data) => {
            this.keywords = data || [];
        });
    }

    public deleteKeyword(item: TablesKeyword) {
        const rq = { phrase: item.phrase, roomid: item.roomid };
        RobotApi.keywordDelete(rq).then(() => {
            this.getKeywords();
        });
    }

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            data.forEach((item) => this.wcfChatrooms[item.wxid] = item);
        });
    }

}
