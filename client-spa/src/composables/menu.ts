import { ref } from "vue";

const isOpen = ref(false);
const setMenu = (value: boolean) => (isOpen.value = value);
const toggle = () => (isOpen.value = !isOpen.value);

export const useMenu = () => {
  return { isOpen, setMenu, toggle };
};
