export const environment = {
    production: false,
    apiUrl: 'http://localhost:8080/api',
    useMsalAuth: false,
    msalConfig: {
        clientId: 'DEV CLIENTID FROM APP REGISTRATION HERE',
        authority: 'https://login.microsoftonline.com/b2748d0a-856e-4184-bda8-831f9ffa8a48',
        redirectUri: 'http://localhost:4200/auth',
        apiScope: 'OUR SCOPE (api.access for sure) NEEDS TO BE HERE',
        graphUrl: 'https://graph.microsoft.com/v1.0/me',
        graphScope: 'User.Read'
    }
};