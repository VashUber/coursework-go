<script setup lang="ts">
const props = defineProps<{
  modelValue: string;
}>();
const emits = defineEmits<{
  (e: "update:modelValue", value: string): void;
}>();

const { modelValue } = toRefs(props);

const onChange = (e: Event) => {
  emits("update:modelValue", (e.target as HTMLInputElement).value);
};
</script>

<template>
  <input type="text" :value="modelValue" @input="onChange" />
</template>

<style scoped></style>
