import { useNavigation } from "@react-navigation/native";
import React from "react";
import { Text, TouchableOpacity, View } from "react-native";
import { SubmitButtonStyle } from "./style";

export const SubmitButton = () => {
  const navigation = useNavigation();

  const handlePress = () => {
    navigation.navigate("ItemSelection");
  };

  return (
    <View style={SubmitButtonStyle.buttonContainer}>
      <TouchableOpacity style={SubmitButtonStyle.button} onPress={handlePress}>
        <Text style={SubmitButtonStyle.buttonText}>Submit!</Text>
      </TouchableOpacity>
    </View>
  );
};
