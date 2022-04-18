import { useNavigation } from "@react-navigation/native";
import React from "react";
import { Text, TouchableOpacity, View } from "react-native";
import { BaseStyle } from "../../../../styles/base";
import { SubmitButtonStyle } from "./style";

type SubmitButtonProps = {
  setLoading: React.Dispatch<React.SetStateAction<boolean>>;
};

export const SubmitButton = ({ setLoading }: SubmitButtonProps) => {
  const navigation = useNavigation();

  const handlePress = () => {
    setLoading(true);
    setTimeout(() => {
      setLoading(false);
      navigation.navigate("ItemSelection");
    }, 1500);
  };

  return (
    <View style={SubmitButtonStyle.buttonContainer}>
      <TouchableOpacity style={SubmitButtonStyle.button} onPress={handlePress}>
        <Text style={SubmitButtonStyle.buttonText}>Submit!</Text>
      </TouchableOpacity>
    </View>
  );
};
