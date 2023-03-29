<script setup lang="ts">
const switchLocalePath = useSwitchLocalePath();
</script>

<template>
  <div>
    <div style="display: flex; gap: 20px">
      <NuxtLink :to="switchLocalePath('ru')">ru</NuxtLink>
      <NuxtLink :to="switchLocalePath('en')">en</NuxtLink>
    </div>

    <slot />
  </div>
</template>

<style scoped></style>
