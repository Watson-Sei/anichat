const enviroment = process.env.NODE_ENV
console.log(enviroment)
require('dotenv').config({path: `config/.env.${enviroment}`})
const { API_URL, BASE_URL, WebSocket_URL, apiKey } = process.env;

export default {
  // Global page headers (https://go.nuxtjs.dev/config-head)
  head: {
    title: 'AniChat - チャットアプリ ',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'AniChatは、地上波放送などのアニメの視聴感想をリアルタイムで発信できるチャットサービスです。' },
      { hid: 'keywords', name: 'keywords', content: 'アニメチャット,AniChat,アニメ,チャット'},
      { hid: 'og:site_name', property: 'og:site_name', content: 'AniChat'},
      { hid: 'og:type', property: 'og:type', content: 'website'},
      { hid: 'og:url', property: 'og:url', content: 'https://www.anichat.jp'},
      { hid: 'og:title', property: 'og:title', content: 'AniChat'},
      { hid: 'og:description', property: 'og:description', content: 'AniChatは、地上波放送などのアニメの視聴感想をリアルタイムで発信できるチャットサービスです。'},
      { hid: 'og:image', property: 'og:image', content: 'https://media.discordapp.net/attachments/616249287030079488/825360800763478016/anichat.png'}
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'apple-touch-icon', type: 'image/x-icon', href: './apple-touch-icon.png'},
    ]
  },

  // Global CSS (https://go.nuxtjs.dev/config-css)
  css: [
  ],

  // Plugins to run before rendering page (https://go.nuxtjs.dev/config-plugins)
  plugins: [
    '~/plugins/axios.js'
  ],

  // Auto import components (https://go.nuxtjs.dev/config-components)
  components: true,

  // Modules for dev and build (recommended) (https://go.nuxtjs.dev/config-modules)
  buildModules: [
    // https://go.nuxtjs.dev/tailwindcss
    '@nuxtjs/axios',
    '@nuxtjs/vuetify',
  ],

  // Modules (https://go.nuxtjs.dev/config-modules)
  modules: [
    '@nuxtjs/dotenv',
    '@nuxtjs/pwa',
    'nuxt-fontawesome',
    [
      '@nuxtjs/firebase',
      {
        config: {
          apiKey: process.env.apiKey,
          authDomain: process.env.authDomain,
          databaseURL: process.env.databaseURL,
          projectId: process.env.projectId,
          storageBucket: process.env.storageBucket,
          messagingSenderId: process.env.messagingSenderId,
          appId: process.env.appId,
          measurementId: process.env.measurementId
        },
        services: {
          auth: {
            persistence: 'local',
            initialize: {
              onAuthStateChangedAction: 'onAuthStateChanged',
              onIdTokenChanged: 'onIdTokenChanged',
            },
            ssr: true
          }
        }
      }
    ]
  ],

  pwa: {
    meta: false,
    icon: false,

    workbox: {
      importScripts: [
        '/firebase-auth-sw.js'
      ],

      dev: true
    }
  },

  fontawesome: {
    imports: [
      {
        set: '@fortawesome/free-solid-svg-icons',
        icons: ['fas']
      }
    ]
  },

  // Build Configuration (https://go.nuxtjs.dev/config-build)
  build: {
  },

  axios: {
  },
  env: {
    API_URL,
    BASE_URL,
    WebSocket_URL,
    apiKey,
  }
}
