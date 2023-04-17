<script setup lang="ts">
import TicketPreviewCard from "~/components/misc/TicketPreviewCard.vue";

const mockTickets = [
  {
    id: 1,
    time: "1",
    price: 14000,
    title: "Title",
    info: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt",
  },
  {
    id: 1,
    time: "1",
    price: 14000,
    title: "Title",
    info: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt",
  },
  {
    id: 1,
    time: "1",
    price: 14000,
    title: "Title",
    info: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt",
  },
];
</script>

<template>
  <div class="flex min-h-screen justify-center items-center py-10">
    <div class="flex flex-wrap gap-4 justify-center">
      <ticket-preview-card
        class="hover:-translate-y-6 transition-transform"
        v-for="ticket in mockTickets"
        :key="ticket.id"
        :ticket="ticket"
      />
    </div>
  </div>
</template>

<style scoped></style>
