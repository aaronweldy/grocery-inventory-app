import { StyleSheet } from "react-native";

export const ItemListStyle = StyleSheet.create({
  title: {
    flex: 1,
    alignItems: "center",
    flexBasis: 30,
  },
  titleText: {
    fontWeight: "bold",
    fontSize: 24,
  },
  body: {
    height: 400,
    minWidth: 300,
    borderWidth: 2,
  },
});
