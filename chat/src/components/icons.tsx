import { useState } from "react";
import { Post } from "../types";
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalFooter,
  HStack,
  Text,
  Button,
} from "@chakra-ui/react";
import { aCl } from "../api-client";

type Props = {
  post: Post;
  loginUserName: string;
  onDelete: (postId: number) => void;
};

export const PostIcons: React.VFC<Props> = ({
  loginUserName,
  post,
  onDelete,
}) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);

  const handleOpenClick: React.MouseEventHandler<HTMLButtonElement> = () => {
    setIsOpen(true);
  };

  const handleCloseClick = () => {
    setIsOpen(false);
  };

  const handleDeletePostAndCloseClick = () => {
    // post methodをよぶ
    // onDelete(post.id);
    setIsOpen(false);
    aCl
      .delete<Post>(`http://127.0.0.1:8080/api/post/delete/${post.id}`)
      .then((res) => {
        onDelete(post.id);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <HStack>
      <Text>いいね</Text>
      <Text>リツイート</Text>
      {loginUserName == post.name ? (
        <Button
          colorScheme="red"
          size="xs"
          variant="outline"
          mr={2}
          onClick={handleOpenClick}
        >
          削除
        </Button>
      ) : null}
      <Modal
        isOpen={isOpen}
        closeOnOverlayClick={true}
        onClose={handleCloseClick}
        size={"xs"}
      >
        <ModalOverlay />
        <ModalContent paddingLeft={"20px"} paddingTop={"20px"}>
          <Text>本当に削除しますか？</Text>
          <ModalFooter>
            <Button
              size={"xs"}
              mr={3}
              colorScheme={"red"}
              onClick={handleDeletePostAndCloseClick}
            >
              削除
            </Button>
            <Button size={"xs"} onClick={handleCloseClick}>
              戻る
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </HStack>
  );
};
