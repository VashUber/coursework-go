<script setup lang="ts">
import { toRefs } from "vue";

const props = defineProps<{
  thumbs: {
    thumb: string;
    title: string;
    id: number;
  }[];
  to: string;
}>();

const { thumbs, to } = toRefs(props);
</script>

<template>
  <div class="grid gap-4 py-2 lg:grid-cols-3 md:grid-cols-2 grid-cols-1">
    <router-link v-for="thumb in thumbs" :key="thumb.id" :to="`/${to}/${thumb.id}`">
      <div class="flex flex-col gap-2">
        <div class="relative pt-[56.25%] bg-slate-200">
          <img :src="thumb.thumb" class="absolute top-0 left-0 w-full h-full rounded-md object-cover" />
        </div>
        <h2 class="h-6">
          {{ thumb.title }}
        </h2>
      </div>
    </router-link>
  </div>
</template>
