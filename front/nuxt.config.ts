// https://nuxt.com/docs/api/configuration/nuxt-config
import svgLoader from 'vite-svg-loader'
import { i18n } from './locales/i18n-nuxt-config'
// @ts-ignore
export default defineNuxtConfig({
  devtools: { enabled: true },
  css: [
    '~/assets/scss/main.scss',
  ],
  vite: {
    define: {
      'process.env.DEBUG': false
    },
    css: {
      devSourcemap: false,
      preprocessorOptions: {
        scss: {
          additionalData: '@use "~/assets/scss/_colors.scss" as *;'
        }
      }
    },
    plugins: [
      svgLoader()
    ]
  },
  app: {
    head: {
      charset: 'utf-8',
      viewport: 'width=device-width, initial-scale=1',
    }
  },
  modules: [
    ['@nuxtjs/i18n', i18n],
    ['@nuxtjs/device', { refreshOnResize: true }],
  ],
  runtimeConfig: {
    public: {
      baseUrl: process.env.NUXT_MAIN_DOMAIN,
      serverUrl: process.env.NUXT_SERVER_DOMAIN || 'http://app:4000',
    }
  },
})
