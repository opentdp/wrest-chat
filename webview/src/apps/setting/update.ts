import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import { FieldTypes } from '../../openapi/const';
import { RobotApi, TablesLLModel, SettingUpdateParam } from '../../openapi/wrobot';


@Component({
    selector: 'page-setting-update',
    templateUrl: 'update.html'
})
export class SettingUpdateComponent implements OnInit {

    public fieldTypes = FieldTypes;

    public llmodels: Array<TablesLLModel> = [];

    public formdata = {} as SettingUpdateParam;

    constructor(
        private router: Router,
        private route: ActivatedRoute
    ) {
        this.getLLModels();
    }

    public ngOnInit() {
        const name = this.route.snapshot.paramMap.get('name');
        name && this.getSetting(name);
    }

    public getSetting(name: string) {
        RobotApi.settingDetail({ name }).then((data) => {
            data && Object.assign(this.formdata, data);
        });
    }

    public updateSetting() {
        this.formdata.value = String(this.formdata.value);
        RobotApi.settingUpdate(this.formdata).then(() => {
            this.router.navigate(['setting/list']);
        });
    }

    public getLLModels() {
        RobotApi.llmodelList({}).then((data) => {
            this.llmodels = data || [];
        });
    }
}
