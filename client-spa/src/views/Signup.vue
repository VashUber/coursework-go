<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import dayjs from "dayjs";
import Input from "~/components/ui/Input.vue";
import AuthLayout from "./components/auth/AuthLayout.vue";
import { authService } from "~/services/auth.service";

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
</script>

<template>
  <AuthLayout @submit="onSubmit">
    <template #form>
      <Input v-model="formData.email"> Email </Input>
      <Input v-model="formData.name"> Name </Input>
      <Input v-model="formData.birthday" v-maska data-maska="##.##.####">Birthday</Input>
      <Input v-model="formData.password" type="password">Password</Input>

      <button type="submit" class="button">
        {{ $t("page.signup.btn") }}
      </button>

      <router-link to="/signin" class="text-center">
        {{ $t("page.signup.go-to-signin") }}
        {{ $t("misc.signin") }}
      </router-link>
    </template>

    <template #title>
      {{ $t("misc.signup") }}
    </template>
  </AuthLayout>
</template>
