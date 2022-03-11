import { StyleSheet } from "react-native";
import { BaseStyle } from "../../../styles/base";

export const ItemEntryStyle = StyleSheet.create({
  entryRow: {
    flex: 1,
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
    //flexGrow: 0,
    marginTop: 40,
    marginBottom: 40,
  },
  textbox: {
    flex: 4,
    height: 40,
    borderWidth: 2,
    margin: 10,
    padding: 2,
  },
});
