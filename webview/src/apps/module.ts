import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { BasicModule } from '../helpers/basic';

import { AppLayouts } from '../layouts';

import { RootComponent } from './root';
import { AppComponents, AppRoutes } from './pages';


@NgModule({
    imports: [
        BasicModule,
        RouterModule.forRoot(AppRoutes, {
            scrollPositionRestoration: 'top'
        }),
    ],
    declarations: [
        AppLayouts,
        RootComponent,
        AppComponents,
    ],
    bootstrap: [RootComponent]
})
export class AppModule { }
