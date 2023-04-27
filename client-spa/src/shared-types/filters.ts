import { TOption } from "./select";

export type TFilters = Record<
  string,
  {
    options: TOption[];
    model: TOption;
  }
>;
