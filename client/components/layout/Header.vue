<script setup lang="ts">
const switchLocalePath = useSwitchLocalePath();

const { locale } = useI18n();
</script>

<template>
  <div class="header">
    <div class="wrapper">
      <div class="header__language-switcher">
        <NuxtLink :to="switchLocalePath('ru')" v-if="locale === 'en'">
          {{ $t("lang.ru") }}
        </NuxtLink>
        <NuxtLink :to="switchLocalePath('en')" v-else>
          {{ $t("lang.en") }}
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.header {
  height: 60px;
  display: flex;
  align-items: center;
  border-bottom: 1px solid $line;

  &__language-switcher {
    display: flex;
    align-items: center;
    gap: 10px;
    margin-left: auto;
    width: max-content;
  }
}
</style>
