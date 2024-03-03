import { Component } from '@angular/core';

import { RobotApi, TablesSetting } from '../../openapi/wrobot';


@Component({
    selector: 'page-setting-list',
    templateUrl: 'list.html'
})
export class SettingListComponent {

    public settings: Array<TablesSetting> = [];

    constructor() {
        this.getSettings();
    }

    public getSettings() {
        RobotApi.settingList({}).then((data) => {
            this.settings = data || [];
        });
    }

}
