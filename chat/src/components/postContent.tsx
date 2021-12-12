import React from "react";
import { Post } from "../types";
import { Center, VStack, Text } from "@chakra-ui/react";

type PostContentProps = {
  post: Post;
};

export const DisplayPost: React.FC<PostContentProps> = (post) => {
  return (
    <Center>
       <VStack border={"solid 2px white"} borderTop={"0"} width={"50%"} align={"left"} paddingLeft={"20px"}>
        <Text>@ {post.post.name}</Text>
        <Text>{post.post.text}</Text>
      </VStack>
    </Center>
  );
};
