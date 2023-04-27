<script setup lang="ts">
import { computed } from "vue";
import { useRoute } from "vue-router";
import { clubsService } from "~/services/clubs.service";
import { useLoader } from "~/composables/loader";
import Equipment from "~/components/misc/Equipment.vue";

const route = useRoute();

const clubId = computed(() => +route.params.id || -1);
const { data: club } = useLoader(clubsService.getClubById, clubId);
</script>

<template>
  <div class="page">
    <div class="grid grid-cols-2 gap-4 mb-2">
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

    <h2 class="mb-2">
      {{ $t("misc.equipment") }}
    </h2>
    <div class="flex gap-4 overflow-x-scroll overflow-y-hidden min-h-max">
      <Equipment
        v-for="equipment in club?.equipment"
        :key="equipment.name"
        :equipment="equipment"
        class="flex-shrink-0"
      />
    </div>
  </div>
</template>
