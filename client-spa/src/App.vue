<script setup lang="ts">
import { useI18n } from "vue-i18n";
import Default from "./layouts/Default.vue";
const { t } = useI18n();
</script>

<template>
  <Default>
    <div>
      {{ t("hello", { v: "egor" }) }}
    </div>
  </Default>
</template>

<style scoped></style>
