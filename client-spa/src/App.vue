<script setup lang="ts">
import Head from "~/components/meta/Head.vue";
import Notifications from "./components/misc/Notifications.vue";
import Modal from "./components/modal/Modal.vue";
</script>

<template>
  <Head />
  <Modal />
  <Notifications />
  <router-view />
</template>

<style lang="scss">
body {
  @apply bg-zinc-100;
}
</style>
