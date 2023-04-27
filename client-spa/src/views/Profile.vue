<script setup lang="ts">
import { ref, unref, watch } from "vue";
import dayjs from "dayjs";
import { useUser } from "~/composables/user";
import Input from "~/components/ui/Input.vue";
import { userService } from "~/services/user.service";
import Upload from "~/components/icons/Upload.vue";
import { useLoader } from "~/composables/loader";

const { user } = useUser();
const { data, refetch } = useLoader(userService.getProfileInfo);

const userFormData = ref({
  email: user.value!.email,
  name: user.value!.name,
  birthday: "",
  password: "",
  avatar: "",
});

const onSubmit = () => {
  const body = unref(userFormData);

  const [d, m, y] = body.birthday.split(".");

  userService.updateProfileInfo({
    ...body,
    birthday: dayjs([m, d, y].join("-")).format(),
  });
};

const onAvatarUpload = async (e: Event) => {
  const target = e.target as HTMLInputElement;
  const file = target.files![0];

  const formData = new FormData();
  formData.append("avatar", file);
  await userService.uploadAvatar(formData);
  await refetch();
};

watch(data, (user) => {
  if (!user) return;

  userFormData.value.avatar = user.avatar;
  userFormData.value.birthday = dayjs(user.birthday).format("DD.MM.YYYY");
});
</script>

<template>
  <div class="page">
    <h1>
      {{ $t("page.names.profile") }}
    </h1>

    <div class="flex justify-center">
      <div
        class="flex gap-8 items-center justify-center border border-gray-300 bg-white shadow-lg p-6 rounded-lg max-[640px]:flex-wrap"
      >
        <div class="w-40 h-40 relative">
          <img v-if="userFormData.avatar" class="object-cover w-full h-full rounded-full" :src="userFormData.avatar" />

          <div
            v-else
            class="bg-orange-400 w-full h-full text-6xl flex items-center justify-center text-white uppercase rounded-full"
          >
            {{ userFormData.name[0] }}
          </div>

          <div class="absolute top-2 right-2">
            <input type="file" id="file" class="hidden" @change="onAvatarUpload" />
            <label for="file" class="bg-zinc-800 bg-opacity-70 block p-2 rounded-full cursor-pointer">
              <Upload class="w-6 h-6 fill-white" />
            </label>
          </div>
        </div>

        <form class="form w-72" @submit.prevent="onSubmit">
          <Input v-model="userFormData.email">Email</Input>
          <Input v-model="userFormData.name">Name</Input>
          <Input v-model="userFormData.birthday" v-maska data-maska="##.##.####">Birthday</Input>
          <Input v-model="userFormData.password" type="password"> Password </Input>

          <button class="button" type="submit">Change</button>
        </form>
      </div>
    </div>
  </div>
</template>
