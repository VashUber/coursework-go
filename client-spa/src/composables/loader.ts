import { computed, ref, unref, watch } from "vue";
import { useRoute } from "vue-router";

export const useLoader = <T, ARG extends any[]>(func: (...args: ARG) => Promise<T>, ...args: any) => {
  const route = useRoute();
  const data = ref<T>();
  const isLoading = computed(() => !!data.value);

  const fetch = async () => {
    const res = await func(...args.map((arg: any) => unref(arg)));
    data.value = res;
  };
  fetch();

  watch([() => route.name, ...args], ([newName], [prevName]) => {
    if (newName !== prevName) return;

    fetch();
  });

  return { data, isLoading, refetch: fetch };
};
