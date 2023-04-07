import { $http } from "~/libs/axios";

export const auth = {
  signin: async () => {},
  signup: async () => {},
  getUserInfo: async () => {
    await $http("/api");
  },
};
