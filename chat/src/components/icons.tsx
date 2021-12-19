import { useState } from "react";
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  HStack,
  Text,
  Button,
} from "@chakra-ui/react";

type Props = {
  postName: string;
  loginUserName: string;
};

export const PostIcons: React.VFC<Props> = ({ loginUserName, postName }) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);

  const handleOpenClick: React.MouseEventHandler<HTMLButtonElement> = () => {
    setIsOpen(true);
  };

  const handleCloseClick = () => {
    setIsOpen(false);
  };

  return (
    <HStack>
      <Text>いいね</Text>
      <Text>リツイート</Text>
      {loginUserName == postName ? (
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
              onClick={handleCloseClick}
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
