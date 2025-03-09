import { HttpBackend } from "@angular/common/http";

export const environment = {
    production: false,
    API_BASE_PATH: 'http://localhost:8080/api',

};

// Is called by the app initializer during the bootstrap process
export async function getConfig() {
    const response = await fetch('/config');
    const config = await response.json();
    if (config?.API_BASE_PATH) {
        environment.API_BASE_PATH = config.API_BASE_PATH;
    }
}
        