<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import dayjs from "dayjs";
import Input from "~/components/ui/Input.vue";
import AuthLayout from "./components/auth/AuthLayout.vue";
import { authService } from "~/services/auth.service";
import { useMeta } from "~/composables/meta";
import { useI18n } from "vue-i18n";

const router = useRouter();

const onSubmit = async () => {
  const [d, m, y] = formData.value.birthday.split(".");

  await authService.signup({
    email: formData.value.email,
    password: formData.value.password,
    name: formData.value.name,
    birthday: dayjs([m, d, y].join("-")).format(),
  });

  router.push({
    name: "signin",
  });
};

const formData = ref({
  email: "",
  password: "",
  name: "",
  birthday: "",
});

const { setHead } = useMeta();
const { t } = useI18n();
setHead({
  title: t("page.signup.title"),
  description: t("page.index.description"),
});
</script>

<template>
  <AuthLayout @submit="onSubmit">
    <template #form>
      <Input v-model="formData.email"> {{ $t("misc.email") }} </Input>
      <Input v-model="formData.name">
        {{ $t("misc.name") }}
      </Input>
      <Input v-model="formData.birthday" v-maska data-maska="##.##.####">
        {{ $t("misc.birthday") }}
      </Input>
      <Input v-model="formData.password" type="password">
        {{ $t("misc.password") }}
      </Input>

      <button type="submit" class="button">
        {{ $t("page.signup.btn") }}
      </button>

      <router-link to="/signin" class="text-center">
        {{ $t("page.signup.go-to-signin") }}
        {{ $t("page.signin.btn") }}
      </router-link>
    </template>

    <template #title>
      {{ $t("misc.signup") }}
    </template>
  </AuthLayout>
</template>
