<script setup lang="ts">
import { ticketService } from "~/services/ticket.service";
import { useLoader } from "~/composables/loader";
import TicketCard from "~/components/misc/TicketCard.vue";
import { useMeta } from "~/composables/meta";
import { useI18n } from "vue-i18n";
import { useUser } from "~/composables/user";

const { data, isLoading } = useLoader(ticketService.getUserTicket);

const { user } = useUser();
const { setHead } = useMeta();
const { t } = useI18n();
setHead({
  title: t("page.ticket.title", { v: user.value?.name }),
  description: t("page.ticket.description", { v: user.value?.name }),
});
</script>

<template>
  <div class="page">
    {{ $t("page.names.ticket") }}
    <div class="flex flex-wrap gap-4 justify-center items-center min-h-[calc(100vh_-_120px)]">
      <ticket-card :ticket="data" v-if="data"></ticket-card>
      <div v-else-if="!isLoading">
        {{ $t("components.card.empty") }}
      </div>
    </div>
  </div>
</template>

<style scoped></style>
