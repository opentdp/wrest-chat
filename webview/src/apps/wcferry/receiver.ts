import { Component, OnDestroy } from '@angular/core';


@Component({
    selector: 'page-wcferry-receiver',
    templateUrl: 'receiver.html'
})
export class WcferryReceiverComponent implements OnDestroy {

    public wst!: WebSocket;
    public messages: Array<string> = [];

    constructor() {
        this.startSocket();
    }

    ngOnDestroy(): void {
        this.wst && this.wst.close();
        this.messages = [];
    }

    public async startSocket() {
        const url = location.origin.replace(/^http/, 'ws');
        const wst = new WebSocket(url + '/wcf/socket_receiver');
        wst.onopen = () => {
            this.messages.push('WebSocket is connected.');
            this.wst = wst;
        };
        wst.onclose = () => {
            this.messages.push('WebSocket is closed now.');
        };
        wst.onerror = (error) => {
            this.messages.push('WebSocket Error:' + error);
        };
        wst.onmessage = (event) => {
            this.messages.push(event.data);
        };
    }

}
