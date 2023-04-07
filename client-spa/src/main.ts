import { createApp } from "vue";
import App from "~/App.vue";
import Error from "~/views/Error.vue";
import { i18n, i18nError, loadLocaleAsync } from "~/libs/i18n";
import { router } from "~/libs/router";
import { auth } from "~/services/auth.service";
import "~/styles/tailwind.css";

const setup = async () => {
  const app = createApp(App);
  app.use(router);
  app.use(i18n);

  try {
    await Promise.all([loadLocaleAsync(), auth.getUserInfo()]);
    app.mount("#app");
  } catch (e) {
    console.log(e);
    const error = createApp(Error);
    error.use(i18nError);
    error.mount("#app");
  }
};

setup();
