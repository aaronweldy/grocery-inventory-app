import React from "react";
import { Button, View } from "react-native";
import { BaseStyle } from "../../../styles/base";
import { SubmitButtonStyle } from "./style";

export const SubmitButton = () => {
  return (
    <View style={SubmitButtonStyle.button}>
      <Button color="green" title="Submit!" onPress={() => {}} />
    </View>
  );
};
