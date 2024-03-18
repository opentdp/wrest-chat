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
        return SundryApi.cronjobList({}).then((data) => {
            this.cronjobs = data || [];
        });
    }

    public getCronStatus() {
        return SundryApi.cronjobStatus({}).then((data) => {
            this.status = data || [];
        });
    }

    public deleteCronjob(item: TablesCronjob) {
        const rq = { rd: item.rd };
        return SundryApi.cronjobDelete(rq).then(() => {
            this.getCronjobs();
        });
    }

}
