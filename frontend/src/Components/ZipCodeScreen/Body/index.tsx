import React, { useContext, useState } from "react";
import { View, TextInput } from "react-native";
import { SubmitButton } from "../SubmitButton";
import { BaseStyle } from "../../../../styles/base";
import { ZipCodeContext } from "../../../Context/ZipCodeContext";

export const Body = () => {
  const { zipCode, setZipCode } = useContext(ZipCodeContext);
  return (
    <View style={BaseStyle.body}>
      <TextInput
        style={BaseStyle.input}
        placeholder="Zip Code"
        keyboardType="numeric"
        onChangeText={(text) => setZipCode(text)}
      />
      <SubmitButton />
    </View>
  );
};
