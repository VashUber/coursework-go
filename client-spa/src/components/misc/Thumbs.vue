<script setup lang="ts">
import { computed, toRefs } from "vue";

const props = withDefaults(
  defineProps<{
    thumbs: {
      thumb: string;
      title: string;
      id: number;
    }[];
    to?: string;
    isVertical?: boolean;
  }>(),
  {
    isVertical: false,
  }
);

const { thumbs, to, isVertical } = toRefs(props);

const component = computed(() => {
  if (to?.value) return "router-link";

  return "div";
});
</script>

<template>
  <div class="grid gap-4 py-2" :class="[isVertical ? 'vertical' : 'horizontal']">
    <component :is="component" v-for="thumb in thumbs" :key="thumb.id" :to="to ? `/${to}/${thumb.id}` : null">
      <div class="flex flex-col gap-2">
        <div class="relative bg-slate-200" :class="[isVertical ? 'aspect-[9/16]' : 'aspect-video']">
          <img :src="thumb.thumb" class="absolute top-0 left-0 w-full h-full rounded-md object-cover" />
        </div>
        <h2 class="h-6 overflow-hidden whitespace-nowrap text-ellipsis">
          {{ thumb.title }}
        </h2>
      </div>
    </component>
  </div>
</template>

<style scoped lang="scss">
.vertical {
  @apply lg:grid-cols-6 sm:grid-cols-3 min-[480px]:grid-cols-2 grid-cols-1;
}

.horizontal {
  @apply lg:grid-cols-3 min-[480px]:grid-cols-2 grid-cols-1;
}
</style>
