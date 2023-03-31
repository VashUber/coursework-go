<script setup lang="ts">
const ticket = {
  id: 1,
  title: "Обычный",
  date: "22.02.2023 - 22.02.2024",
};
</script>

<template>
  <div class="tickets page">
    <div class="tickets__wrapper">
      <Ticket :ticket="ticket" />
    </div>
  </div>
</template>

<style scoped lang="scss">
.tickets {
  display: flex;
  align-items: center;
  justify-content: center;

  &__wrapper {
    display: flex;
    flex-wrap: wrap;
    gap: 20px;
    justify-content: center;
  }
}
</style>
