<template>
  <div>
    <nuxt-link :to="localePath('club')">
      {{ $t("samara") }}
    </nuxt-link>
    <nuxt-link :to="localePath('/')">
      {{ $t("hello") }}
    </nuxt-link>
    <NuxtPage />
  </div>
</template>

<script setup lang="ts">
const localePath = useLocalePath();
</script>
