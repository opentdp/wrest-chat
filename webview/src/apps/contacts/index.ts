import { Component } from '@angular/core';

import { WrestApi, WcferryRpcContact } from '../../openapi/wcfrest';


@Component({
    selector: 'page-contacts',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class ContactsComponent {

    public contacts: Array<WcferryRpcContact> = [];

    constructor() { 
        WrestApi.contacts().then((contacts) => {
            this.contacts = contacts;
        });
    }

}
