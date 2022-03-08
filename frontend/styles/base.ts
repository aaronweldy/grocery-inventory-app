import { StyleSheet } from "react-native";

export const BaseStyle = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
  },
  row: {
    flex: 1,
    flexDirection: "row",
    alignItems: "center",
  },
  column: {
    flexDirection: "column",
    alignItems: "center",
  },
});
