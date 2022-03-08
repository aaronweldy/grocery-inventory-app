import { BaseStyle } from "../../../styles/base";
import React from "react";
import { View, Text } from "react-native";
import { BodyStyles } from "./style";

export const Body = () => (
  <View style={BodyStyles.body}>
    <Text>Test1</Text>
    <Text>Test2</Text>
  </View>
);
