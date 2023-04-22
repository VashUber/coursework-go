<script setup lang="ts">
import { toRefs } from "vue";

const props = defineProps<{
  thumbs: number;
}>();

const { thumbs } = toRefs(props);
</script>

<template>
  <div class="grid gap-4 py-2 lg:grid-cols-3 md:grid-cols-2 grid-cols-1">
    <div v-for="thumb in thumbs" :key="thumb" class="flex flex-col gap-2 animate-pulse">
      <div class="relative pt-[56%] bg-slate-300 rounded-md"></div>
      <h2 class="bg-slate-300 w-full h-6"></h2>
    </div>
  </div>
</template>
