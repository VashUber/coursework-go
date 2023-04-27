<script setup lang="ts">
import Header from "~/components/layout/Header.vue";
import Footer from "~/components/layout/Footer.vue";
import MobileMenu from "~/components/layout/MobileMenu.vue";
</script>

<template>
  <Header />
  <MobileMenu />

  <div class="wrapper">
    <router-view />
  </div>

  <Footer />
</template>

<style scoped></style>
