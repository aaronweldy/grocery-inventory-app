export type RootStackNavigator = {
  ItemSelection: undefined;
  Results: undefined;
};

declare global {
  namespace ReactNavigation {
    interface RootParamList extends RootStackNavigator {}
  }
}
