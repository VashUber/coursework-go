import { IClub } from "types/club";
import { $http } from "~/libs/axios";

export const clubsService = {
  getClubsPerPage: async (page: number) => {
    const data = (
      await $http.get<{
        clubs: IClub[];
        pages: number;
      }>(`/api/clubs/get-clubs?page=${page}`)
    ).data;

    return data;
  },
};
