<script setup lang="ts">
import { ref } from "vue";
import Input from "~/components/ui/Input.vue";
import AuthLayout from "./components/auth/AuthLayout.vue";
import { authService } from "~/services/auth.service";

const onSubmit = async () => {
  await authService.signin({
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

      <button type="submit" class="button">
        {{ $t("page.signin.btn") }}
      </button>

      <router-link to="/signup" class="text-center">
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
