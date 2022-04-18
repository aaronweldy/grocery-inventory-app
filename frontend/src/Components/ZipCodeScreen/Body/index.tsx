import React, { useState } from "react";
import { ActivityIndicator, View, TextInput } from "react-native";
import { SubmitButton } from "../SubmitButton";
import { BaseStyle } from "../../../../styles/base";

export const Body = () => {
  const [loading, setLoading] = useState(false);
  return (
    <View style={BaseStyle.body}>
      <TextInput
        style={BaseStyle.input}
        placeholder="Zip Code"
        keyboardType="numeric"
      />
      <SubmitButton setLoading={setLoading} />
      {loading && <ActivityIndicator size="large" style={BaseStyle.loading} />}
    </View>
  );
};
