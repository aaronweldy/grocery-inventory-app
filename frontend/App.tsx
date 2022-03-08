import { StatusBar } from "expo-status-bar";
import { View } from "react-native";
import { BaseStyle } from "./styles/base";
import { Header } from "./src/Components/Header";
import { Body } from "./src/Components/Body";
import React from "react";

export default function App() {
  return (
    <View style={BaseStyle.container}>
      <Header />
      <Body />
      <StatusBar style="auto" />
    </View>
  );
}
