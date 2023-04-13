<script setup lang="ts">
import { ref } from "vue";
import Input from "~/components/ui/Input.vue";
import Button from "~/components/ui/Button.vue";
import AuthLayout from "./components/auth/AuthLayout.vue";
import { auth } from "~/services/auth.service";

const onSubmit = async () => {
  await auth.signin({
    email: formData.value.email,
    password: formData.value.password,
  });

  window.location.replace("/");
};

const formData = ref({
  email: "",
  password: "",
});
</script>

<template>
  <AuthLayout @submit="onSubmit">
    <template #form>
      <Input v-model="formData.email"> Email </Input>
      <Input v-model="formData.password" type="password"> Password </Input>

      <Button type="submit">
        {{ $t("page.signin.btn") }}
      </Button>

      <router-link to="/signup">
        {{ $t("page.signin.go-to-signup") }}
        {{ $t("misc.signup") }}
      </router-link>
    </template>

    <template #title>
      {{ $t("misc.signin") }}
    </template>
  </AuthLayout>
</template>

<style scoped></style>
