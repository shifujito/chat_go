import React, { useEffect } from "react";
import { aCl } from "./api-client";
import { useState } from "react";
import { DisplayPost } from "./components/postContent";
import { Post } from "./types";
import { HeaderLayout } from "./components/header";
import { Stack, Box } from "@chakra-ui/react";
import { useRecoilValue } from "recoil";
import { titleSelector, singInUserState } from "./atom";
import { CreatePost } from "./components/createPost";

export const Posts: React.VFC = () => {
  const [posts, setPosts] = useState<Post[]>([]);
  const sample = useRecoilValue(titleSelector);
  const singInUser = useRecoilValue(singInUserState);

  useEffect(() => {
    aCl.get<Post[]>("http://127.0.0.1:8080/api/posts").then((res) => {
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
      <HeaderLayout
        id={singInUser.id}
        name={singInUser.name}
        isLogined={singInUser.isLogined}
      />
      <Stack spacing="0">
        <CreatePost />
        {posts.map((post) => (
          <Box key={post.id}>
            <DisplayPost post={post} loginUserName={singInUser.name} />
          </Box>
        ))}
      </Stack>
    </>
  );
};
