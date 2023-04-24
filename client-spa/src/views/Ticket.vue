<script setup lang="ts">
import { onMounted } from "vue";
import { ticketService } from "~/services/ticket.service";

onMounted(() => {
  ticketService.getUserTicket();
});
</script>

<template>
  <div class="page">
    {{ $t("page.names.ticket") }}
  </div>
</template>

<style scoped></style>
