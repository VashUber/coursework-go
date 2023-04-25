<script setup lang="ts">
import { onMounted, ref } from "vue";
import { ticketService } from "~/services/ticket.service";
import TicketCard from "~/components/misc/TicketCard.vue";
import { ITicket } from "types/ticket";

const ticket = ref<ITicket>();

const setTicket = async () => {
  const response = await ticketService.getUserTicket();
  if (!response) return;

  ticket.value = response;
};

onMounted(() => {
  setTicket();
});
</script>

<template>
  <div class="page">
    {{ $t("page.names.ticket") }}
    <div class="py-2 flex items-center justify-center">
      <ticket-card :ticket="ticket" v-if="ticket"></ticket-card>
    </div>
  </div>
</template>

<style scoped></style>
