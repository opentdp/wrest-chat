import { LoginGuard } from '../helpers/login-guard';

import { WelcomeComponent } from './welcome';
import { ErrorComponent } from './error';

import { ChatroomsComponent } from './wcferry/chatrooms';
import { ContactsComponent } from './wcferry/contacts';
import { ReceiverComponent } from './wcferry/receiver';


export const AppComponents = [
    WelcomeComponent,
    ErrorComponent,

    ChatroomsComponent,
    ContactsComponent,
    ReceiverComponent,
];

//////////////////////////////////////////////////////////////////

import { Routes } from '@angular/router';

export const AppRoutes: Routes = [
    { path: 'welcome', component: WelcomeComponent },

    { path: 'wcf/chatrooms', component: ChatroomsComponent, canActivate: [LoginGuard] },
    { path: 'wcf/contacts', component: ContactsComponent, canActivate: [LoginGuard] },
    { path: 'wcf/receiver', component: ReceiverComponent, canActivate: [LoginGuard] },

    { path: '', redirectTo: 'welcome', pathMatch: 'full' },
    { path: '**', component: ErrorComponent, data: { error: 404 } },
];
