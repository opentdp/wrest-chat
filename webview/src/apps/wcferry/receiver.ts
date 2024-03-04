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

    ngOnDestroy(): void {
        this.wss && this.wss.close();
        this.messages = [];
    }

    public async startSocket() {
        const url = location.origin.replace(/^http/, 'ws');
        const wss = new WebSocket(url + '/wcf/socket_receiver');
        wss.onopen = () => {
            this.messages.push('websocket is connected');
            this.wss = wss;
        };
        wss.onclose = () => {
            this.messages.push('websocket is closed');
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
