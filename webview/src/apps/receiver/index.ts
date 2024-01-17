import { Component } from '@angular/core';


@Component({
    selector: 'page-receiver',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class ReceiverComponent {

    public messages: Array<string> = [];

    constructor() {
        this.startSocket();
    }

    public async startSocket() {
        const url = location.origin.replace(/^http/, 'ws');
        const websocket = new WebSocket(url + '/api/socket_receiver');
        websocket.onopen = () => {
            this.messages.push('WebSocket is connected.');
        };
        websocket.onmessage = event => {
            this.messages.push(event.data);
        };
        websocket.onerror = (error) => {
            this.messages.push('WebSocket Error:' + error);
        };
        websocket.onclose = () => {
            this.messages.push('WebSocket is closed now.');
        };
    }

}
