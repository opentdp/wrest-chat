import { Component } from '@angular/core';

import { KeywordGroups, KeywordLevels } from '../../openapi/const';
import { RobotApi, TablesKeyword, KeywordFetchAllParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-keyword-list',
    templateUrl: 'list.html'
})
export class KeywordListComponent {

    public keywordGroups = KeywordGroups;
    public keywordLevels = KeywordLevels;

    public wcfChatrooms: Record<string, WcfrestContactPayload> = {};

    public keywords: Array<TablesKeyword> = [];

    public formdata: KeywordFetchAllParam = {
        group: 'badword',
        roomid: '-',
    };

    constructor() {
        this.getKeywords();
        this.getWcfChatrooms();
    }

    public getKeywords() {
        RobotApi.keywordList(this.formdata).then((data) => {
            this.keywords = data || [];
        });
    }

    public deleteKeyword(item: TablesKeyword) {
        RobotApi.keywordDelete({ rd: item.rd }).then(() => {
            this.getKeywords();
        });
    }

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            data.forEach((item) => this.wcfChatrooms[item.wxid] = item);
        });
    }

}
