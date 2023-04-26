<script setup lang="ts">
import { onMounted, ref, toRefs } from "vue";
import { useRoute, useRouter } from "vue-router";

const props = defineProps<{
  pages: number;
}>();

const router = useRouter();
const route = useRoute();

const count = 5;
const { pages } = toRefs(props);
const currPages = ref<number[]>([]);

const getNewPagesArray = (page: number) => {
  const arr = [];

  if (page + count > pages.value) {
    for (let i = pages.value - count + 1; i <= pages.value; i++) {
      arr.push(i);
    }
  } else {
    for (let i = page; i < page + count; i++) {
      arr.push(i);
    }
  }

  return arr;
};

const onPagination = (page: number) => {
  if (page < 1 || page > pages.value) return;

  router.push({
    ...route,
    params: {
      page,
    },
  });

  if (pages.value < count) return;

  currPages.value = getNewPagesArray(page);
};

onMounted(() => {
  const pageFromRoute = route.params.page ? +route.params.page : 1;

  currPages.value =
    pages.value < count
      ? Array(pages.value)
          .fill(0)
          .map((_, idx) => idx + 1)
      : getNewPagesArray(pageFromRoute);
});
</script>

<template>
  <div class="flex items-center justify-center">
    <div class="flex gap-1">
      <button class="btn" @click="onPagination(currPages[0] - 1)" :disabled="currPages[0] === 1">
        {{ "<" }}
      </button>
      <button v-for="page in currPages" :key="page" class="btn" @click="onPagination(page)">
        {{ page }}
      </button>
      <button class="btn" @click="onPagination(currPages[0] + 1)" :disabled="currPages[currPages.length - 1] === pages">
        {{ ">" }}
      </button>
    </div>
  </div>
</template>

<style scoped lang="scss">
.btn {
  @apply w-10 h-10 bg-zinc-900 text-white flex items-center justify-center rounded-xl disabled:bg-zinc-800 disabled:text-gray-200;
}
</style>
