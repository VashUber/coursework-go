<script setup lang="ts">
import { Menu, MenuButton, MenuItems } from "@headlessui/vue";
</script>

<template>
  <Menu as="div" class="relative">
    <MenuButton
      class="inline-flex w-full justify-center rounded-md bg-zinc-800 px-4 py-2 text-sm font-medium text-white hover:bg-opacity-90"
    >
      <slot name="button" />
    </MenuButton>
    <MenuItems
      class="absolute right-0 mt-2 w-56 origin-top-right divide-y divide-gray-200 rounded-md bg-white shadow-lg focus:outline-none"
    >
      <slot name="items" />
    </MenuItems>
  </Menu>
</template>
