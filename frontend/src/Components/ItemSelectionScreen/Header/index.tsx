import React from "react";
import { Text, View } from "react-native";
import { Logo } from "../Logo";
import { BaseStyle } from "../../../../styles/base";
import { HeaderStyles } from "./style";

export const Header = () => (
  <View style={HeaderStyles.header}>
    <Logo />
    <View style={BaseStyle.column}>
      <Text style={HeaderStyles.appTitle}>Stockify</Text>
    </View>
  </View>
);
