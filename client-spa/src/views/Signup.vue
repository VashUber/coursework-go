<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import Input from "~/components/ui/Input.vue";
import Button from "~/components/ui/Button.vue";
import AuthLayout from "./components/auth/AuthLayout.vue";
import { auth } from "~/services/auth.service";

const router = useRouter();

const onSubmit = async () => {
  await auth.signup({
    email: formData.value.email,
    password: formData.value.password,
    name: formData.value.name,
  });

  router.push({
    name: "signin",
  });
};

const formData = ref({
  email: "",
  password: "",
  name: "",
});
</script>

<template>
  <AuthLayout @submit="onSubmit">
    <template #form>
      <Input v-model="formData.email"> Email </Input>
      <Input v-model="formData.name"> Name </Input>
      <Input v-model="formData.password" type="password"> Password </Input>

      <Button type="submit">
        {{ $t("page.signup.btn") }}
      </Button>

      <router-link to="/signin">
        {{ $t("page.signup.go-to-signin") }}
        {{ $t("misc.signin") }}
      </router-link>
    </template>

    <template #title>
      {{ $t("misc.signup") }}
    </template>
  </AuthLayout>
</template>

<style scoped></style>
