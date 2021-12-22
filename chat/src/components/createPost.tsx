import React from "react";
import { useState } from "react";
import { useHistory } from "react-router-dom";
import { aCl } from "../api-client";
import { loginInfo } from "../types";
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

type Porps = {
  loginUser: loginInfo;
  onCreate: () => void;
};

export const CreatePost: React.VFC<Porps> = ({ loginUser, onCreate }) => {
  const history = useHistory();
  const [postContent, setPostContent] = useState<string>("");

  const handleInput = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setPostContent(e.target.value);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    aCl
      .post("http://localhost:8080/api/post/create", {
        userId: loginUser.id,
        content: postContent,
      })
      .then((res) => {
        onCreate();
        setPostContent("");
      });
  };

  return (
    <Center>
      <Stack borderTop={"0"} width={"50%"} border={"solid 3px skyblue"}>
        <form onSubmit={handleSubmit}>
          <FormControl isRequired>
            <Textarea
              minHeight={"125px"}
              placeholder="What are you doing now?"
              border={"none"}
              onChange={handleInput}
              value={postContent}
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
