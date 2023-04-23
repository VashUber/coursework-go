<script setup lang="ts">
import { toRefs } from "vue";
import { TNotification } from "~/composables/notification";

const props = defineProps<{
  notification: TNotification;
}>();

const { notification } = toRefs(props);
</script>

<template>
  <div
    class="w-60 h-10 bg-opacity-75 backdrop-blur-sm rounded-md flex items-center p-4 text-white"
    :class="{
      'bg-red-500': !notification.success,
      'bg-green-500': notification.success,
    }"
  >
    {{ notification.message }}
  </div>
</template>
