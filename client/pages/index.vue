<script setup lang="ts">
const { t } = useI18n();
const input = ref("");

const data = await useAsyncData(() => $http("http://localhost:3055/api"));

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
    {{ data }}
    <ui-input v-model="input"></ui-input>
  </div>
</template>
