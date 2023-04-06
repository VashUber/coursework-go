<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { changeLocale } from "~/libs/i18n";
import Head from "~/components/meta/Head.vue";

const { locale } = useI18n();
</script>

<template>
  <Head />

  <div>
    <div @click="changeLocale('ru')" v-if="locale === 'en'">ru</div>
    <div @click="changeLocale('en')" v-else>en</div>
  </div>

  <slot></slot>
</template>

<style scoped></style>
