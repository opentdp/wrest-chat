import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { UserLevels } from 'src/openapi/const';
import { RobotApi, ProfileCreateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-profile-create',
    templateUrl: 'create.html'
})
export class ProfileCreateComponent {

    public userLevels = UserLevels;

    public wcfAvatars: Record<string, string> = {};
    public wcfFriends: Array<WcfrestContactPayload> = [];
    public wcfContacts: Record<string, WcfrestContactPayload> = {};
    public wcfChatrooms: Array<WcfrestContactPayload> = [];
    public wcfRoomMembers: Record<string, Array<WcfrestContactPayload>> = {};

    public conacts: Array<WcfrestContactPayload> = [];
    public formdata: ProfileCreateParam = {
        wxid: '',
        roomid: '',
        level: 2,
    };

    constructor(private router: Router) {
        this.getWcfFriends();
        this.getWcfChatrooms();
    }

    public createProfile() {
        if (this.formdata.level) {
            this.formdata.level = +this.formdata.level;
        }
        RobotApi.profileCreate(this.formdata).then(() => {
            this.router.navigate(['profile/list']);
        });
    }

    public changeConacts() {
        const id = this.formdata.roomid || '-';
        this.conacts = id == '-' ? this.wcfFriends : this.wcfRoomMembers[id] || [];
    }

    public getWcfFriends() {
        WrestApi.friends().then((data) => {
            this.wcfFriends = data || [];
        });
    }

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.wcfChatrooms = data || [];
            this.getWcfRoomMembers(this.wcfChatrooms.map((item) => item.wxid));
        });
    }

    public getWcfRoomMembers(ids: string[]) {
        [...new Set(ids)].forEach((id) => {
            WrestApi.chatroomMembers({ roomid: id }).then((data) => {
                this.wcfRoomMembers[id] = data || [];
            });
        });
    }

}
