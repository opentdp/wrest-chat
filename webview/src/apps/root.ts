import { Component } from '@angular/core';


@Component({
    selector: 'app-root',
    template: `
        <layout-toast></layout-toast>
        <router-outlet></router-outlet>
    `
})
export class RootComponent {

    public progress = 0;

}
