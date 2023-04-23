<script setup lang="ts">
import Head from "~/components/meta/Head.vue";
import Notifications from "./components/misc/Notifications.vue";
</script>

<template>
  <Head />
  <Notifications />
  <router-view />
</template>

<style lang="scss">
body {
  @apply bg-zinc-100;
}
</style>
