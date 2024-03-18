import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { BasicModule } from '../helpers/basic';
import { Pipes } from '../helpers/pipes';

import { AppLayouts } from '../layouts';

import { RootComponent } from './root';
import { AppComponents, AppRoutes } from './pages';


@NgModule({
    imports: [
        BasicModule,
        RouterModule.forRoot(AppRoutes, {
            scrollPositionRestoration: 'top',
            useHash: true
        }),
    ],
    declarations: [
        Pipes,
        AppLayouts,
        RootComponent,
        AppComponents,
    ],
    bootstrap: [RootComponent]
})
export class AppModule { }
