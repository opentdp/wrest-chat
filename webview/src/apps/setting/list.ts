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
        return RobotApi.settingList({}).then((data) => {
            this.settings = data || [];
        });
    }

    public deleteSetting(item: TablesSetting) {
        if (item.group === 'bot') {
            window.postMessage({ message: '系统配置不可删除', type: 'danger' });
            return Promise.resolve();
        }
        return RobotApi.settingDelete({ rd: item.rd }).then(() => {
            this.getSettings();
        });
    }

}
