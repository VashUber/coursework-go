<template>
  <div>
    {{ $t("hello") }}

    <button @click="changeLocale">{{ $i18n.global.locale }}</button>
  </div>
</template>

<script setup>
const { $i18n } = useNuxtApp();

const lang = $i18n.global.locale;

const changeLocale = () => {
  if (lang.value === "en") {
    $i18n.global.locale.value = "ru";
    return;
  }

  $i18n.global.locale.value = "en";
};
</script>
