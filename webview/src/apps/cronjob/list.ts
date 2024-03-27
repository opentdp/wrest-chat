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
    public status: Record<number, CronjobStatusPayload> = {};

    public execStatus: Record<number, boolean> = {};

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

    public executeCronjob(item: TablesCronjob) {
        this.execStatus[item.rd] = true;
        const rq = { rd: item.rd };
        return SundryApi.cronjobExecute(rq).finally(() => {
            this.execStatus[item.rd] = false;
        });
    }

    public deleteCronjob(item: TablesCronjob) {
        const rq = { rd: item.rd };
        return SundryApi.cronjobDelete(rq).then(() => {
            this.getCronjobs();
        });
    }

}
