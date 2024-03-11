import { Component } from '@angular/core';

import { FieldTypes } from '../../openapi/const';
import { RobotApi, TablesSetting } from '../../openapi/wrobot';


@Component({
    selector: 'page-setting-list',
    templateUrl: 'list.html'
})
export class SettingListComponent {

    public fieldTypes = FieldTypes;

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
