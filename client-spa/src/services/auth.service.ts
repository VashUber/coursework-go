import { $http } from "~/libs/axios";
import { IUser } from "types/user";
import { useUser } from "~/composables/user";

export const auth = {
  signin: async (body: { email: string; password: string }) => {
    try {
      await $http.post("/api/auth/signin", body);
    } catch (e) {
      throw e;
    }
  },
  signup: async (body: { email: string; password: string; name: string }) => {
    try {
      await $http.post("/api/auth/signup", body);
    } catch (e) {
      throw e;
    }
  },
  getUserInfo: async () => {
    const { setUser } = useUser();

    try {
      const response = (
        await $http<{
          user: IUser | null;
        }>("/api/auth/get-user-info")
      ).data;

      if (response.user) {
        setUser(response.user);
      }
    } catch (e) {
      throw e;
    }
  },
};
