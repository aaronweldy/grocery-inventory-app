import { createContext } from "react";

type ListContextType = {
  list: string[];
  addToList: (i: string) => void;
  removeFromList: (idx: number) => void;
};

export const ListContext = createContext<ListContextType>({
  list: [],
  addToList: (item: string) => {},
  removeFromList: (idx: number) => {},
});
