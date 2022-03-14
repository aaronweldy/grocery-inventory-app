import React from "react";
import { BaseStyle } from "../../../../styles/base";
import { View, Text } from "react-native";
import { HeaderStyles } from "./style";

export const Header = () => (
  <View style={HeaderStyles.header}>
    <Text style={HeaderStyles.headerText}>Results</Text>
  </View>
);
