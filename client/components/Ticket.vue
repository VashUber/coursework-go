<script setup lang="ts">
defineProps<{
  ticket: {
    id: number;
    title: string;
    date: string;
  };
}>();
</script>

<template>
  <div class="ticket" :key="ticket.id">
    <div class="ticket__title">
      {{ ticket.title }}
    </div>

    <div class="ticket__date">
      {{ ticket.date }}
    </div>

    <div class="ticket__info">info</div>
  </div>
</template>

<style scoped lang="scss">
$ticket: #f65200;
.ticket {
  $root: &;
  background: url("@/assets/bg/waves-ticket.svg") $ticket no-repeat;
  background-size: cover;
  width: 350px;
  height: 200px;
  border-radius: 20px;
  padding: 20px;
  box-shadow: $shadow;
  display: flex;
  flex-direction: column;
  position: relative;

  &__info {
    opacity: 0;
    position: absolute;
    top: 0;
    right: 0;
    transition: opacity $transition ease-in-out;
    background: $backdrop;
    width: 100%;
    height: 100%;
    border-radius: 20px;
    padding: 20px;
    backdrop-filter: blur(2px);
  }

  & > * {
    color: #fff;
  }

  &__title {
    font-size: 20px;
  }

  &__date {
    margin-top: auto;
    margin-left: auto;
    text-shadow: 0px 2px 0px rgba(0, 0, 0, 0.4);
  }

  &:hover {
    #{$root}__info {
      opacity: 1;
    }
  }
}
</style>
