import { LoginGuard } from '../helpers/login-guard';

import { WelcomeComponent } from './welcome';

import { ChatroomsComponent } from './chatrooms';
import { ContactsComponent } from './contacts';
import { ReceiverComponent } from './receiver';

import { ErrorComponent } from './error';


export const AppComponents = [
    WelcomeComponent,

    ChatroomsComponent,
    ContactsComponent,
    ReceiverComponent,

    ErrorComponent,
];

//////////////////////////////////////////////////////////////////

import { Routes } from '@angular/router';

export const AppRoutes: Routes = [
    { path: '', redirectTo: 'welcome', pathMatch: 'full' },
    { path: 'welcome', component: WelcomeComponent },

    { path: 'chatrooms', component: ChatroomsComponent, canActivate: [LoginGuard] },
    { path: 'contacts', component: ContactsComponent, canActivate: [LoginGuard] },
    { path: 'receiver', component: ReceiverComponent, canActivate: [LoginGuard] },

    { path: '**', component: ErrorComponent, data: { error: 404 } }
];
