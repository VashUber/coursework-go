<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { clubsService } from "~/services/clubs.service";
import { IClub } from "~/shared-types/club";

const route = useRoute();

const club = ref<IClub>();

const setClubById = async (id: number) => {
  club.value = await clubsService.getClubById(id);
};

onMounted(() => {
  const clubId = +route.params.id || -1;
  setClubById(clubId);
});
</script>

<template>
  <div class="page">
    <div class="grid grid-cols-2 gap-4">
      <div class="pt-[56.25%] h-full bg-slate-200 relative">
        <img :src="club.image" v-if="club" class="absolute top-0 left-0 w-full h-full rounded-md" />
      </div>
      <div>
        <p v-if="club?.address">
          {{ club.address.home + " " + club.address.street }}
        </p>
        <p>{{ club?.info }}</p>
      </div>
    </div>
  </div>
</template>
