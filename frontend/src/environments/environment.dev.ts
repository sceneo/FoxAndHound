export const environment = {
    production: false,
    apiUrl: 'OUR DEV API URL HERE',
    useMsalAuth: true,
    msalConfig: {
        clientId: 'CLIENTID FROM APP REGISTRATION HERE',
        authority: 'https://login.microsoftonline.com/b2748d0a-856e-4184-bda8-831f9ffa8a48',
        redirectUri: 'OUR_BASE_URI/auth',
        apiScope: 'OUR SCOPE (api.access for sure) NEEDS TO BE HERE',
        graphUrl: 'https://graph.microsoft.com/v1.0/me',
        graphScope: 'User.Read'
    }
};