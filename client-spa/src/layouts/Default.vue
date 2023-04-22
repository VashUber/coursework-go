<script setup lang="ts">
import Header from "~/components/layout/Header.vue";
import Footer from "~/components/layout/Footer.vue";
import Head from "~/components/meta/Head.vue";
</script>

<template>
  <Head />

  <Header />

  <div class="wrapper">
    <router-view />
  </div>

  <Footer />
</template>

<style scoped></style>
