<script setup lang="ts">
const { t } = useI18n();

definePageMeta({
  auth: true,
});
useHead({
  title: t("clubs.seo.title"),
});
</script>

<template>
  <h1>
    {{ $t("clubs.seo.h1") }}
  </h1>
</template>

<style scoped></style>
