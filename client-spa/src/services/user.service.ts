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

  updateProfileInfo: async (body: { email: string; password: string; name: string; birthday: string }) => {
    const response = (await $http.post("/api/user/update-profile-info", body)).data;

    return response;
  },

  uploadAvatar: async (body: FormData) => {
    const response = (await $http.post("/api/user/upload-avatar", body)).data;

    return response;
  },
};
