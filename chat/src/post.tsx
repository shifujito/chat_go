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

  return (
    <div>
      {posts.map((post) => (
        <div key={post.id}>
          <p >{post.name}</p>
          <div>{post.text}</div>
        </div>
      ))}
    </div>
  )
};
