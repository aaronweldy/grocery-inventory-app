import { StyleSheet } from "react-native";

export const ContentRowStyles = StyleSheet.create({
  contentRow: {
    flexDirection: "row",
    justifyContent: "space-around",
    width: "95%",
    borderWidth: 2,
  },
  logo: {
    flex: 1,
    height: "100%",
  },
  percentageContainer: {
    flex: 1,
    alignItems: "center",
    justifyContent: "center",
  },
  percentage: {
    fontSize: 36,
    color: "#097969",
  },
});
