<script setup lang="ts">
import { toRefs } from "vue";

const props = withDefaults(
  defineProps<{
    thumbs: number;
    isVertical: boolean;
  }>(),
  {
    isVertical: false,
  }
);

const { thumbs, isVertical } = toRefs(props);
</script>

<template>
  <div class="grid gap-4 py-2" :class="[isVertical ? 'vertical' : 'horizontal']">
    <div v-for="thumb in thumbs" :key="thumb" class="flex flex-col gap-2 animate-pulse">
      <div class="relative bg-slate-300 rounded-md" :class="[isVertical ? 'aspect-[9/16]' : 'aspect-video']"></div>
      <h2 class="bg-slate-300 w-full h-6"></h2>
    </div>
  </div>
</template>

<style scoped lang="scss">
.vertical {
  @apply lg:grid-cols-6 md:grid-cols-3 grid-cols-1;
}

.horizontal {
  @apply lg:grid-cols-3 md:grid-cols-2 grid-cols-1;
}
</style>
