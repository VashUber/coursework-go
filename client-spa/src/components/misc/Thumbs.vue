<script setup lang="ts">
import { toRefs } from "vue";

const props = defineProps<{
  thumbs: {
    thumb: string;
    title: string;
    id: number;
  }[];
}>();

const { thumbs } = toRefs(props);
</script>

<template>
  <div class="grid gap-4 py-2 lg:grid-cols-3 md:grid-cols-2 grid-cols-1">
    <div v-for="thumb in thumbs" :key="thumb.id" class="flex flex-col gap-2">
      <div class="relative pt-[56%] bg-slate-200">
        <img :src="thumb.thumb" class="absolute top-0 left-0 w-full h-full aspect-video" />
      </div>
      <h2 class="h-6">
        {{ thumb.title }}
      </h2>
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
