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
    flex: 3,
    height: 40,
    borderWidth: 2,
    margin: 10,
    padding: 2,
  },
  button: {
    flex: 1,
    backgroundColor: "#24a0ed",
    alignItems: "center",
    justifyContent: "center",
    height: 40,
    padding: 3,
    marginRight: 5,
  },
  buttonText: {
    color: "white",
  },
});
