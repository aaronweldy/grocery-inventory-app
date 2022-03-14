import React from "react";
import { Text, TouchableOpacity, View } from "react-native";
import { BaseStyle } from "../../../styles/base";
import { SubmitButtonStyle } from "./style";

export const SubmitButton = () => {
  return (
    <View style={SubmitButtonStyle.buttonContainer}>
      <TouchableOpacity style={SubmitButtonStyle.button} onPress={() => {}}>
        <Text style={SubmitButtonStyle.buttonText}>Submit!</Text>
      </TouchableOpacity>
    </View>
  );
};
