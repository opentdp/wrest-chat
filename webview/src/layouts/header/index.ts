import { Component, Input } from '@angular/core';


@Component({
    selector: 'layout-header',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class LayoutHeaderComponent {

    public collapse = false;

    @Input()
    public set title(val: string) {
        document.title = (val ? val + ' - ' : '') + 'Wrest';
    }

}
