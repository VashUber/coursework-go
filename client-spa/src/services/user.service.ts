import { $http } from "~/libs/axios";
import { IProfile } from "types/profile";

export const userService = {
  getProfileInfo: async () => {
    const response = (
      await $http.get<{
        profile: IProfile;
      }>("/api/user/get-profile-info")
    ).data.profile;

    return response;
  },
};
