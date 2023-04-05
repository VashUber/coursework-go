import { createI18n } from "vue-i18n";
import cookies from "js-cookie";
import en from "~/locales/en.json";

const langCookie = cookies.get("app_lang");

export const i18n = createI18n({
  locale: "en",
  fallbackLocale: "en",
  legacy: false,
  messages: {
    en,
  },
});

export const loadLocaleAsync = async () => {
  const lang = langCookie || "en";

  if (lang === "en") return;

  const locale = await import(`../locales/${lang}.json`);

  i18n.global.setLocaleMessage(lang, locale);
  (i18n.global.locale.value as string) = lang;

  document.querySelector("html")!.setAttribute("lang", lang);
};
