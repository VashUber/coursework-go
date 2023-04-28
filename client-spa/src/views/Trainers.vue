<script setup lang="ts">
import { computed, ref } from "vue";
import { useRoute } from "vue-router";
import { useLoader } from "~/composables/loader";
import Thumbs from "~/components/misc/Thumbs.vue";
import Pagination from "~/components/ui/Pagination.vue";
import ThumbsLoader from "~/components/misc/ThumbsLoader.vue";
import Search from "~/components/misc/Search.vue";
import Filters from "~/components/misc/Filters.vue";
import { trainersService } from "~/services/trainers.service";
import { useMeta } from "~/composables/meta";
import { useI18n } from "vue-i18n";

const route = useRoute();

const page = computed(() => +route.params.page || 1);
const searchQ = computed(() => route.query.search || "");
const experienceQ = computed(() => route.query.experience || "");
const { data } = useLoader(trainersService.getTrainersPerPage, page, searchQ, experienceQ);

const trainers = computed(() => {
  if (!data.value) return;

  return data.value.trainers.map((trainer) => ({
    id: trainer.id,
    thumb: trainer.thumb,
    title: trainer.name,
  }));
});

const { setHead } = useMeta();
const { t } = useI18n();
setHead({
  title: t("page.trainers.title"),
  description: t("page.trainers.description"),
});

const filters = ref({
  experience: {
    options: [
      { label: "Experience more than 5 years", value: 5 },
      { label: "Experience more than 3 years", value: 3 },
      { label: "Experience more than 1 year", value: 1 },
      { label: "Not selected", value: null },
    ],
    model: { label: "Not selected", value: null },
  },
});
</script>

<template>
  <div class="page">
    <div class="flex justify-between">
      <Search />

      <Filters :filters="filters" />
    </div>

    <Thumbs :thumbs="trainers" v-if="trainers" :is-vertical="true" />
    <ThumbsLoader :thumbs="6" v-else :is-vertical="true" />
    <Pagination v-if="data?.pages" :pages="data.pages" />
  </div>
</template>
