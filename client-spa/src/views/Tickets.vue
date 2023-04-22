<script setup lang="ts">
import { onMounted, ref } from "vue";
import type { ITicketPreview } from "types/ticket-preview";
import TicketPreviewCard from "~/components/misc/TicketPreviewCard.vue";
import { ticketsPreviewService } from "~/services/tickets-preview.service";

const ticketsPreview = ref<ITicketPreview[]>();

onMounted(() => {
  ticketsPreviewService.getAllTicketsPreview().then((res) => {
    ticketsPreview.value = res;
  });
});
</script>

<template>
  <div class="page flex justify-center items-center">
    <div class="flex flex-wrap gap-4 justify-center">
      <ticket-preview-card
        class="hover:-translate-y-6 transition-transform"
        v-for="ticket in ticketsPreview"
        :key="ticket.ID"
        :ticket="ticket"
      />
    </div>
  </div>
</template>

<style scoped></style>
