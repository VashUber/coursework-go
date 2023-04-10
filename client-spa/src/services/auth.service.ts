import { $http } from "~/libs/axios";

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
    try {
      const data = (await $http("/api/auth/get-user-info")).data;
    } catch (e) {
      throw e;
    }
  },
};
