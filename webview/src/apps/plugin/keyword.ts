import { Component } from '@angular/core';

import { UserLevels, SpecialRooms } from '../../openapi/const';
import { SundryApi, KeywordPlugin } from '../../openapi/sundry';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-plugin-keyword',
    templateUrl: 'keyword.html'
})
export class PluginKeywordComponent {

    public userLevels = UserLevels;
    public specialRooms = SpecialRooms;

    public keywordPlugins: Array<KeywordPlugin> = [];

    public wcfChatrooms: Record<string, WcfrestContactPayload> = {};

    constructor() {
        this.getWcfChatrooms();
        this.getKeywordPlugins();
    }

    public getKeywordPlugins() {
        return SundryApi.pluginKeywords({}).then((data) => {
            this.keywordPlugins = data || [];
        });
    }

    public getWcfChatrooms() {
        return WrestApi.chatrooms().then((data) => {
            data.forEach((item) => this.wcfChatrooms[item.wxid] = item);
        });
    }

}
