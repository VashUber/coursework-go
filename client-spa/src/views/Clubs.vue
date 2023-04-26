<script setup lang="ts">
import { computed, ref } from "vue";
import { useRoute } from "vue-router";
import { clubsService } from "~/services/clubs.service";
import Pagination from "~/components/ui/Pagination.vue";
import Thumbs from "~/components/misc/Thumbs.vue";
import Search from "~/components/misc/Search.vue";
import ThumbsLoader from "~/components/misc/ThumbsLoader.vue";
import { useLoader } from "~/composables/loader";

const route = useRoute();
const page = computed(() => +route.params.page || 1);
const searchQ = computed(() => route.query.search || "");
const { data } = useLoader(clubsService.getClubsPerPage, page, searchQ);

const search = ref("");

const clubs = computed(() => {
  if (!data.value?.clubs) return;

  return data.value.clubs.map((club) => ({
    id: club.ID,
    thumb: club.thumb,
    title: club.name,
  }));
});
</script>

<template>
  <div class="page">
    <div class="flex items-center justify-center">
      <Search v-model="search" />
    </div>

    <Thumbs :thumbs="clubs" to="club" v-if="clubs" />
    <ThumbsLoader :thumbs="6" v-else />
    <Pagination v-if="data?.pages" :pages="data.pages" />
  </div>
</template>
