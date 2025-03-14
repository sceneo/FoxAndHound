import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { AppComponent } from './app/app.component';

export type RuntimeAppConfig = {
  apiBaseUrl: string;
};

// Default values, a better way would be to use an injection token instead of exporting the value
// but I couldn't get it to work
const runtimeAppConfig: RuntimeAppConfig = {
  apiBaseUrl: 'http://localhost:8080/api',
};

export const getRuntimeConfig = () => runtimeAppConfig;

// Fetch the config from the server and bootstrap the application
fetchApiBasePath().then((apiBaseUrl) => {
  if (apiBaseUrl) {
    runtimeAppConfig.apiBaseUrl = apiBaseUrl;
    console.log('base path set, runtimeAppConfig', runtimeAppConfig);
  }
  bootstrapApplication(AppComponent, appConfig).catch((err) => console.error(err));
});

async function fetchApiBasePath(): Promise<string | undefined> {
  console.log(`getConfig from base path ${window.location.origin}`);
  const response = await fetch('/config');
  console.log('response', response);
  let config;
  try {
    config = await response.json();
  } catch (e) {
    // When running on localhost there is no server providing the config
    console.error('Failed to fetch config: this is expected on localhost without nginx', e);
    return undefined;
  }
  console.log('config from server', config);
  if (config?.apiBaseUrl) {
    console.log('apiBaseUrl', config.apiBaseUrl);
    return config.apiBaseUrl;
  }
  return undefined;
}

// Return Promise that resolves after 5 seconds to localhost:8085/api
// Simulate nginx response on localhost, use this instead of fetchApiBasePath
async function fetchApiBasePathDummy(): Promise<string | undefined> {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve('http://localhost:8085/api');
    }, 5000);
  });
}