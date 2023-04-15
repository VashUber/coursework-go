import { $http } from "~/libs/axios";
import { IUser } from "types/user";

export const auth = {
  signin: async (body: { email: string; password: string }) => {
    try {
      return (await $http.post("/api/auth/signin", body)).data;
    } catch (e) {
      throw e;
    }
  },
  signup: async (body: {
    email: string;
    password: string;
    name: string;
    birthday: number;
  }) => {
    try {
      return (await $http.post("/api/auth/signup", body)).data;
    } catch (e) {
      throw e;
    }
  },
  signout: async () => {
    try {
      return (await $http.get("/api/auth/signout")).data;
    } catch (e) {
      throw e;
    }
  },
  getUserInfo: async () => {
    try {
      const response = (
        await $http<{
          user: IUser | null;
        }>("/api/auth/get-user-info")
      ).data.user;

      return response;
    } catch (e) {
      throw e;
    }
  },
};
