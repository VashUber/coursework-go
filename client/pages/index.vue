<script setup lang="ts">
const { t } = useI18n();

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
    {{ t("index.seo.h1") }}
  </div>
</template>

<style scoped></style>
