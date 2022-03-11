import React, { useContext } from "react";
import { FlatList, Text, View } from "react-native";
import { BaseStyle } from "../../../styles/base";
import { ListContext } from "../Context/ListContext";
import { ItemInList } from "../ItemInList";
import { ItemListStyle } from "./style";

export const ItemList = () => {
  const { list } = useContext(ListContext);

  return (
    <View style={ItemListStyle.list}>
      <View style={ItemListStyle.title}>
        <Text style={ItemListStyle.titleText}>Grocery List</Text>
      </View>
      <View style={ItemListStyle.body}>
        <FlatList
          data={list}
          scrollIndicatorInsets={{
            top: 0,
            left: 0,
            right: 50,
            bottom: 0,
          }}
          renderItem={({ item, index }) => (
            <ItemInList index={index} itemType={item} />
          )}
          keyExtractor={(item, index) => item + index.toString()}
        />
      </View>
    </View>
  );
};
