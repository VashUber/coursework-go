<script setup lang="ts">
const { t } = useI18n();
const input = ref("");

useHead({
  title: t("index.seo.title"),
  meta: [
    {
      property: "og:title",
      content: t("index.seo.title"),
    },
  ],
});
</script>

<template>
  <div>
    <h1>{{ t("index.seo.h1") }}</h1>

    <ui-input v-model="input"></ui-input>
  </div>
</template>

<style scoped></style>
