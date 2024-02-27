import { Component } from '@angular/core';

import { RobotApi, ProfileFetchAllParam, TablesProfile } from '../../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../../openapi/wcfrest';


@Component({
    selector: 'page-bot-profiles',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class BotProfilesComponent {

    public contacts: Record<string, WcfrestContactPayload> = {};

    public profiles: Array<TablesProfile> = [];

    constructor() {
        this.getContacts();
        this.getProfiles();
    }

    public getContacts() {
        WrestApi.contacts().then((data) => {
            data.forEach((item) => this.contacts[item.wxid] = item);
        });
    }

    public getProfiles() {
        RobotApi.profileList({} as ProfileFetchAllParam).then((data) => {
            this.profiles = data;
        });
    }

}
