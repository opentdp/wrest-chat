import { Component } from '@angular/core';

import { WrestApi, WcfrestUserInfoPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-owner',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class OwnerComponent {

    public user!: WcfrestUserInfoPayload;

    constructor() {
        this.getSelfInfo();
    }

    public getSelfInfo() {
        WrestApi.selfInfo().then((user) => {
            this.user = user;
        });
    }

}
