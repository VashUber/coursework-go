import { ref } from "vue";

const size = ref({
  isMd: false,
  isSm: false,
});

const onResize = () => {
  const width = window.innerWidth;

  size.value.isMd = width < 1024;
  size.value.isSm = width < 768;
};

export const useSize = () => {
  onResize();

  return { size, onResize };
};
