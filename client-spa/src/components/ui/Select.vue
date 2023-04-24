<template>
  <Listbox>
    <div class="relative mt-1">
      <ListboxButton
        class="min-w-[250px] outline-none px-4 pb-2 pt-4 border-gray-400 border-b rounded-tl-[5px] rounded-tr-[5px] focus:border-orange-400 w-full text-gray-950 bg-slate-50 flex justify-start"
      >
        {{ modelValue.label }}
      </ListboxButton>

      <transition
        leave-active-class="transition duration-100 ease-in"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
      >
        <ListboxOptions
          class="absolute mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm z-20"
        >
          <ListboxOption
            v-slot="{ active }"
            v-for="option in options"
            :key="option.label"
            :value="option"
            as="template"
          >
            <li
              :class="[
                active ? 'text-white bg-orange-400' : 'text-gray-950',
                'relative cursor-default select-none py-2 pl-4 pr-4',
              ]"
              @click="onUpdate(option)"
            >
              {{ option.label }}
            </li>
          </ListboxOption>
        </ListboxOptions>
      </transition>
    </div>
  </Listbox>
</template>

<script setup lang="ts">
import { toRefs } from "vue";
import { Listbox, ListboxButton, ListboxOptions, ListboxOption } from "@headlessui/vue";
import { TOption } from "types/select";

const props = defineProps<{
  modelValue: TOption;
  options: TOption[];
}>();
const emits = defineEmits<{
  (e: "update:modelValue", value: TOption): void;
}>();
const { modelValue, options } = toRefs(props);

const onUpdate = (val: TOption) => {
  emits("update:modelValue", val);
};
</script>
