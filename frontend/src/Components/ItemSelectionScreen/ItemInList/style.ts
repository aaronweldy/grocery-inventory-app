import { StyleSheet } from "react-native";

export const ItemInListStyle = StyleSheet.create({
  item: {
    flex: 1,
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "flex-start",
    width: 300,
    padding: 5,
  },
  itemText: {
    fontSize: 24,
    fontWeight: "bold",
  },
});
