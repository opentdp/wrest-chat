import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import { UserLevels } from 'src/openapi/const';
import { RobotApi, ProfileUpdateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-profile-update',
    templateUrl: 'update.html'
})
export class ProfileUpdateComponent implements OnInit {

    public userLevels = UserLevels;

    public wcfAvatars: Record<string, string> = {};
    public wcfFriends: Array<WcfrestContactPayload> = [];
    public wcfContacts: Record<string, WcfrestContactPayload> = {};
    public wcfChatrooms: Array<WcfrestContactPayload> = [];
    public wcfRoomMembers: Record<string, Array<WcfrestContactPayload>> = {};

    public conacts: Array<WcfrestContactPayload> = [];
    public formdata = {} as ProfileUpdateParam;

    constructor(
        private router: Router,
        private route: ActivatedRoute
    ) {
        this.getWcfFriends();
        this.getWcfChatrooms();
    }

    public ngOnInit() {
        const rd = this.route.snapshot.paramMap.get('rd');
        rd && this.getProfile(+rd);
    }

    public getProfile(rd: number) {
        RobotApi.profileDetail({ rd }).then((data) => {
            data && Object.assign(this.formdata, data);
        });
    }

    public updateProfile() {
        if (this.formdata.level) {
            this.formdata.level = +this.formdata.level;
        }
        RobotApi.profileUpdate(this.formdata).then(() => {
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
                // 尝试更新当前人员列表
                this.changeConacts();
            });
        });
    }

}
