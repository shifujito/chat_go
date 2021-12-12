import { Flex, Heading, Spacer, Stack, Text } from "@chakra-ui/react";

type Props = {
  id: number
  name:  string
  isLogined: boolean
}

export const HeaderLayout: React.VFC<Props> = (props) => {
  return (
    <Flex as="nav" bg="teal.500" padding={6} borderBottom={"solid 2px white"} justify="space-between">
      <Heading size={"md"}>投稿一覧</Heading>
      <Spacer />
      <Text>{props.name}</Text>
    </Flex>
  );
};
