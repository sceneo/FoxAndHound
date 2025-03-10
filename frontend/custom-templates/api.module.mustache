import { EnvironmentProviders, makeEnvironmentProviders } from '@angular/core';
import { Configuration } from './configuration';

export function provideApiModule(basePath: string): EnvironmentProviders {
    return makeEnvironmentProviders([
        { provide: Configuration, useFactory: () => getConfig(basePath) }
    ]);
}

// Getting the config on runtime from the server which creates it from env variables.
export async function getConfig(basePath: string): Promise<Configuration> {
    console.log('getConfig');
    const response = await fetch('/config');
    const config = await response.json();
    console.log('config from server', config);
    if (config?.API_BASE_PATH) {
        console.log('API_BASE_PATH', config.API_BASE_PATH);
        basePath = config.API_BASE_PATH;
    }
    return new Configuration({ basePath });
}