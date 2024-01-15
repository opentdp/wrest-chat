import { enableProdMode } from '@angular/core';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

import { Package } from './helpers/env.dev';
import { AppModule } from './apps/module';


Package.production && enableProdMode();

document.addEventListener('DOMContentLoaded', () => {
    platformBrowserDynamic().bootstrapModule(AppModule).catch(err => console.error(err));
});
