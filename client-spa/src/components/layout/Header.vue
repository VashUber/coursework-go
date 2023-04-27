<script setup lang="ts">
import { useUser } from "~/composables/user";
import { useSize } from "~/composables/size";
import LangToggler from "../misc/LangToggler.vue";
import Dropdown from "~/components/ui/Dropdown.vue";
import { MenuItem } from "@headlessui/vue";
import { authService } from "~/services/auth.service";
import MenuClosed from "~/components/icons/MenuClosed.vue";
import MenuOpen from "~/components/icons/MenuOpen.vue";
import { useMenu } from "~/composables/menu";

const { size } = useSize();
const { user } = useUser();
const { isOpen, toggle } = useMenu();

const signout = async () => {
  await authService.signout();

  window.location.replace("/");
};
</script>

<template>
  <div class="bg-zinc-900">
    <div class="wrapper flex items-center h-[60px] gap-4 text-white">
      <router-link to="/">
        <img src="/logo.svg" class="h-[50px] aspect-[117/50]" />
      </router-link>

      <div class="flex items-center gap-4 bg-zinc-800 h-[36px] py-2 px-4 rounded-md" v-if="!size.isMd">
        <router-link to="/tickets">
          {{ $t("misc.tickets") }}
        </router-link>

        <router-link to="/clubs/1">
          {{ $t("misc.clubs") }}
        </router-link>

        <router-link to="/trainers/1">
          {{ $t("misc.trainers") }}
        </router-link>
      </div>

      <div class="flex items-center justify-center gap-4 ml-auto text-white">
        <LangToggler v-if="!size.isMd" />

        <Dropdown v-if="user">
          <template #button>
            {{ user.name }}
          </template>

          <template #items>
            <div class="p-1">
              <MenuItem v-slot="{ active, close }">
                <router-link
                  :class="{
                    'dropdown__item--active': active,
                  }"
                  class="dropdown__item"
                  :to="{
                    name: 'profile',
                  }"
                >
                  <button @click="close" class="w-full text-left">
                    {{ $t("page.names.profile") }}
                  </button>
                </router-link>
              </MenuItem>

              <MenuItem v-slot="{ active, close }">
                <router-link
                  :class="{
                    'dropdown__item--active': active,
                  }"
                  class="dropdown__item"
                  :to="{
                    name: 'ticket',
                  }"
                >
                  <button @click="close" class="w-full text-left">
                    {{ $t("page.names.ticket") }}
                  </button>
                </router-link>
              </MenuItem>
            </div>

            <div class="p-1">
              <MenuItem v-slot="{ active }">
                <button
                  :class="{
                    'dropdown__item--active': active,
                  }"
                  class="dropdown__item"
                  @click="signout"
                >
                  {{ $t("misc.signout") }}
                </button>
              </MenuItem>
            </div>
          </template>
        </Dropdown>

        <router-link to="/signin" v-else>
          {{ $t("misc.signin") }}
        </router-link>

        <button class="button button--second" v-if="size.isMd" @click="toggle">
          <MenuClosed v-if="!isOpen" />
          <MenuOpen v-else />
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.dropdown {
  &__item {
    @apply flex w-full items-center rounded-md px-2 py-2 text-sm text-gray-900;

    &--active {
      @apply bg-orange-400 text-white;
    }
  }
}
</style>
