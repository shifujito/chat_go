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

  const handleClick: React.MouseEventHandler<HTMLButtonElement> = () => {
    isOpen ? setIsOpen(false) : setIsOpen(true);
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
          onClick={handleClick}
        >
          削除
        </Button>
      ) : null}
      <Modal isOpen={isOpen} onClose={() => {}} size={"xs"}>
        {/* <ModalOverlay /> */}
        <ModalContent paddingLeft={"20px"} paddingTop={"20px"}>
          <Text>本当に削除しますか？</Text>
          <ModalFooter>
            <Button
              size={"xs"}
              mr={3}
              colorScheme={"red"}
              onClick={handleClick}
            >
              削除
            </Button>
            <Button size={"xs"} onClick={handleClick}>
              戻る
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </HStack>
  );
};