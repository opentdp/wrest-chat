import { Component } from '@angular/core';

import { WrestApi, WcfrestDbTablePayload } from '../../openapi/wcfrest';

@Component({
    selector: 'page-wcferry-dbquery',
    templateUrl: 'dbquery.html'
})
export class WcferryDbqueryComponent {

    public dbList: Array<string> = [];
    public tableList: Array<WcfrestDbTablePayload> = [];

    public dbName = '';
    public tableName = '';

    public sql = 'SELECT * FROM {TABLE} LIMIT 10';
    public result = '';

    public loading = false;

    constructor() {
        this.getDbList();
    }

    public getDbList() {
        return WrestApi.dbNames().then((data) => {
            this.dbList = data;
            this.dbName = '';
        });
    }

    public getDbTables() {
        return WrestApi.dbTables({ db: this.dbName }).then((data) => {
            this.tableList = data;
            this.tableName = '';
        });
    }

    public getDbTableRecords() {
        this.loading = true;
        const db = this.dbName;
        const sql = this.sql.replace(/\{TABLE\}/, this.tableName);
        return WrestApi.dbQuerySql({ db, sql }).then((data) => {
            this.result = JSON.stringify(data, null, 4);
        }).finally(() => {
            this.loading = false;
        });
    }

}
