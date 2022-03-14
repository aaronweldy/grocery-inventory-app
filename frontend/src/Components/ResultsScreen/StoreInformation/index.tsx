import { View, Text, FlatList } from "react-native";
import { StoreInformationStyle } from "./style";

type StoreInformationProps = {
  name: string;
  address: string;
  missingItems: string[];
};

const renderItem = (item: string) => (
  <Text style={StoreInformationStyle.missingItem}>{item}</Text>
);

export const StoreInformation = ({
  name,
  address,
  missingItems,
}: StoreInformationProps) => (
  <View style={StoreInformationStyle.contentColumn}>
    <Text style={StoreInformationStyle.storeNameText}>{name}</Text>
    <Text style={StoreInformationStyle.storeAddressText}>{address}</Text>
    <Text style={StoreInformationStyle.missingItemHeader}>Missing:</Text>
    <FlatList
      data={missingItems}
      renderItem={({ item }) => (
        <Text style={StoreInformationStyle.missingItem}>{item}</Text>
      )}
      keyExtractor={(item, index) => item + index.toString()}
      horizontal={true}
      ItemSeparatorComponent={() => (
        <Text style={StoreInformationStyle.missingItem}>, </Text>
      )}
    />
  </View>
);
