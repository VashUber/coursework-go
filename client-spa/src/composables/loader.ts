import { computed, ref, unref, watch } from "vue";

export const useLoader = <T, ARG extends any[]>(func: (...args: ARG) => Promise<T>, ...args: any) => {
  const data = ref<T>();
  const isLoading = computed(() => !!data.value);

  const fetch = async () => {
    const res = await func(...args.map((arg: any) => unref(arg)));
    data.value = res;
  };
  fetch();

  watch(args, fetch);

  return { data, isLoading, refetch: fetch };
};
