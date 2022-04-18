import { createContext } from "react";

type ZipCodeContextType = {
  zipCode: string;
  setZipCode: (zipCode: string) => void;
};

export const ZipCodeContext = createContext<ZipCodeContextType>({
  zipCode: "",
  setZipCode: (zipCode: string) => {},
});
