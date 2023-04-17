import { createApp } from "vue";
import App from "~/App.vue";
import Error from "~/views/Error.vue";
import { i18n, i18nError, loadLocaleAsync } from "~/libs/i18n";
import { router } from "~/libs/router";
import { auth } from "~/services/auth.service";
import "~/styles/index.scss";
import { useUser } from "./composables/user";

const setup = async () => {
  const { setUser } = useUser();

  try {
    const app = createApp(App);

    const [_, user] = await Promise.all([
      loadLocaleAsync(),
      auth.getUserInfo(),
    ]);

    if (user) {
      setUser(user);
    }

    app.use(router);
    app.use(i18n);
    app.mount("#app");
  } catch (e) {
    console.log(e);
    const error = createApp(Error);
    error.use(i18nError);
    error.mount("#app");
  }
};

setup();
