import { StyleSheet } from "react-native";

export const ItemInListStyle = StyleSheet.create({
  item: {
    flex: 1,
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "flex-start",
    width: 300,
    padding: 5,
    borderWidth: 1,
    borderColor: "#20232a",
  },
  itemText: {
    fontSize: 24,
    fontWeight: "bold",
  }
});
