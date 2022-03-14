import { StyleSheet } from "react-native";

export const StoreInformationStyle = StyleSheet.create({
  contentColumn: {
    flex: 2,
    borderLeftWidth: 1,
    borderRightWidth: 1,
    padding: 3,
  },
  storeNameText: {
    fontSize: 24,
    fontWeight: "bold",
  },
  storeAddressText: {
    fontSize: 18,
    color: "#a9a9a9",
  },
  missingItemHeader: {
    fontSize: 12,
    fontWeight: "bold",
  },
  missingItem: {
    fontSize: 12,
    color: "#c41e3a",
  },
});
