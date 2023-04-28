<script setup lang="ts">
import { computed, watch } from "vue";
import { useRoute } from "vue-router";
import { clubsService } from "~/services/clubs.service";
import { useLoader } from "~/composables/loader";
import Equipment from "~/components/misc/Equipment.vue";
import { useMeta } from "~/composables/meta";
import { useI18n } from "vue-i18n";

const route = useRoute();

const clubId = computed(() => +route.params.id || -1);
const { data: club } = useLoader(clubsService.getClubById, clubId);

const { setHead } = useMeta();
const { t } = useI18n();
watch(club, (value) => {
  if (!value) return;

  setHead({
    title: t("page.club.title", {
      v: value.name,
    }),
    description: t("page.club.description", { v: value.name }),
  });
});
</script>

<template>
  <div class="page">
    <div class="grid grid-cols-1 gap-4 mb-2 md:grid-cols-2">
      <div class="aspect-video h-full bg-slate-200 relative">
        <img :src="club.image" v-if="club" class="absolute top-0 left-0 w-full h-full rounded-md" />
      </div>
      <div class="border border-gray-300 bg-white shadow-lg p-6 rounded-lg">
        <template v-if="club?.address">
          <p>Address: {{ club.address.home + " " + club.address.street }}</p>
          <p></p>
          <p>Subway station: {{ club?.address.subway }}</p>
        </template>

        <div>
          {{ club?.info }}
        </div>
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
