import { StatusBar } from "expo-status-bar";
import { SafeAreaView } from "react-native";
import { BaseStyle } from "./styles/base";
import { Header } from "./src/Components/Header";
import { Body } from "./src/Components/Body";
import React, { useCallback, useState } from "react";
import { ListContext } from "./src/Components/Context/ListContext";

export default function App() {
  const [list, setList] = useState<string[]>([]);

  const addToList = useCallback(
    (item: string) => {
      setList(list.concat(item));
    },
    [list, setList]
  );

  const removeFromList = useCallback(
    (idx: number) => {
      const updateList = [...list];
      updateList.splice(idx, 1);
      setList(updateList);
    },
    [list, setList]
  );

  return (
    <ListContext.Provider value={{ list, addToList, removeFromList }}>
      <SafeAreaView style={BaseStyle.container}>
        <Header />
        <Body />
        <StatusBar style="auto" />
      </SafeAreaView>
    </ListContext.Provider>
  );
}
