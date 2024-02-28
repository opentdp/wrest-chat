import { Component } from '@angular/core';

import { RobotApi, KeywordFetchAllParam, TablesKeyword } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-keyword-list',
    templateUrl: 'list.html',
    styleUrls: ['list.scss']
})
export class KeywordListComponent {

    public contacts: Record<string, WcfrestContactPayload> = {};

    public keywords: Array<TablesKeyword> = [];

    constructor() {
        this.getContacts();
        this.getKeywords();
    }

    public getContacts() {
        WrestApi.contacts().then((data) => {
            data.forEach((item) => this.contacts[item.wxid] = item);
        });
    }

    public getKeywords() {
        const rq = {} as KeywordFetchAllParam;
        RobotApi.keywordList(rq).then((data) => {
            this.keywords = data || [];
        });
    }

    public deleteLLModel(item: TablesKeyword) {
        const rq = { phrase: item.phrase, roomid: item.roomid };
        RobotApi.keywordDelete(rq).then(() => {
            this.getKeywords();
        });
    }

}
