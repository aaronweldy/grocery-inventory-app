import React, { useContext, useState } from "react";
import { TextInput, TouchableOpacity, View, Text } from "react-native";
import { ListContext } from "../../../Context/ListContext";
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
      <TouchableOpacity onPress={handleAddItem} style={ItemEntryStyle.button}>
        <Text style={ItemEntryStyle.buttonText}>Add Item</Text>
      </TouchableOpacity>
    </View>
  );
};
