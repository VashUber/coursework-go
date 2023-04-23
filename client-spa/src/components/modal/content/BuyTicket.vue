<script setup lang="ts">
import { toRefs } from "vue";
import { TBuyTicketProps } from "~/composables/modal";

const props = defineProps<{
  ctx: TBuyTicketProps;
}>();

const { ctx } = toRefs(props);
</script>

<template>
  <div>buy {{ ctx.ticket }}</div>
</template>

<style scoped></style>
