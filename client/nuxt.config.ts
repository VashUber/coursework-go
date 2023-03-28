export default defineNuxtConfig({
  modules: ["@nuxtjs/i18n"],
  i18n: {
    langDir: "locales",
    locales: [
      {
        code: "en",
        file: "en.json",
      },
      {
        code: "ru",
        file: "ru.json",
      },
    ],
    defaultLocale: "ru",
    detectBrowserLanguage: {
      useCookie: true,
      cookieKey: "i18n_lang",
      redirectOn: "root",
      fallbackLocale: "en",
      alwaysRedirect: true,
    },
    vueI18n: {
      fallbackLocale: "en",
      legacy: false,
    },
  },
});
