import React, { useContext } from "react";
import { View, Text, Button } from "react-native";
import { ListContext } from "../Context/ListContext";
import { ItemInListStyle } from "./style";

type ItemInListProps = {
  itemType: string;
  index: number;
};

export const ItemInList = ({ itemType, index }: ItemInListProps) => {
  const { removeFromList } = useContext(ListContext);

  return (
    <View style={ItemInListStyle.item}>
      <Text style={ItemInListStyle.itemText}>{itemType}</Text>
      <Button
        title="X"
        color="red"
        accessibilityLabel="Click to remove an item from your grocery list"
        onPress={() => removeFromList(index)}
      >
        X
      </Button>
    </View>
  );
};
