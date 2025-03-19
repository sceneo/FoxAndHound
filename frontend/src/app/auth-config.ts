import {
    IPublicClientApplication,
    LogLevel,
    PublicClientApplication,
    InteractionType,
    BrowserCacheLocation
} from '@azure/msal-browser';
import { MsalGuardConfiguration, MsalInterceptorConfiguration } from '@azure/msal-angular';
import { environment } from '../environments/environment';

export function MSALInstanceFactory(): IPublicClientApplication {
    return new PublicClientApplication({
      auth: {
        clientId: environment.msalConfig.clientId,
        authority: environment.msalConfig.authority,
        redirectUri: environment.msalConfig.redirectUri
      },
      cache: {
        cacheLocation: BrowserCacheLocation.LocalStorage,
        storeAuthStateInCookie: !environment.production
      },
      system: {
        loggerOptions: {
          logLevel: LogLevel.Info,
          loggerCallback: (_, message, __) => {
            console.log(message);
          },
        },
      },
    });
}

export function MSALInterceptorConfigFactory(): MsalInterceptorConfiguration {
    const protectedResourceMap = new Map<string, Array<string>>();
  
    protectedResourceMap.set(environment.msalConfig.graphUrl, [environment.msalConfig.graphScope]);
    protectedResourceMap.set(environment.apiUrl, [environment.msalConfig.apiScope]);
    
    return {
      interactionType: InteractionType.Redirect,
      protectedResourceMap
    };
}

export function MSALGuardConfigFactory(): MsalGuardConfiguration {
    return {
        interactionType: InteractionType.Redirect,
        authRequest: {
            scopes: [
                environment.msalConfig.graphScope,
                environment.msalConfig.apiScope
            ]
        }
    };
}
