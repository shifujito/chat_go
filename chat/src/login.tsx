import { useState } from "react";
import { useHistory, Link as RouterLink } from "react-router-dom";
import { apiClient, aCl } from "./api-client";
import { useRecoilState } from "recoil";
import { singInUserState } from "./atom";
import {
  Button,
  Input,
  Heading,
  Link,
  Box,
  Center,
  Stack,
  FormControl,
  FormLabel,
  Text,
} from "@chakra-ui/react";

function Login() {
  const history = useHistory();

  const [inputname, setInputName] = useState<string>("");
  const [inputpass, setInputPass] = useState<string>("");

  const [showModal, setShowModal] = useState<boolean>(false);

  const handleInputName = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputName(e.target.value);
  };

  const handleInputPass = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputPass(e.target.value);
  };

  const [signInUser, setSignInUser] = useRecoilState(singInUserState);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    aCl
      .post("http://127.0.0.1:8080/login", {
        name: inputname,
        password: inputpass,
      })
      .then((res) => {
        setSignInUser({
          id: res.data.id,
          name: res.data.name,
          isLogined: true,
        });
        history.push("/posts");
      })
      .catch((err) => {
        // 画面遷移しない。
        setShowModal(true);
      });
  };

  return (
    <>
      <Box w="100%" p={8}>
        <Center w="100%">
          <Stack spacing="5">
            <Heading as="h3" textAlign={"center"}>
              ログインページ
            </Heading>
            {showModal ? (
              <Text textAlign={"center"} color="red">
                ユーザー名またはパスワードが一致しません
              </Text>
            ) : null}
            <Center bg="#1A202C" color="white" padding="20px">
              <form onSubmit={handleSubmit}>
                <Stack spacing={8}>
                  <FormControl id="name" isRequired>
                    <FormLabel>ユーザー名</FormLabel>
                    <Input
                      type="text"
                      name="name"
                      id="name"
                      value={inputname}
                      onChange={handleInputName}
                    />
                  </FormControl>
                  <FormControl id="password" isRequired>
                    <FormLabel>パスワード</FormLabel>
                    <Input
                      type="password"
                      name="password"
                      value={inputpass}
                      onChange={handleInputPass}
                    />
                  </FormControl>
                  <Button colorScheme="blue" type="submit">
                    ログイン
                  </Button>
                  <Text width={"400px"} textAlign={"center"}>
                    <Link as={RouterLink} to="/create_user" color={"blue.400"}>
                      アカウントを作成
                    </Link>
                  </Text>
                </Stack>
              </form>
            </Center>
          </Stack>
        </Center>
      </Box>
    </>
  );
}

export default Login;
