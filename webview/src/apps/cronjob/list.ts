import { Component } from '@angular/core';

import { CronjobTypes } from '../../openapi/const';
import { SundryApi, CronjobStatusPayload, TablesCronjob } from '../../openapi/sundry';


@Component({
    selector: 'page-cronjob-list',
    templateUrl: 'list.html'
})
export class CronjobListComponent {

    public cronjobTypes = CronjobTypes;

    public cronjobs: Array<TablesCronjob> = [];
    public status: Record<number, CronjobStatusPayload> = [];

    constructor() {
        this.getCronjobs();
        this.getCronStatus();
    }

    public getCronjobs() {
        SundryApi.cronjobList({}).then((data) => {
            this.cronjobs = data || [];
        });
    }

    public getCronStatus() {
        SundryApi.cronjobStatus({}).then((data) => {
            this.status = data || [];
        });
    }

    public deleteCronjob(item: TablesCronjob) {
        const rq = { rd: item.rd };
        SundryApi.cronjobDelete(rq).then(() => {
            this.getCronjobs();
        });
    }

}
