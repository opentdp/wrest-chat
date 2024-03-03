import { Component } from '@angular/core';

import { CronApi, CronjobStatusPayload, TablesCronjob } from '../../openapi/cronjob';


@Component({
    selector: 'page-cronjob-list',
    templateUrl: 'list.html',
    styleUrls: ['list.scss']
})
export class CronjobListComponent {

    public cronjobs: Array<TablesCronjob> = [];
    public status: Record<number, CronjobStatusPayload> = [];

    constructor() {
        this.getCronjobs();
        this.getCronStatus();
    }

    public getCronjobs() {
        CronApi.cronjobList({}).then((data) => {
            this.cronjobs = data || [];
        });
    }

    public getCronStatus() {
        CronApi.cronjobStatus({}).then((data) => {
            this.status = data || [];
        });
    }

    public deleteCronjob(item: TablesCronjob) {
        const rq = { rd: item.rd };
        CronApi.cronjobDelete(rq).then(() => {
            this.getCronjobs();
        });
    }

}
