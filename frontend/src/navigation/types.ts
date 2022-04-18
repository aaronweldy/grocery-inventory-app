export type RootStackNavigator = {
  ItemSelection: undefined;
  Results: undefined;
  ZipCode: undefined;
};

declare global {
  namespace ReactNavigation {
    interface RootParamList extends RootStackNavigator {}
  }
}
