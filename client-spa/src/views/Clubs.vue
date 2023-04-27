<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { clubsService } from "~/services/clubs.service";
import Pagination from "~/components/ui/Pagination.vue";
import Thumbs from "~/components/misc/Thumbs.vue";
import Search from "~/components/misc/Search.vue";
import Filters from "~/components/misc/Filters.vue";
import ThumbsLoader from "~/components/misc/ThumbsLoader.vue";
import { useLoader } from "~/composables/loader";
import { TFilters } from "types/filters";

const route = useRoute();

const page = computed(() => +route.params.page || 1);
const searchQ = computed(() => route.query.search || "");
const subway = computed(() => route.query.subway || "");
const { data } = useLoader(clubsService.getClubsPerPage, page, searchQ, subway);

const clubs = computed(() => {
  if (!data.value?.clubs) return;

  return data.value.clubs.map((club) => ({
    id: club.ID,
    thumb: club.thumb,
    title: club.name,
  }));
});

const filters = ref<TFilters>({
  subway: {
    options: [{ label: "Not selected", value: null }],
    model: { label: "Not selected", value: null },
  },
});

onMounted(() => {
  clubsService.getAllSubwayStations().then((res) => {
    const data = res.map((elem) => ({
      label: elem.subway,
      value: elem.subway,
    }));

    filters.value.subway.options = [...filters.value.subway.options, ...data];
  });
});
</script>

<template>
  <div class="page">
    <div class="flex items-center justify-between">
      <Search />

      <Filters :filters="filters" />
    </div>

    <Thumbs :thumbs="clubs" to="club" v-if="clubs" />
    <ThumbsLoader :thumbs="6" v-else />
    <Pagination v-if="data?.pages" :pages="data.pages" />
  </div>
</template>
