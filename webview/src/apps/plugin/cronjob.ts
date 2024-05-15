import { Component } from '@angular/core';

import { UserLevels, SpecialRooms } from '../../openapi/const';
import { SundryApi, CronjobPlugin } from '../../openapi/sundry';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-plugin-cronjob',
    templateUrl: 'cronjob.html'
})
export class PluginCronjobComponent {

    public userLevels = UserLevels;
    public specialRooms = SpecialRooms;

    public cronjobPlugins: Array<CronjobPlugin> = [];

    public wcfChatrooms: Record<string, WcfrestContactPayload> = {};

    constructor() {
        this.getWcfChatrooms();
        this.getCronjobPlugins();
    }

    public getCronjobPlugins() {
        return SundryApi.pluginCronjobs({}).then((data) => {
            this.cronjobPlugins = data || [];
        });
    }

    public getWcfChatrooms() {
        return WrestApi.chatrooms().then((data) => {
            data.forEach((item) => this.wcfChatrooms[item.wxid] = item);
        });
    }

}
