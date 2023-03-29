<script setup lang="ts">
const { t } = useI18n();
const localePath = useLocalePath();

useHead({
  title: t("404.seo.title"),
});
</script>

<template>
  <div class="page-404">
    <h1>{{ t("404.seo.h1") }}</h1>

    <router-link :to="localePath('/')">
      {{ t("404.go_to_main") }}
    </router-link>
  </div>
</template>

<style scoped lang="scss">
.page-404 {
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 10px;
  font-weight: bold;
  background: url("./assets/bg/waves.svg") #fcfcfe no-repeat;
  background-size: cover;

  & > h1 {
    font-size: 38px;
    text-align: center;
  }

  & > a {
    color: $accent;
    text-decoration: none;
  }
}
</style>
