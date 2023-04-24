<script setup lang="ts">
import { onMounted, ref, toRefs } from "vue";
import type { TBuyTicketProps } from "~/composables/modal";
import Input from "~/components/ui/Input.vue";
import Select from "~/components/ui/Select.vue";
import { clubsService } from "~/services/clubs.service";
import { TOption } from "types/select";

const props = defineProps<{
  ctx: TBuyTicketProps;
}>();
const { ctx } = toRefs(props);

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

onMounted(() => {
  setAllClubs();
});
</script>

<template>
  <div>
    <h2>{{ $t("misc.payment") }}</h2>

    <form @submit.prevent="" class="form mt-4">
      <Select v-model="selectedClub" :options="optionsClubs" v-if="selectedClub" />

      <Input v-model="creditCard.number" v-maska data-maska="#### #### #### ####">
        {{ $t("misc.credit-card.number") }}
      </Input>
      <Input v-model="creditCard.expireDate" v-maska data-maska="##/##">
        {{ $t("misc.credit-card.expired-date") }}
      </Input>
      <Input v-model="creditCard.cvc" v-maska data-maska="###">
        {{ $t("misc.credit-card.cvc") }}
      </Input>
    </form>
  </div>
</template>

<style scoped></style>
