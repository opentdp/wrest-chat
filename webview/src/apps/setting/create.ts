import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { FieldTypes } from '../../openapi/const';
import { RobotApi, TablesLLModel, SettingCreateParam } from '../../openapi/wrobot';


@Component({
    selector: 'page-setting-create',
    templateUrl: 'create.html'
})
export class SettingCreateComponent {

    public fieldTypes = FieldTypes;

    public llmodels: Array<TablesLLModel> = [];

    public formdata: SettingCreateParam = {
        group: 'custom',
        type: 'string',
        name: '',
    };

    constructor(private router: Router) {
        this.getLLModels();
    }

    public createSetting() {
        this.formdata.value = String(this.formdata.value);
        return RobotApi.settingCreate(this.formdata).then(() => {
            this.router.navigate(['setting/list']);
        });
    }

    public getLLModels() {
        return RobotApi.llmodelList({}).then((data) => {
            this.llmodels = data || [];
        });
    }
}
