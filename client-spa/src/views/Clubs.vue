<script setup lang="ts">
import { computed, onMounted, ref, watch } from "vue";
import { useRoute } from "vue-router";
import { clubsService } from "~/services/clubs.service";
import Pagination from "~/components/ui/Pagination.vue";
import Thumbs from "~/components/misc/Thumbs.vue";
import ThumbsLoader from "~/components/misc/ThumbsLoader.vue";
import { IClub } from "~/shared-types/club";

const route = useRoute();
const clubsStorage = ref<{
  clubs: IClub[];
  pages: number;
  isLoading: boolean;
}>({
  clubs: [],
  pages: 0,
  isLoading: false,
});

const setClubsPerPage = async (page: number) => {
  clubsStorage.value.isLoading = true;
  const data = await clubsService.getClubsPerPage(page);
  clubsStorage.value = { ...data, isLoading: false };
};

const clubs = computed(() => {
  return clubsStorage.value.clubs.map((club) => ({
    id: club.ID,
    thumb: club.thumb,
    title: club.name,
  }));
});

watch(
  () => route.params,
  (params) => {
    const page = +params.page || 1;
    setClubsPerPage(page);
  }
);

onMounted(() => {
  const page = +route.params.page || 1;
  setClubsPerPage(page);
});
</script>

<template>
  <div class="page">
    <Thumbs :thumbs="clubs" to="club" v-if="!clubsStorage.isLoading" />
    <ThumbsLoader :thumbs="6" v-else />
    <Pagination v-if="clubsStorage.pages > 0" :pages="clubsStorage.pages" />
  </div>
</template>

<style scoped></style>
