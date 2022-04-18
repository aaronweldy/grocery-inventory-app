import { StyleSheet } from "react-native";

export const ContentRowStyles = StyleSheet.create({
  listBody: {
    marginTop: 80,
  },
  contentRow: {
    flexDirection: "row",
    justifyContent: "space-around",
    width: "95%",
    borderWidth: 2,
  },
  logo: {
    flex: 1,
    height: "100%",
    paddingHorizontal: 50,
  },
  percentageContainer: {
    flex: 1,
    alignItems: "center",
    justifyContent: "center",
  },
  percentage: {
    fontSize: 36,
    color: "#097969",
    paddingHorizontal: 40,
  },
});
