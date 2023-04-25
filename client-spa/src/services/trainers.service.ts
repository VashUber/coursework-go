import { $http } from "~/libs/axios";
import { ITrainerSimple } from "types/trainers";

export const trainersService = {
  getTrainersPerPage: async (page: number) => {
    const resp = await $http.get<{
      trainers: ITrainerSimple[];
      pages: number;
    }>(`/api/trainers/get-trainers?page=${page}`);

    return resp.data;
  },
};
