import React, { useState } from "react";
import { ActivityIndicator, View } from "react-native";
import { ItemEntry } from "../ItemEntry";
import { ItemList } from "../ItemList";
import { SubmitButton } from "../SubmitButton";
import { BaseStyle } from "../../../../styles/base";

export const Body = () => {
  const [loading, setLoading] = useState(false);
  return (
    <View style={BaseStyle.body}>
      <ItemEntry />
      <ItemList />
      <SubmitButton setLoading={setLoading} />
      {loading && <ActivityIndicator size="large" style={BaseStyle.loading} />}
    </View>
  );
};
