import { $http } from "~/libs/axios";

export const auth = {
  signin: async () => {},
  signup: async () => {},
  getUserInfo: async () => {
    try {
      const data = (await $http("/api/auth/get-user-info")).data;
    } catch (e) {
      throw e;
    }
  },
};
