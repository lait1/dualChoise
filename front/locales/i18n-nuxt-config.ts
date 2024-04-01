export const i18n = {
  strategy: 'no_prefix',
  defaultLocale: 'en',
  langDir: 'locales',
  locales: [
    {
      code: 'en',
      name: 'English',
      iso: 'en-US',
      file: 'api.ts'
    },
    {
      code: 'es',
      name: 'Español',
      iso: 'es-ES',
      file: 'api.ts'
    },
    {
      code: 'ru',
      name: 'Русский',
      iso: 'ru-RU',
      file: 'api.ts'
    },
  ],
  detectBrowserLanguage: {
    useCookie: true,
    cookieKey: 'i18n_redirected',
    redirectOn: 'root'
  }

}