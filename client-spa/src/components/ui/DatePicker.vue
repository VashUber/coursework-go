<script setup lang="ts">
import DatePicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";
import { useI18n } from "vue-i18n";

const { locale } = useI18n();
</script>

<template>
  <date-picker
    :locale="locale"
    :enable-time-picker="false"
    input-class-name="datepicker__input"
    :max-date="new Date()"
    vertical
  />
</template>

<style lang="scss">
.datepicker {
  &__input {
    @apply h-[50px] border-gray-400 border-b border-x-0 border-t-0 rounded-tl-[5px] rounded-tr-[5px] rounded-none placeholder:text-gray-400;
  }
}
</style>
