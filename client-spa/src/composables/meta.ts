import { ref } from "vue";

const head = ref({
  title: "",
  description: "",
});

const setHead = (newMeta: { title: string; description: string }) => {
  head.value = newMeta;
};

export const useMeta = () => {
  return { head, setHead };
};
