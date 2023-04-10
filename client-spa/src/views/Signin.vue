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
      <Input v-model="formData.password"> Password </Input>

      <Button type="submit">Register</Button>
    </template>

    <template #title>
      {{ $t("misc.signin") }}
    </template>
  </AuthLayout>
</template>

<style scoped></style>
