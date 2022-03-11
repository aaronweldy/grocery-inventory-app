import { StyleSheet } from "react-native";

export const ItemListStyle = StyleSheet.create({
  list: {
    flex: 5,
    width: 375,
    maxHeight: 500,
    borderWidth: 1,
    marginBottom: 20,
  },
  title: {
    flex: 1,
    alignItems: "center",
    flexGrow: 0,
    flexBasis: 30,
    marginBottom: 50,
  },
  titleText: {
    fontWeight: "bold",
    fontSize: 24,
  },
  body: {
    flex: 5,
    width: 370,
    alignItems: "center",
  },
});
