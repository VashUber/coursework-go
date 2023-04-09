<script setup lang="ts">
import { ref } from "vue";
import Input from "~/components/ui/Input.vue";
import Button from "~/components/ui/Button.vue";

const onSubmit = () => {
  console.log(formData.value);
};

const formData = ref({
  email: "",
  password: "",
  name: "",
  surname: "",
  age: "",
});
</script>

<template>
  <div
    class="w-full flex items-center justify-center gap-40 h-[calc(100vh_-_60px)] mx-4"
  >
    <form
      class="flex-grow flex flex-col gap-2 items-center"
      @submit.prevent="onSubmit"
    >
      <Input v-model="formData.email"> Email </Input>
      <Input v-model="formData.name"> Name </Input>
      <Input v-model="formData.surname"> Surname </Input>
      <Input v-model="formData.age"> Age </Input>
      <Input v-model="formData.password"> Password </Input>

      <Button type="submit">Register</Button>
    </form>
    <div
      class="flex-grow aspect-[2/3] overflow-hidden h-[calc(100vh_-_60px)] relative"
    >
      <h1
        class="absolute text-white z-20 top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2"
      >
        {{ $t("misc.signup") }}
      </h1>
      <img
        src="/auth-page-img.jpeg"
        class="w-full h-full object-cover brightness-[60%]"
      />
    </div>
  </div>
</template>

<style scoped></style>
