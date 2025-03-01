import { EnvironmentProviders, makeEnvironmentProviders } from '@angular/core';
import { Configuration } from './configuration';

export function provideApiModule(basePath: string): EnvironmentProviders {
    return makeEnvironmentProviders([
        { provide: Configuration, useValue: new Configuration({ basePath }) }
    ]);
}
