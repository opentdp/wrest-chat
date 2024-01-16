import { Component } from '@angular/core';

import { WrestApi, WcfrestUserInfoPayload, WcfrestAvatarPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-owner',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class OwnerComponent {

    public user!: WcfrestUserInfoPayload;
    public avatar!: WcfrestAvatarPayload;

    constructor() {
        this.getSelfInfo();
    }

    public async getSelfInfo() {
        await WrestApi.selfInfo().then((user) => {
            this.user = user;
        });
        WrestApi.avatars({ wxids: [this.user.wxid] }).then((avatars) => {
            if (avatars && avatars.length > 0) {
                this.avatar = avatars[0];
            }
        });
    }

}
