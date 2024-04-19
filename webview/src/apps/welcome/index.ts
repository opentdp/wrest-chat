import { Component } from '@angular/core';

import { SundryApi, SystemVersion } from '../../openapi/sundry';
import { WrestApi, WcfrestCommonPayload, WcfrestUserInfoPayload, WcfrestAvatarPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-welcome',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class WelcomeComponent {

    public islogin = false;
    public loginqr!: WcfrestCommonPayload;

    public user!: WcfrestUserInfoPayload;
    public avatar!: WcfrestAvatarPayload;

    public system!: SystemVersion;
    public upgrade = {
        url: '', version: ''
    };

    public formdata = {
        token: '',
    };

    constructor() {
        this.checkLogin();
        this.refreshQrcode();
        this.getSystemVersion();
        // 获取会话存储的令牌
        this.formdata.token = sessionStorage.getItem('token') || '';
    }

    public submitForm() {
        sessionStorage.setItem('token', this.formdata.token);
        location.reload();
    }

    public async checkLogin() {
        this.islogin = await WrestApi.isLogin();
        if (this.islogin) {
            return this.getSelfInfo();
        }
        setTimeout(() => this.checkLogin(), 5 * 1000);
    }

    public async refreshQrcode() {
        this.loginqr = await WrestApi.loginQr();
        if (!this.islogin) {
            const t = this.loginqr.result ? 200 : 5;
            setTimeout(() => this.refreshQrcode(), t * 1000);
        }
    }

    public async getSelfInfo() {
        this.user = await WrestApi.selfInfo();
        if (this.user && this.user.wxid) {
            const avatars = await WrestApi.avatars({ wxids: [this.user.wxid] });
            if (avatars && avatars.length > 0) {
                this.avatar = avatars[0];
            }
        }
    }

    public async getSystemVersion() {
        await SundryApi.systemVersion({}).then((data) => {
            this.system = data;
        });
        // 新版检查
        const ghApi = 'https://api.github.com/repos/opentdp/wrest-chat/releases/latest';
        fetch(ghApi).then(r => r.json()).then(data => {
            if (!data || !data.created_at) {
                return;
            }
            if (this.compareVersions(this.system.version, data.tag_name) < 1) {
                this.upgrade.url = data.assets[1].browser_download_url;
                this.upgrade.version = data.tag_name;
            }
        });

    }

    public compareVersions(v1: string, v2: string) {
        const v1s = v1.replace(/^[a-zA-Z]+/, '').split('.');
        const v2s = v2.replace(/^[a-zA-Z]+/, '').split('.');
        const len = Math.min(v1s.length, v2s.length);
        // 逐级比较
        for (let i = 0; i < len; i++) {
            const n1 = parseInt(v1s[i], 10);
            const n2 = parseInt(v2s[i], 10);
            if (n1 > n2) {
                return 1;
            }
            if (n1 < n2) {
                return -1;
            }
        }
        // 版本号长度不同
        if (v1s.length > v2s.length) {
            return 1;
        }
        if (v1s.length < v2s.length) {
            return -1;
        }
        return 0;
    }

}