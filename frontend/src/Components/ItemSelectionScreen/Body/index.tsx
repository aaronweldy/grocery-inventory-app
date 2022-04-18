import React, { useContext, useState } from "react";
import { ActivityIndicator, View } from "react-native";
import { ItemEntry } from "../ItemEntry";
import { ItemList } from "../ItemList";
import { SubmitButton } from "../SubmitButton";
import { BaseStyle } from "../../../../styles/base";
import { ZipCodeContext } from "../../../Context/ZipCodeContext";

export const Body = () => {
  const [loading, setLoading] = useState(false);
  const { zipCode } = useContext(ZipCodeContext);
  return (
    <View style={BaseStyle.body}>
      <ItemEntry />
      <ItemList />
      <SubmitButton zip={zipCode} setLoading={setLoading} />
      {loading && <ActivityIndicator size="large" style={BaseStyle.loading} />}
    </View>
  );
};
