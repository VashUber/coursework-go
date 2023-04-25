<script setup lang="ts">
import { onMounted, ref, toRefs } from "vue";
import { useRouter } from "vue-router";
import { useI18n } from "vue-i18n";
import type { TBuyTicketProps } from "~/composables/modal";
import Input from "~/components/ui/Input.vue";
import Select from "~/components/ui/Select.vue";
import { clubsService } from "~/services/clubs.service";
import { ticketService } from "~/services/ticket.service";
import { TOption } from "types/select";
import { useNotifications } from "~/composables/notification";

const props = defineProps<{
  ctx: TBuyTicketProps;
}>();
const { ctx } = toRefs(props);
const { t } = useI18n();
const { setNotification } = useNotifications();
const router = useRouter();

const creditCard = ref({
  number: "",
  cvc: "",
  expireDate: "",
});

const selectedClub = ref<TOption>();
const optionsClubs = ref<TOption[]>([]);

const setAllClubs = async () => {
  const data = await clubsService.getAllClubsLightWeight();

  const options = data.map((club) => ({ label: club.name, value: club }));
  optionsClubs.value = options;
  selectedClub.value = options[0];
};

const buyTicket = async () => {
  if (Object.entries(creditCard.value).some((elem) => elem[1].length === 0)) {
    setNotification({
      success: false,
      message: t("misc.credit-card.fill-data"),
    });
    return;
  }

  await ticketService.buyTicket({
    ticket: ctx.value.ticket,
    club: selectedClub.value?.value.id,
  });
  router.push("/ticket");
};

onMounted(() => {
  setAllClubs();
});
</script>

<template>
  <div>
    <h2>{{ $t("misc.payment") }}</h2>

    <form @submit.prevent="buyTicket" class="form mt-4">
      <Select v-model="selectedClub" :options="optionsClubs" v-if="selectedClub" />

      <Input v-model="creditCard.number" v-maska data-maska="#### #### #### ####">
        {{ $t("misc.credit-card.number") }}
      </Input>
      <div class="flex gap-2">
        <Input v-model="creditCard.expireDate" v-maska data-maska="##/##" class="flex-shrink-0 min-w-0">
          {{ $t("misc.credit-card.expired-date") }}
        </Input>
        <Input v-model="creditCard.cvc" v-maska data-maska="###" class="flex-1 min-w-min">
          {{ $t("misc.credit-card.cvc") }}
        </Input>
      </div>

      <button class="button button--accent" type="submit">
        {{ $t("misc.buy") }}
      </button>
    </form>
  </div>
</template>
