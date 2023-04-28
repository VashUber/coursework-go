<script setup lang="ts">
import { IEquipment } from "types/equipment";
import { toRefs } from "vue";

const props = defineProps<{
  equipment: IEquipment;
}>();
const { equipment } = toRefs(props);
</script>

<template>
  <div class="flex flex-col gap-2 w-40">
    <div class="aspect-square bg-slate-200 relative">
      <img :src="equipment.image" class="absolute left-0 top-0 w-full h-full" />
    </div>

    <h2 class="h-6">
      {{ equipment.name }}
    </h2>
  </div>
</template>
