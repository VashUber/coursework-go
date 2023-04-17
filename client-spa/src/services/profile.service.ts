import { $http } from "~/libs/axios";
import { IUserProfile } from "types/profile";

export const profile = {
  getProfileInfo: async () => {
    const response = (await $http.get<IUserProfile>("/api/profile/get-info"))
      .data;

    return response;
  },
};
