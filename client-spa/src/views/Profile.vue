<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useUser } from "~/composables/user";
import Input from "~/components/ui/Input.vue";
import { userService } from "~/services/user.service";
import DatePicker from "~/components/ui/DatePicker.vue";

const { user } = useUser();
const userFormData = ref({
  email: user.value!.email,
  name: user.value!.name,
  birthday: "",
  password: "",
  avatar: "",
});

onMounted(() => {
  userService.getProfileInfo().then((response) => {
    userFormData.value.birthday = response.birthday;
    userFormData.value.avatar = response.avatar;
  });
});
</script>

<template>
  <div class="">
    <h1>
      {{ $t("page.names.profile") }}
    </h1>

    <div class="flex gap-4 items-center">
      <div class="w-40 h-40">
        <img
          v-if="userFormData.avatar"
          class="object-cover w-full h-full rounded-full"
          :src="userFormData.avatar"
        />

        <div
          v-else
          class="bg-orange-400 w-full h-full text-6xl flex items-center justify-center text-white uppercase rounded-full"
        >
          {{ userFormData.name[0] }}
        </div>
      </div>

      <form class="form w-72">
        <Input v-model="userFormData.email">Email</Input>
        <Input v-model="userFormData.name">Name</Input>
        <date-picker v-model="userFormData.birthday" placeholder="Birthday" />
        <Input v-model="userFormData.password">Password</Input>
      </form>
    </div>
  </div>
</template>