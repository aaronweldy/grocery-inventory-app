import { StatusBar } from "expo-status-bar";
import React from "react";
import { SafeAreaView, TextInput } from "react-native";
import { BaseStyle } from "../../../styles/base";
import { Body } from "./Body";
import { Header } from "./Header";


export const ZipCodeScreen = () => (
  <SafeAreaView style={BaseStyle.container}>
    <Header />
    <Body />
    <StatusBar style="auto" />
  </SafeAreaView>
);
