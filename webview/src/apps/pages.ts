import { ChatroomsComponent } from './chatrooms';
import { ContactsComponent } from './contacts';

import { ErrorComponent } from './error';


export const AppComponents = [
    ChatroomsComponent,
    ContactsComponent,
    ErrorComponent,
];

//////////////////////////////////////////////////////////////////

import { Routes } from '@angular/router';

export const AppRoutes: Routes = [
    { path: '', redirectTo: 'chatrooms', pathMatch: 'full' },
    { path: 'chatrooms', component: ChatroomsComponent },
    { path: 'contacts', component: ContactsComponent },
    { path: '**', component: ErrorComponent, data: { error: 404 } }
];
