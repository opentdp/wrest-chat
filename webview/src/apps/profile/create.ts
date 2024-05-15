import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { UserLevels } from '../../openapi/const';
import { RobotApi, TablesLLModel, ProfileCreateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-profile-create',
    templateUrl: 'create.html'
})
export class ProfileCreateComponent {

    public userLevels = UserLevels;

    public llmodels: Array<TablesLLModel> = [];

    public wcfAvatars: Record<string, string> = {};
    public wcfFriends: Array<WcfrestContactPayload> = [];
    public wcfChatrooms: Array<WcfrestContactPayload> = [];
    public wcfRoomMembers: Record<string, Array<WcfrestContactPayload>> = {};

    public conacts: Array<WcfrestContactPayload> = [];
    public conactsFilter = '';

    public formdata: ProfileCreateParam = {
        wxid: '',
        roomid: '',
        level: -1,
    };

    constructor(private router: Router) {
        this.getLLModels();
        this.getWcfFriends();
        this.getWcfChatrooms();
    }

    public createProfile() {
        if (this.formdata.level) {
            this.formdata.level = +this.formdata.level;
        }
        return RobotApi.profileCreate(this.formdata).then(() => {
            this.router.navigate(['profile/list']);
        });
    }

    public changeConacts() {
        const id = this.formdata.roomid || '-';
        return this.getWcfRoomMembers(this.formdata.roomid).then(() => {
            this.conacts = id == '-' ? this.wcfFriends : this.wcfRoomMembers[id] || [];
        });
    }

    public getLLModels() {
        return RobotApi.llmodelList({}).then((data) => {
            this.llmodels = data || [];
        });
    }

    public getWcfFriends() {
        return WrestApi.friends().then((data) => {
            this.wcfFriends = data || [];
        });
    }

    public getWcfChatrooms() {
        return WrestApi.chatrooms().then((data) => {
            this.wcfChatrooms = data || [];
        });
    }

    public getWcfRoomMembers(id: string) {
        if (this.wcfRoomMembers[id]) {
            return Promise.resolve(); //已获取
        }
        this.wcfRoomMembers[id] = []; //初始化
        return WrestApi.chatroomMembers({ roomid: id }).then((data) => {
            this.wcfRoomMembers[id] = data || [];
        });
    }

}
