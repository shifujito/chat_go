import React from "react";
import { Post } from "../types";
import { Center, VStack, Text, Stack } from "@chakra-ui/react";
import { PostIcons } from './icons';

type PostContentProps = {
  post: Post;
};

export const DisplayPost: React.FC<PostContentProps> = (post) => {
  return (
    <Center>
       <VStack border={"solid 2px white"} borderTop={"0"} width={"50%"} align={"left"} paddingLeft={"20px"}>
        <Stack minHeight={"75px"}>
          <Text>@ {post.post.name}</Text>
          <Text>{post.post.text}</Text>
        </Stack>
        <PostIcons />
      </VStack>
    </Center>
  );
};
