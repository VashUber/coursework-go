<script setup lang="ts">
import { toRefs } from "vue";

const props = defineProps<{
  modelValue: string | number;
}>();

const emits = defineEmits<{
  (e: "update:modelValue", value: string): void;
}>();

const { modelValue } = toRefs(props);
const onChange = (e: Event) => {
  const value = (e.target as HTMLInputElement).value;

  emits("update:modelValue", value);
};
</script>

<template>
  <label class="relative">
    <input
      :value="modelValue"
      @input="onChange"
      class="min-w-[250px] outline-none px-4 pb-2 pt-4 border-gray-400 border-b rounded-tl-[5px] rounded-tr-[5px] focus:border-blue-700"
      placeholder="placeholder"
    />

    <div
      class="absolute top-1/2 -translate-y-1/2 left-4 text-gray-400 label transition-all"
    >
      <slot />
    </div>
  </label>
</template>

<style scoped lang="scss">
input {
  &::placeholder {
    opacity: 0;
  }

  &:placeholder-shown:focus + .label,
  &:not(:placeholder-shown) + .label {
    top: 10px;
    font-size: 12px;
  }
}
</style>
