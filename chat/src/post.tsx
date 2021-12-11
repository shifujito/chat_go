import React, { useEffect } from "react";
import { aCl } from './api-client';
import { useState } from 'react';

type Post = {
  id: number
  name: string
  text: string
}

export const Post: React.VFC = () => {
  const [posts, setPosts] = useState<Post[]>([])
  useEffect(() => {
    aCl.get("http://127.0.0.1:8080/api/posts").then((res) => {
      setPosts(res.data.map((val) => ({
        id: val.id,
        name: val.name,
        text: val.text,
      })))
    })}
    ,[])

  console.log(posts)
  // const [posts, setPosts] = useState<Post[]>([])

  // const allPosts = () => {
  //   aCl.get("http://127.0.0.1:8080/api/posts").then((res) => {
  //       // setPosts([...posts, res.data[0]])
  //       console.log(posts)
  //     }
  //   )
  // }
  return <h2>Hello</h2>;
};
