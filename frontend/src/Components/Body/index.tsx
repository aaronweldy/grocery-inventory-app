import React from "react";
import { View } from "react-native";
import { BodyStyles } from "./style";
import { ItemEntry } from "../ItemEntry";
import { ItemList } from "../ItemList";
import { SubmitButton } from "../SubmitButton";

export const Body = () => (
  <View style={BodyStyles.body}>
    <ItemEntry />
    <ItemList />
    <SubmitButton />
  </View>
);
