<script setup lang="ts">
import Header from "~/components/layout/Header.vue";
import Head from "~/components/meta/Head.vue";
</script>

<template>
  <Head />

  <Header></Header>

  <div class="wrapper">
    <router-view />
  </div>
</template>

<style scoped></style>
