import { useNavigation, useRoute } from "@react-navigation/native";
import { BaseStyle } from "../../../../styles/base";
import { View, Button, FlatList } from "react-native";
import { ContentRow } from "../ContentRow";
import { ResultsScreenRouteProp } from "../../../types";
import React from "react";
import { ContentRowStyles } from "../ContentRow/style";

const storeDefaultImages: Record<string, string> = {
  kroger:
    "https://upload.wikimedia.org/wikipedia/commons/thumb/9/92/Kroger_Logo_11-6-19.svg/768px-Kroger_Logo_11-6-19.svg.png?20191106175859",
  shell:
    "https://upload.wikimedia.org/wikipedia/en/thumb/e/e8/Shell_logo.svg/1200px-Shell_logo.svg.png",
};

export const Body = () => {
  const navigation = useNavigation();
  const route = useRoute<ResultsScreenRouteProp>();
  const sortedStores = route.params.stores.sort((s1, s2) => {
    return parseInt(s2.PercentageInStock) - parseInt(s1.PercentageInStock);
  });
  const handlePress = () => {
    navigation.navigate("ItemSelection");
  };

  return (
    <View style={BaseStyle.body}>
      <FlatList
        style={ContentRowStyles.listBody}
        data={sortedStores}
        keyExtractor={(item) => item.StoreName + item.StoreAddress}
        renderItem={({ item }) => {
          console.log(item);
          return (
            <ContentRow
              storeAddress={item.StoreAddress}
              storeName={item.StoreName}
              storeLogo={
                item.StoreLogo ||
                storeDefaultImages[item.StoreName.split(" ")[0].toLowerCase()]
              }
              percentageInStock={parseInt(item.PercentageInStock)}
              missingItems={item.MissingItems}
            />
          );
        }}
      />
      <Button title="Back" onPress={handlePress}></Button>
    </View>
  );
};
