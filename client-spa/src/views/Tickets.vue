<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import type { ITicketPreview } from "types/ticket-preview";
import TicketPreviewCard from "~/components/misc/TicketPreviewCard.vue";
import { ticketsPreviewService } from "~/services/tickets-preview.service";
import { useUser } from "~/composables/user";
import { useNotifications } from "~/composables/notification";
import { useI18n } from "vue-i18n";
import { useModal } from "~/composables/modal";
import { useMeta } from "~/composables/meta";

const { user } = useUser();
const router = useRouter();
const { t } = useI18n();
const { setModal } = useModal();
const ticketsPreview = ref<ITicketPreview[]>();
const { setNotification } = useNotifications();

onMounted(() => {
  ticketsPreviewService.getAllTicketsPreview().then((res) => {
    ticketsPreview.value = res;
  });
});

const onBuy = async (ticket: number) => {
  if (!user.value) {
    router.push({
      name: "signin",
    });
    setNotification({
      message: t("notifications.signin"),
      success: false,
    });
    return;
  }

  setModal({
    type: "BuyTicket",
    props: {
      ticket,
    },
  });
};

const { setHead } = useMeta();
setHead({
  title: t("page.tickets.title"),
  description: t("page.tickets.description"),
});
</script>

<template>
  <div class="page flex justify-center items-center">
    <div class="flex flex-wrap gap-4 justify-center items-center">
      <ticket-preview-card v-for="ticket in ticketsPreview" :key="ticket.ID" :ticket="ticket" @buy="onBuy" />
    </div>
  </div>
</template>

<style scoped></style>
