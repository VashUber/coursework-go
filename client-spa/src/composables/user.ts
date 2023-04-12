import { ref } from "vue";
import type { IUser } from "types/user";

const user = ref<IUser | null>(null);
const setUser = (value: IUser) => {
  user.value = value;
};

export const useUser = () => {
  return { user, setUser };
};
