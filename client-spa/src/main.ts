import { createApp } from "vue";
import App from "./App.vue";
import { i18n, loadLocaleAsync } from "~/libs/i18n";
import "~/libs/axios";
import { auth } from "./services/auth.service";

const setup = async () => {
  const app = createApp(App);

  await Promise.all([loadLocaleAsync(), auth.signin()]);

  app.use(i18n);
  app.mount("#app");
};

setup();
