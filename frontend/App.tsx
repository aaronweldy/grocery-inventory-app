import React, { useCallback, useState } from "react";
import { ItemSelectionScreen } from "./src/Components/ItemSelectionScreen";
import { ListContext } from "./src/Context/ListContext";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { RootStackNavigator } from "./src/navigation/types";
import { ResultsScreen } from "./src/Components/ResultsScreen";

export default function App() {
  const [list, setList] = useState<string[]>([]);
  const Stack = createNativeStackNavigator<RootStackNavigator>();

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
    <NavigationContainer>
      <ListContext.Provider value={{ list, addToList, removeFromList }}>
        <Stack.Navigator screenOptions={{ headerShown: false }}>
          <Stack.Screen name="ItemSelection" component={ItemSelectionScreen} />
          <Stack.Screen name="Results" component={ResultsScreen} />
        </Stack.Navigator>
      </ListContext.Provider>
    </NavigationContainer>
  );
}
