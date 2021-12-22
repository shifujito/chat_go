import { Flex, Heading, Spacer, Stack, Link } from "@chakra-ui/react";
import { useHistory } from "react-router-dom";

type Props = {
  id: number;
  name: string;
  isLogined: boolean;
};

export const HeaderLayout: React.VFC<Props> = (props) => {
  const history = useHistory();
  const handleProfileClick = () => {
    history.push("/profile");
  };

  return (
    <Flex as="nav" bg="teal.500" padding={6} justify="space-between">
      <Heading size={"md"}>投稿一覧</Heading>
      <Spacer />
      <Link onClick={handleProfileClick}>{props.name}</Link>
    </Flex>
  );
};
