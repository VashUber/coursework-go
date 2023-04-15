<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import DatePicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";
import Input from "~/components/ui/Input.vue";
import Button from "~/components/ui/Button.vue";
import AuthLayout from "./components/auth/AuthLayout.vue";
import { auth } from "~/services/auth.service";
import { useI18n } from "vue-i18n";

const router = useRouter();
const { locale } = useI18n();

const onSubmit = async () => {
  await auth.signup({
    email: formData.value.email,
    password: formData.value.password,
    name: formData.value.name,
    birthday: formData.value.birthday,
  });

  router.push({
    name: "signin",
  });
};

const formData = ref({
  email: "",
  password: "",
  name: "",
  birthday: 0,
});
</script>

<template>
  <AuthLayout @submit="onSubmit">
    <template #form>
      <Input v-model="formData.email"> Email </Input>
      <Input v-model="formData.name"> Name </Input>
      <date-picker
        :locale="locale"
        :enable-time-picker="false"
        v-model="formData.birthday"
        placeholder="Birthday"
        input-class-name="datepicker__input"
        :max-date="new Date()"
        vertical
      />
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

<style lang="scss">
.datepicker__input {
  @apply h-[50px] border-gray-400 border-b border-x-0 border-t-0 rounded-tl-[5px] rounded-tr-[5px] rounded-none placeholder:text-gray-400;
}
</style>
