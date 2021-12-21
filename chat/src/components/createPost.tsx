import React from "react";
import { useState } from "react";
import { aCl } from "../api-client";
import {
  Textarea,
  Center,
  Stack,
  Button,
  Divider,
  Flex,
  Spacer,
  FormControl,
} from "@chakra-ui/react";

export const CreatePost: React.VFC = () => {
  const [postContent, setPostContent] = useState<string>("");

  const handleInput = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setPostContent(e.target.value);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    aCl.post("http://127.0.0.1:8080/api/cratepost", {
      // userId　1 => kimura
      userId: 1,
      content: postContent,
    });
  };

  return (
    <Center>
      <Stack borderTop={"0"} width={"50%"} border={"solid 3px skyblue"}>
        {/* <form onSubmit={hadleSubmit}> */}
        <form>
          <FormControl isRequired>
            <Textarea
              minHeight={"125px"}
              placeholder="投稿内容"
              border={"none"}
              onChange={handleInput}
            />
          </FormControl>
          <Divider />
          <Flex padding={"1"}>
            <Spacer />
            <Button width={"100px"} colorScheme={"blue"} type="submit">
              投稿
            </Button>
          </Flex>
        </form>
      </Stack>
    </Center>
  );
};
