import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { RobotApi, ProfileCreateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-profile-create',
    templateUrl: 'create.html',
    styleUrls: ['create.scss']
})
export class ProfileCreateComponent {

    public conacts: Array<WcfrestContactPayload> = [];
    public friends: Array<WcfrestContactPayload> = [];
    public chatrooms: Array<WcfrestContactPayload> = [];

    public formdata = {} as ProfileCreateParam;

    constructor(private router: Router) {
        this.getChatrooms();
        this.getFriends();
    }

    public createProfile() {
        RobotApi.profileCreate(this.formdata).then(() => {
            this.router.navigate(['profile/list']);
        });
    }

    public changeRoomid() {
        if (this.formdata.roomid == '-') {
            this.conacts = this.friends;
        } else {
            this.getRoomMembers(this.formdata.roomid);
        }
    }

    public getFriends() {
        WrestApi.friends().then((data) => {
            this.friends = data || [];
        });
    }

    public getChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.chatrooms = data || [];
        });
    }

    public getRoomMembers(id: string) {
        WrestApi.chatroomMembers({ roomid: id }).then((data) => {
            this.conacts = data || [];
        });
    }

}
