<script setup lang="ts">
import { ticketService } from "~/services/ticket.service";
import { useLoader } from "~/composables/loader";
import TicketCard from "~/components/misc/TicketCard.vue";

const { data, isLoading } = useLoader(ticketService.getUserTicket);
</script>

<template>
  <div class="page">
    {{ $t("page.names.ticket") }}
    <div class="py-2 flex items-center justify-center">
      <ticket-card :ticket="data" v-if="data"></ticket-card>
      <div v-else-if="!isLoading">
        {{ $t("components.card.empty") }}
      </div>
    </div>
  </div>
</template>

<style scoped></style>
