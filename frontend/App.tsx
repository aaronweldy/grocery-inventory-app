import React, { useCallback, useState } from "react";
import { ItemSelectionScreen } from "./src/Components/ItemSelectionScreen";
import { ListContext } from "./src/Context/ListContext";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { RootStackNavigator } from "./src/navigation/types";
import { ResultsScreen } from "./src/Components/ResultsScreen";
import { ZipCodeScreen } from "./src/Components/ZipCodeScreen";
import { RootStackParamList } from "./src/types";
import { ZipCodeContext } from "./src/Context/ZipCodeContext";

export default function App() {
  const [list, setList] = useState<string[]>([]);
  const [zipCode, setZipCode] = useState<string>("");
  const Stack = createNativeStackNavigator<RootStackParamList>();

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
      <ZipCodeContext.Provider value={{ zipCode, setZipCode }}>
        <ListContext.Provider value={{ list, addToList, removeFromList }}>
          <Stack.Navigator screenOptions={{ headerShown: false }}>
            <Stack.Screen name="ZipCode" component={ZipCodeScreen} />
            <Stack.Screen
              name="ItemSelection"
              component={ItemSelectionScreen}
            />
            <Stack.Screen name="Results" component={ResultsScreen} />
          </Stack.Navigator>
        </ListContext.Provider>
      </ZipCodeContext.Provider>
    </NavigationContainer>
  );
}
