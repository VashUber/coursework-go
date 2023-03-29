export default defineNuxtConfig({
  modules: ["@nuxtjs/i18n"],
  i18n: {
    langDir: "locales",
    locales: [
      {
        code: "en",
        iso: "en-US",
        file: "en.json",
      },
      {
        code: "ru",
        iso: "ru",
        file: "ru.json",
      },
    ],
    lazy: true,
    defaultLocale: "ru",
    detectBrowserLanguage: {
      useCookie: true,
      cookieKey: "i18n_lang",
      redirectOn: "root",
      fallbackLocale: "en",
    },
    vueI18n: {
      fallbackLocale: "en",
      legacy: false,
    },
  },
  css: ["@/styles/global.scss"],
  vite: {
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: "@import '@/styles/vars.scss';",
        },
      },
    },
  },
});
