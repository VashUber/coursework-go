<script setup lang="ts">
import { toRefs } from "vue";
import { ITicketPreview } from "types/ticket-preview";

const props = defineProps<{
  ticket: ITicketPreview;
}>();
const emits = defineEmits<{
  (e: "buy", value: number): void;
}>();

const { ticket } = toRefs(props);

const handleBuy = (v: number) => {
  emits("buy", v);
};
</script>

<template>
  <div class="bg-zinc-900 text-orange-400 w-96 h-60 rounded-lg p-4 shadow-md flex flex-col group card relative">
    <div
      class="group-hover:opacity-100 opacity-0 transition-opacity absolute w-full h-full bg-black bg-opacity-75 top-0 left-0 flex justify-center items-center flex-col gap-4 rounded-lg"
    >
      <div class="text-2xl flex flex-col gap-2 items-center p-2">
        {{ ticket.price }} {{ $t("misc.ruble") }}
        <div class="text-center text-base">
          {{ ticket.info }}
        </div>
      </div>
      <button class="button button--accent" @click="handleBuy(ticket.time)">
        {{ $t("misc.buy") }}
      </button>
    </div>
    <div class="text-3xl flex justify-between">
      <div>
        {{ ticket.title }}
      </div>

      <div>
        {{ ticket.time }}
      </div>
    </div>

    <div class="mt-auto ml-auto">
      {{ ticket.ID }}
    </div>
  </div>
</template>

<style scoped lang="scss">
.card {
  background-image: url("/logo.svg");
  background-repeat: no-repeat;
  background-size: 234px 100px;
  background-position: center;
}
</style>
