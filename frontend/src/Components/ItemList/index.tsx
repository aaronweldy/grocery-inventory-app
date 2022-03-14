import React, { useContext } from "react";
import { FlatList, Text, View } from "react-native";
import { ListContext } from "../Context/ListContext";
import { ItemInList } from "../ItemInList";

import { ItemListStyle } from "./style";

export const ItemList = () => {
  const { list } = useContext(ListContext);

  return (
    <View style={ItemListStyle.body}>
      <FlatList
        ListHeaderComponent={
          <Text style={ItemListStyle.titleText}>Grocery List</Text>
        }
        ListHeaderComponentStyle={ItemListStyle.title}
        data={list}
        renderItem={({ item, index }) => (
          <ItemInList index={index} itemType={item} />
        )}
        keyExtractor={(item, index) => item + index.toString()}
      />
    </View>
  );
};
