<template>
  <div>
    {{ $t("clubs.seo.h1") }}
  </div>
</template>

<script setup lang="ts">
const { t } = useI18n();

useHead({
  title: t("clubs.seo.title"),
});
</script>

<style scoped></style>
