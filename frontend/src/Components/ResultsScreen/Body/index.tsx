import { useNavigation } from "@react-navigation/native";
import { BaseStyle } from "../../../../styles/base";
import { View, Text, TouchableOpacity, Button } from "react-native";
import { ContentRow } from "../ContentRow";
import { SubmitButtonStyle } from "./style";

export const Body = () => {
  const navigation = useNavigation();

  const handlePress = () => {
    setTimeout(() => {
      navigation.navigate("ItemSelection");
    }, 1500);
  };

  return (
    <View style={BaseStyle.body}>
      <ContentRow
        storeName="Kroger"
        storeAddress="500 W Huron St, Ann Arbor, MI, 48109"
        missingItems={["apples", "bananas"]}
        percentageInStock={87}
        storeLogo="https://logos-world.net/wp-content/uploads/2021/09/Kroger-Logo.png"
      />
      <ContentRow
        storeName="Meijer"
        storeAddress="200 S State St, Ann Arbor, MI, 48104"
        missingItems={["bananas", "eggs", "milk"]}
        percentageInStock={76}
        storeLogo="https://www.ywcawcmi.org/wp-content/uploads/2021/09/meijer-logo.jpg"
      />
      <ContentRow
        storeName="Trader Joe's"
        storeAddress="404 South Stadium Rd, Ann Arbor, MI, 48104"
        missingItems={[
          "chicken tenders",
          "chocolate bars",
          "sage",
          "kale",
          "oranges",
        ]}
        percentageInStock={58}
        storeLogo="https://shelbyreport.nyc3.cdn.digitaloceanspaces.com/wp-content/uploads/2020/03/Trader-Joes-logo.jpg"
      />
      <Button title="Back" onPress={handlePress}></Button>
    </View>

  );

};
