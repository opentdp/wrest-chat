import { Component } from '@angular/core';

import { UserLevels, SpecialRooms } from '../../openapi/const';
import { SundryApi, CronjobPlugin, KeywordPlugin } from '../../openapi/sundry';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-plugin-list',
    templateUrl: 'list.html'
})
export class PluginListComponent {

    public userLevels = UserLevels;
    public specialRooms = SpecialRooms;

    public type = 'cronjob';
    public cronjobPlugins: Array<CronjobPlugin> = [];
    public keywordPlugins: Array<KeywordPlugin> = [];

    public wcfChatrooms: Record<string, WcfrestContactPayload> = {};

    constructor() {
        this.getWcfChatrooms();
        this.getCronjobPlugins();
        this.getKeywordPlugins();
    }

    public getCronjobPlugins() {
        return SundryApi.pluginCronjobs({}).then((data) => {
            this.cronjobPlugins = data || [];
        });
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
