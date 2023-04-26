import { IClub } from "types/club";
import { $http } from "~/libs/axios";

export const clubsService = {
  getClubsPerPage: async (page: number, search: string) => {
    const data = (
      await $http.get<{
        clubs: IClub[];
        pages: number;
      }>(`/api/clubs/get-clubs?page=${page}&search=${search}`)
    ).data;

    return data;
  },

  getClubById: async (id: number) => {
    const data = (
      await $http.get<{
        club: IClub;
      }>(`/api/clubs/get-club?id=${id}`)
    ).data.club;

    return data;
  },

  getAllClubsLightWeight: async () => {
    const data = (
      await $http.get<{
        clubs: {
          id: number;
          name: string;
        }[];
      }>(`/api/clubs/get-all-clubs`)
    ).data.clubs;

    return data;
  },
};
