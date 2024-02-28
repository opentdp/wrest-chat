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
    public roomMembers: Record<string, Array<WcfrestContactPayload>> = {};

    public formdata: ProfileCreateParam = {};

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
        this.formdata.wxid = '';
        const id = this.formdata.roomid || '-';
        this.conacts = id == '-' ? this.friends : this.roomMembers[id] || [];
    }

    public getFriends() {
        WrestApi.friends().then((data) => {
            this.friends = data || [];
        });
    }

    public getChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.chatrooms = data || [];
            this.getRoomMembers(this.chatrooms.map((item) => item.wxid));
        });
    }

    public getRoomMembers(ids: string[]) {
        [...new Set(ids)].forEach((id) => {
            WrestApi.chatroomMembers({ roomid: id }).then((data) => {
                this.roomMembers[id] = data;
            });
        });
    }

}
