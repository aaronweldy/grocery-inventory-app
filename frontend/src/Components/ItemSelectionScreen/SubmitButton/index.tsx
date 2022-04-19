import { useNavigation } from "@react-navigation/native";
import React, { useContext } from "react";
import { Text, TouchableOpacity, View } from "react-native";
import { ListContext } from "../../../Context/ListContext";
import { ResultsScreenNavigationProp, StoreResult } from "../../../types";
import { SubmitButtonStyle } from "./style";

type SubmitButtonProps = {
  zip: string;
  setLoading: React.Dispatch<React.SetStateAction<boolean>>;
};

export const SubmitButton = ({ zip, setLoading }: SubmitButtonProps) => {
  const navigation = useNavigation<ResultsScreenNavigationProp>();
  const { list } = useContext(ListContext);

  const handlePress = async () => {
    setLoading(true);
    const url = "https://git.heroku.comt/quiet-oasis-59230.git/list";
    const body = {
      products: list,
      zip,
    };
    const headers = {
      mode: "no-cors",
      method: "POST",
      "Content-Type": "application/json",
      body: JSON.stringify(body),
    };
    const data = await fetch(url, headers);
    const response = (await data.json()) as { Stores: StoreResult[] };
    setLoading(false);
    navigation.navigate("Results", { stores: response.Stores });
  };

  return (
    <View style={SubmitButtonStyle.buttonContainer}>
      <TouchableOpacity style={SubmitButtonStyle.button} onPress={handlePress}>
        <Text style={SubmitButtonStyle.buttonText}>Submit!</Text>
      </TouchableOpacity>
    </View>
  );
};
