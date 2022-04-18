import { RouteProp } from "@react-navigation/native";
import { NativeStackNavigationProp } from "@react-navigation/native-stack";

export type StoreResult = {
  StoreName: string;
  StoreAddress: string;
  MissingItems: string[];
  PercentageInStock: string;
  StoreLogo: string;
};

export type RootStackParamList = {
  ZipCode: undefined;
  ItemSelection: undefined;
  Results: { stores: StoreResult[] };
};

export type ResultsScreenNavigationProp = NativeStackNavigationProp<
  RootStackParamList,
  "Results"
>;

export type ResultsScreenRouteProp = RouteProp<RootStackParamList, "Results">;
