<script setup lang="ts"></script>

<template>
  <div class="layout">
    <layout-header />
    <div class="wrapper layout__content">
      <slot />
    </div>
    <layout-footer />
  </div>
</template>

<style scoped lang="scss">
.layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  height: 1px;

  &__content {
    flex-grow: 1;
  }
}
</style>
