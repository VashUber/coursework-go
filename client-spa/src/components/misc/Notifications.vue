<template>
  <div class="absolute top-4 right-4 flex flex-col gap-4">
    <transition-group name="list">
      <Notification v-for="notify in notifications" :notification="notify" :key="notify.id" />
    </transition-group>
  </div>
</template>

<script setup lang="ts">
import { useNotifications } from "~/composables/notification";
import Notification from "~/components/misc/Notification.vue";

const { notifications } = useNotifications();
</script>
