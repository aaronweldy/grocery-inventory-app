import React from "react";
import { View, Image } from "react-native";
import { BaseStyle } from "../../../../styles/base";
import { LogoStyle } from "./style";

export const Logo = () => (
  <View style={BaseStyle.column}>
    <Image
      style={LogoStyle.thumbnail}
      source={require("../../../../assets/grocery.png")}
    />
  </View>
);
