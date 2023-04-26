<script setup lang="ts">
import { watch } from "vue";
import { useRoute } from "vue-router";
import { useMenu } from "~/composables/menu";
import { useSize } from "~/composables/size";

const route = useRoute();

const { size } = useSize();
const { isOpen, setMenu } = useMenu();

watch(
  () => route.name,
  () => setMenu(false)
);
</script>

<template>
  <transition name="fade">
    <div
      class="fixed top-[60px] left-0 w-full flex justify-end bg-black bg-opacity-20 backdrop-blur-sm z-20 h-[calc(100vh_-_60px)]"
      v-if="size.isMd && isOpen"
      @click.self="setMenu(false)"
    >
      <div class="bg-zinc-900 border-zinc-800 flex flex-col gap-4 p-4 border w-56 text-white">
        <router-link to="/tickets" class="hover:bg-zinc-800 p-2 rounded-md">
          {{ $t("misc.tickets") }}
        </router-link>

        <router-link to="/clubs/1" class="hover:bg-zinc-800 p-2 rounded-md">
          {{ $t("misc.clubs") }}
        </router-link>

        <router-link to="/trainers/1" class="hover:bg-zinc-800 p-2 rounded-md">
          {{ $t("misc.trainers") }}
        </router-link>
      </div>
    </div>
  </transition>
</template>
