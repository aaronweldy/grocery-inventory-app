import { StatusBar } from "expo-status-bar";
import React from "react";
import { SafeAreaView } from "react-native";
import { BaseStyle } from "../../../styles/base";

export const ZipCodeScreen = () => (
  <SafeAreaView style={BaseStyle.container}>
    <StatusBar style="auto" />
  </SafeAreaView>
);
