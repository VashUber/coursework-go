<script setup lang="ts">
import { useMeta } from "~/composables/meta";

const { head } = useMeta();
</script>

<template>
  <teleport to="head title">
    {{ head.title }}
  </teleport>

  <teleport to="head">
    <meta name="description" :content="head.description" />
  </teleport>
</template>
