import React from "react";
import { StatusBar } from "expo-status-bar";
import { SafeAreaView } from "react-native";
import { BaseStyle } from "../../../styles/base";
import { Body } from "./Body";
import { Header } from "./Header";

export const ItemSelectionScreen = () => (
  <SafeAreaView style={BaseStyle.container}>
    <Header />
    <Body />
    <StatusBar style="auto" />
  </SafeAreaView>
);
