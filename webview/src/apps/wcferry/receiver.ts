import { Component, OnDestroy } from '@angular/core';


@Component({
    selector: 'page-wcferry-receiver',
    templateUrl: 'receiver.html'
})
export class WcferryReceiverComponent implements OnDestroy {

    public wss!: WebSocket;
    public messages: Array<string> = [];

    constructor() {
        this.startSocket();
    }

    public ngOnDestroy() {
        this.stopSocket();
    }

    public stopSocket() {
        this.wss && this.wss.close();
        this.messages = [];
    }

    public startSocket() {
        const token = sessionStorage.getItem('token');
        const url = location.origin.replace(/^http/, 'ws') + '/wcf/socket_receiver';
        const wss = new WebSocket(url + (token ? '?token=' + token : ''));
        wss.onopen = () => {
            this.messages.push('websocket is connected');
            this.wss = wss;
        };
        wss.onclose = () => {
            this.messages.push('websocket is closed, retry in 5s');
            setTimeout(() => this.startSocket(), 5 * 1000);
        };
        wss.onerror = (event) => {
            this.messages.push('websocket error, details to console');
            console.log(event);
        };
        wss.onmessage = (event) => {
            this.messages.push(event.data);
        };
    }

}
