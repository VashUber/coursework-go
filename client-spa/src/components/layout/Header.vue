<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { useUser } from "~/composables/user";
import Dropdown from "~/components/ui/Dropdown.vue";
import { MenuItem } from "@headlessui/vue";
import { changeLocale } from "~/libs/i18n";
import { auth } from "~/services/auth.service";

const { user } = useUser();
const { locale } = useI18n();

const signout = async () => {
  await auth.signout();

  window.location.replace("/");
};
</script>

<template>
  <div class="bg-zinc-900">
    <div class="wrapper flex items-center h-[60px]">
      <div class="flex items-center justify-center gap-2 ml-auto text-white">
        <div
          @click="changeLocale(locale === 'ru' ? 'en' : 'ru')"
          class="cursor-pointer"
        >
          {{ $t("lang." + (locale === "ru" ? "en" : "ru")) }}
        </div>

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
                    {{ $t("nav.profile") }}
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
                    {{ $t("nav.ticket") }}
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
