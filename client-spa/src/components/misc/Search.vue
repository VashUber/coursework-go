<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import Search from "../icons/Search.vue";

const router = useRouter();
const route = useRoute();
const search = ref("");
const isFocused = ref(false);
let timeoutId: NodeJS.Timeout;

watch(search, () => {
  clearTimeout(timeoutId);

  timeoutId = setTimeout(() => {
    router.push({
      query: {
        ...route.query,
        search: search.value,
      },
    });
  }, 400);
});
onMounted(() => {
  search.value = String(route.query.search || "");
});
</script>

<template>
  <div
    class="flex items-center px-4 py-2 gap-2 border border-gray-400 text-gray-950 bg-slate-50 rounded-md h-10 transition-colors max-[380px]:w-40"
    :class="{
      'border-orange-400': isFocused,
    }"
  >
    <input
      type="text"
      v-model="search"
      @focus="isFocused = true"
      @blur="isFocused = false"
      class="outline-none bg-slate-50 w-full"
      :placeholder="$t('misc.search')"
    />
    <Search class="text-gray-400" />
  </div>
</template>
