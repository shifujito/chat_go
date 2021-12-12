import React, { useEffect } from "react";
import { aCl } from "./api-client";
import { useState } from "react";
import { DisplayPost } from "./components/postContent";
import { Post } from "./types";
import { HeaderLayout } from "./components/header";
import { Stack, Box } from '@chakra-ui/react';
export const Posts: React.VFC = () => {
  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => {
    aCl.get("http://127.0.0.1:8080/api/posts").then((res) => {
      setPosts(
        res.data.map((val) => ({
          id: val.id,
          name: val.name,
          text: val.text,
        }))
      );
    });
  }, []);

  return (
    <>
      <HeaderLayout />
      <Stack spacing="0" height={""}>
        {posts.map((post) => (
          <Box key={post.id}>
            <DisplayPost post={post} />
          </Box>
        ))}
      </Stack>
    </>
  );
};
