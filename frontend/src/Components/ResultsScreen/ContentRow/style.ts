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
    flexGrow: 2,
    flexWrap: "nowrap",
    alignItems: "center",
    justifyContent: "center",
  },
  percentage: {
    fontSize: 32,
    color: "#097969",
    paddingHorizontal: 5,
  },
});
