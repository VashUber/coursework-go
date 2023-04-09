<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { changeLocale } from "~/libs/i18n";

const { locale } = useI18n();
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

        <router-link to="/signin">
          {{ $t("misc.signin") }}
        </router-link>
      </div>
    </div>
  </div>
</template>
