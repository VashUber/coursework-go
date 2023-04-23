<script setup lang="ts">
import { computed, defineAsyncComponent } from "vue";
import { useModal } from "~/composables/modal";

const { currModal, close } = useModal();

const BuyTicket = defineAsyncComponent(() => import("./content/BuyTicket.vue"));
const components = {
  BuyTicket,
};

const currComponent = computed(() => {
  if (!currModal.value) return;

  return components[currModal.value.type];
});
</script>

<template>
  <transition name="fade">
    <div
      class="fixed w-full h-screen top-0 left-0 bg-black bg-opacity-20 backdrop-blur-sm flex justify-center items-center z-10"
      v-if="currComponent"
      @click="close"
    >
      <div class="bg-zinc-800 text-white w-96 min-h-[300px] p-4 rounded-md">
        <Suspense>
          <template #fallback> Loading... </template>
          <component :is="currComponent" :ctx="currModal?.props"></component>
        </Suspense>
      </div>
    </div>
  </transition>
</template>
