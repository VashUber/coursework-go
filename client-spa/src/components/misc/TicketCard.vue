<script setup lang="ts">
import { toRefs } from "vue";
import { ITicket } from "types/ticket";

const props = defineProps<{
  ticket: ITicket;
}>();

const { ticket } = toRefs(props);
</script>

<template>
  <div class="bg-zinc-900 text-orange-400 w-96 h-60 rounded-lg p-4 shadow-md flex flex-col group card relative">
    <div
      class="group-hover:opacity-100 opacity-0 transition-opacity absolute w-full h-full bg-black bg-opacity-75 top-0 left-0 flex justify-center items-center flex-col gap-4 rounded-lg"
    >
      <div class="text-2xl flex flex-col gap-2 items-center p-2">
        {{ ticket.price }} {{ $t("misc.ruble") }}
        <div class="text-center text-base">
          {{ ticket.start_date }}
          {{ ticket.expired_date }}
        </div>
      </div>
    </div>
    <div class="text-3xl flex justify-between">
      <div>
        {{ ticket.time }}
      </div>
    </div>

    <div class="mt-auto ml-auto">
      {{ ticket. }}
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
