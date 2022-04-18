import React from "react";
import { View, Text, Image } from "react-native";
import { StoreInformation } from "../StoreInformation";
import { ContentRowStyles } from "./style";

type ContentRowProps = {
  storeName: string;
  storeAddress: string;
  missingItems: string[];
  percentageInStock: number;
  storeLogo: string;
};

export const ContentRow = ({
  storeName,
  storeAddress,
  missingItems,
  percentageInStock,
  storeLogo,
}: ContentRowProps) => {
  return (
    <View style={ContentRowStyles.contentRow}>
      <Image
        source={{ uri: storeLogo }}
        style={ContentRowStyles.logo}
        resizeMode="center"
      />
      <StoreInformation
        name={storeName}
        address={storeAddress}
        missingItems={missingItems}
      />
      <View style={ContentRowStyles.percentageContainer}>
        <Text style={ContentRowStyles.percentage}>
          {percentageInStock.toString() + "%"}
        </Text>
      </View>
    </View>
  );
};
