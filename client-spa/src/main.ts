import { createApp } from "vue";
import App from "~/App.vue";
import { i18n, loadLocaleAsync } from "~/libs/i18n";
import { auth } from "~/services/auth.service";
import "~/styles/tailwind.css";

const setup = async () => {
  const app = createApp(App);

  await Promise.all([loadLocaleAsync(), auth.getUserInfo()]);

  app.use(i18n);
  app.mount("#app");
};

setup();
