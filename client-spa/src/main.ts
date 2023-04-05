import { createApp } from "vue";
import App from "./App.vue";
import { i18n, loadLocaleAsync } from "~/libs/i18n";
import "~/libs/axios";

const setup = async () => {
  const app = createApp(App);

  await loadLocaleAsync();

  app.use(i18n);
  app.mount("#app");
};

setup();
