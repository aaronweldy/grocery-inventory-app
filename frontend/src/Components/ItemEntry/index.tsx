import React, { useContext, useState } from "react";
import { Button, TextInput, View } from "react-native";
import { ListContext } from "../Context/ListContext";
import { ItemEntryStyle } from "./style";

export const ItemEntry = () => {
  const [text, setText] = useState("");
  const { addToList } = useContext(ListContext);

  const onChange = (text: string) => {
    setText(text);
  };

  const handleAddItem = () => {
    addToList(text);
    setText("");
  };

  return (
    <View style={ItemEntryStyle.entryRow}>
      <TextInput
        style={ItemEntryStyle.textbox}
        placeholder="Enter items..."
        value={text}
        onChangeText={onChange}
      />
      <Button
        title="Add Item"
        accessibilityLabel="Click here to add an item to your grocery list!"
        onPress={handleAddItem}
      />
    </View>
  );
};
