<script setup lang="ts">
import { ref, toRefs } from "vue";
import Select from "../ui/Select.vue";
import Filter from "../icons/Filter.vue";
import { useRoute, useRouter } from "vue-router";
import { TFilters } from "~/shared-types/filters";

const props = defineProps<{
  filters: TFilters;
}>();

const router = useRouter();
const route = useRoute();
const { filters } = toRefs(props);
const isVisible = ref(false);

const toggle = () => {
  isVisible.value = !isVisible.value;
};
const close = () => {
  isVisible.value = false;
};
const setFilter = (filter: string, value: any) => {
  router.push({
    query: {
      ...route.query,
      [filter]: value,
    },
  });
};
</script>

<template>
  <div class="relative">
    <button class="button ml-auto" @click="toggle">
      <Filter></Filter>
    </button>

    <div
      v-if="isVisible"
      class="absolute top-[102%] right-[50%] bg-zinc-900 text-white z-10 p-4 rounded-md"
      v-click-outside="close"
    >
      <Select
        v-for="(filter, key) in filters"
        :key="key"
        v-model="filter.model"
        :options="filter.options"
        @click="setFilter(key, filter.model.value)"
      />
    </div>
  </div>
</template>
