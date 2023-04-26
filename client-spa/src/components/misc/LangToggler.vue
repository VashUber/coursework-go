<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { changeLocale } from "~/libs/i18n";

const { locale } = useI18n();
</script>

<template>
  <div @click="changeLocale(locale === 'ru' ? 'en' : 'ru')" class="cursor-pointer">
    {{ $t("lang." + (locale === "ru" ? "en" : "ru")) }}
  </div>
</template>
