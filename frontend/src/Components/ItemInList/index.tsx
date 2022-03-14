import React, { useContext, useState } from "react";
import { View, Text, Button } from "react-native";
import { ListContext } from "../Context/ListContext";
import { ItemInListStyle } from "./style";

type ItemInListProps = {
  itemType: string;
  index: number;
};

export const ItemInList = ({ itemType, index }: ItemInListProps) => {
  const { removeFromList } = useContext(ListContext);
  let [quantity, setQuantity] = useState(1);

  function incrementQuantity() {
    quantity = quantity + 1;
    setQuantity(quantity);
  }
  function decrementQuantity(){
    quantity = quantity - 1;
    setQuantity(quantity);
  }

  return (
    <View style={ItemInListStyle.item}>
      <Text style={ItemInListStyle.itemText}>{itemType}</Text>
      <Button
        title="+"
        color="green"
        onPress={incrementQuantity}
      >+</Button>
      <Text style={ItemInListStyle.itemText}>{quantity}</Text>
      <Button
        title="-"
        color="red"
        onPress={decrementQuantity}
      >-</Button>
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
