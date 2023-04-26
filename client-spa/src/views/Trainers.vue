<script setup lang="ts">
import { computed } from "vue";
import { useRoute } from "vue-router";
import { useLoader } from "~/composables/loader";
import Thumbs from "~/components/misc/Thumbs.vue";
import Pagination from "~/components/ui/Pagination.vue";
import ThumbsLoader from "~/components/misc/ThumbsLoader.vue";
import Search from "~/components/misc/Search.vue";
import { trainersService } from "~/services/trainers.service";

const route = useRoute();

const page = computed(() => +route.params.page || 1);
const searchQ = computed(() => route.query.search || "");
const { data } = useLoader(trainersService.getTrainersPerPage, page, searchQ);

const trainers = computed(() => {
  if (!data.value) return;

  return data.value.trainers.map((trainer) => ({
    id: trainer.id,
    thumb: trainer.thumb,
    title: trainer.name,
  }));
});
</script>

<template>
  <div class="page">
    <div class="flex items-center justify-center">
      <Search />
    </div>

    <Thumbs :thumbs="trainers" to="club" v-if="trainers" :is-vertical="true" />
    <ThumbsLoader :thumbs="6" v-else :is-vertical="true" />
    <Pagination v-if="data?.pages" :pages="data.pages" />
  </div>
</template>
