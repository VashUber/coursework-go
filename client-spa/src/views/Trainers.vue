<script setup lang="ts">
import { useLoader } from "~/composables/loader";
import { trainersService } from "~/services/trainers.service";
import Thumbs from "~/components/misc/Thumbs.vue";
import ThumbsLoader from "~/components/misc/ThumbsLoader.vue";
import Pagination from "~/components/ui/Pagination.vue";
import { useRoute } from "vue-router";
import { computed } from "vue";

const route = useRoute();

const page = computed(() => +route.params.page || 1);
const { data } = useLoader(trainersService.getTrainersPerPage, page);

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
    <Thumbs :thumbs="trainers" to="club" v-if="trainers" :is-vertical="true" />
    <ThumbsLoader :thumbs="6" v-else :is-vertical="true" />
    <Pagination v-if="data?.pages" :pages="data.pages" />
  </div>
</template>
