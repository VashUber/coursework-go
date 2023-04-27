import { IEquipment } from "./equipment";
export interface IClubSimple {
  ID: number;
  name: string;
  address: {
    street: string;
    home: string;
  };
}

export interface IClub extends IClubSimple {
  thumb: string;
  image: string;
  info: string;
  schedule: string[];
  equipment: IEquipment[];
}
