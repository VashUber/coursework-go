import { IClubSimple } from "./club";

export interface ITrainerSimple {
  id: number;
  thumb: string;
  name: string;
  services: string[];
}

export interface ITrainer extends ITrainerSimple {
  clubs: IClubSimple[];
}
