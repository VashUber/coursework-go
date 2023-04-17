<script setup lang="ts"></script>

<template>
  <div class="h-screen flex flex-col items-center justify-center">
    <h1>404</h1>
    <div>
      {{ $t("page.404.text") }}

      <router-link class="underline text-orange-400" to="/">
        {{ $t("page.404.link") }}
      </router-link>
    </div>
  </div>
</template>

<style scoped></style>
