import React from "react";
import { Text, View } from "react-native";
import { BaseStyle } from "../../../../styles/base";
import { HeaderStyles } from "./style";

export const Header = () => (
  <View style={HeaderStyles.header}>
    <View style={BaseStyle.column}>
      <Text style={HeaderStyles.appTitle}>Stockify</Text>
    </View>
  </View>
);
