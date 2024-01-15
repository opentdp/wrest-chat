import { ChatsComponent } from './chats';
import { ContactsComponent } from './contacts';

import { ErrorComponent } from './error';


export const AppComponents = [
    ChatsComponent,
    ContactsComponent,
    ErrorComponent,
];

//////////////////////////////////////////////////////////////////

import { Routes } from '@angular/router';

export const AppRoutes: Routes = [
    { path: '', redirectTo: 'chats', pathMatch: 'full' },
    { path: 'chats', component: ChatsComponent },
    { path: 'contacts', component: ContactsComponent },
    { path: '**', component: ErrorComponent, data: { error: 404 } }
];
