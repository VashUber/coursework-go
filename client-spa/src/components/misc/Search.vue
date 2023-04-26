<script setup lang="ts">
import { ref, toRefs } from "vue";
import { useRouter } from "vue-router";
import Search from "../icons/Search.vue";

const router = useRouter();

const props = defineProps<{
  modelValue: string;
}>();
const { modelValue } = toRefs(props);
const emits = defineEmits<{
  (e: "update:modelValue", value: string): void;
}>();
const isFocused = ref(false);
let timeoutId: NodeJS.Timeout;

const onInput = (e: Event) => {
  clearTimeout(timeoutId);
  const value = (e.target as HTMLInputElement).value;

  emits("update:modelValue", value);
  timeoutId = setTimeout(() => {
    router.push({
      query: {
        search: value,
      },
    });
  }, 400);
};
</script>

<template>
  <div
    class="flex items-center px-4 py-2 gap-2 border border-gray-400 text-gray-950 bg-slate-50 rounded-md h-10 transition-colors"
    :class="{
      'border-orange-400': isFocused,
    }"
  >
    <input
      type="text"
      @input="onInput"
      @focus="isFocused = true"
      @blur="isFocused = false"
      :value="modelValue"
      class="outline-none bg-slate-50"
      placeholder="Search"
    />
    <Search class="text-gray-400" />
  </div>
</template>
