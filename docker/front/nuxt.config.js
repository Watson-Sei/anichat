const enviroment = process.env.NODE_ENV
require('dotenv').config({path: `config/.env.${enviroment}`})
const {
  API_URL,
  API_URL_BROWSER,
  BASE_URL,
  apiKey,
  databaseURL,
  projectId,
  messagingSenderId,
  appId,
  measurementId,
} = process.env
export default {
  // Global page headers (https://go.nuxtjs.dev/config-head)
  head: {
    title: 'front',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS (https://go.nuxtjs.dev/config-css)
  css: [
  ],

  // Plugins to run before rendering page (https://go.nuxtjs.dev/config-plugins)
  plugins: [
    '~/plugins/firebase.js'
  ],

  // Auto import components (https://go.nuxtjs.dev/config-components)
  components: true,

  // Modules for dev and build (recommended) (https://go.nuxtjs.dev/config-modules)
  buildModules: [
    // https://go.nuxtjs.dev/tailwindcss
    '@nuxtjs/tailwindcss',
    '@nuxtjs/axios',
  ],

  // Modules (https://go.nuxtjs.dev/config-modules)
  modules: [
  ],

  // Build Configuration (https://go.nuxtjs.dev/config-build)
  build: {
  },
  env: {
    API_URL,
    API_URL_BROWSER,
    BASE_URL,
    apiKey,
    databaseURL,
    projectId,
    messagingSenderId,
    appId,
    measurementId
  }
}
