<script setup lang="ts">
import { onMounted, watch } from "vue";
import { useScrollLock } from "@vueuse/core";
import Head from "~/components/meta/Head.vue";
import Notifications from "./components/misc/Notifications.vue";
import Modal from "./components/modal/Modal.vue";
import { useSize } from "~/composables/size";
import { useMenu } from "./composables/menu";

const isLocked = useScrollLock(document.body);
const { onResize } = useSize();
const { isOpen } = useMenu();

watch(isOpen, (val) => {
  isLocked.value = val;
});

onMounted(() => {
  window.addEventListener("resize", onResize);
});
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
