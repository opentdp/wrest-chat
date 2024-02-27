import { Component } from '@angular/core';

import { RobotApi, ProfileFetchAllParam, TablesProfile } from '../../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../../openapi/wcfrest';


@Component({
    selector: 'page-bot-profiles',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class BotProfilesComponent {

    public contacts: Array<WcfrestContactPayload> = [];

    public profiles: Array<TablesProfile> = [];

    constructor() {
        this.getContacts();
        this.getProfiles();
    }

    public getContacts() {
        WrestApi.contacts().then((data) => {
            this.contacts = data;
        });
    }

    public getProfiles() {
        RobotApi.profileList({} as ProfileFetchAllParam).then((data) => {
            this.profiles = data;
        });
    }

}
