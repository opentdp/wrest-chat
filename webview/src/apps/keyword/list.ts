import { Component } from '@angular/core';

import { UserLevels, SpecialRooms, KeywordGroups, BadwordLevels } from '../../openapi/const';
import { RobotApi, TablesKeyword, KeywordFetchAllParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-keyword-list',
    templateUrl: 'list.html'
})
export class KeywordListComponent {

    public userLevels = UserLevels;
    public specialRooms = SpecialRooms;
    public keywordGroups = KeywordGroups;
    public badwordLevels = BadwordLevels;

    public wcfChatrooms: Record<string, WcfrestContactPayload> = {};

    public keywords: Array<TablesKeyword> = [];

    public formdata: KeywordFetchAllParam = {
        group: '',
        roomid: '',
    };

    constructor() {
        this.getKeywords();
        this.getWcfChatrooms();
    }

    public getKeywords() {
        return RobotApi.keywordList(this.formdata).then((data) => {
            this.keywords = data || [];
        });
    }

    public deleteKeyword(item: TablesKeyword) {
        return RobotApi.keywordDelete({ rd: item.rd }).then(() => {
            this.getKeywords();
        });
    }

    public getWcfChatrooms() {
        return WrestApi.chatrooms().then((data) => {
            data.forEach((item) => this.wcfChatrooms[item.wxid] = item);
        });
    }

}
