import React from "react";
import {
  Textarea,
  Center,
  Stack,
  Button,
  Divider,
  Flex,
  Spacer,
} from "@chakra-ui/react";

export const CreatePost: React.VFC = () => {
  return (
    <Center>
      <Stack borderTop={"0"} width={"50%"} border={"solid 3px skyblue"}>
        <Textarea minHeight={"125px"} placeholder="投稿内容" border={"none"} />
        <Divider />
        <Flex padding={"1"}>
          <Spacer />
          <Button width={"100px"} colorScheme={"blue"}>
            投稿
          </Button>
        </Flex>
      </Stack>
    </Center>
  );
};
