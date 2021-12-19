import React from "react";
import { Post } from "../types";
import { Center, VStack, Text, Stack } from "@chakra-ui/react";
import { PostIcons } from './icons';

type Props = {
  loginUserName: string
  post: Post;
};

export const DisplayPost: React.VFC<Props> = ({post, loginUserName}) => {
  return (
    <Center>
       <VStack border={"solid 2px white"} borderTop={"0"} width={"50%"} align={"left"} paddingLeft={"20px"}>
        <Stack minHeight={"75px"}>
          <Text>@ {post.name}</Text>
          <Text>{post.text}</Text>
        </Stack>
        <PostIcons postName={post.name} loginUserName={loginUserName}/>
      </VStack>
    </Center>
  );
};
