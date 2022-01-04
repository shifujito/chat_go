import { Link } from "@chakra-ui/react";
import { useHistory } from "react-router-dom";

export const Profile: React.VFC = () => {
  // ログアウト機能を実装
  const history = useHistory();
  const handelLogout = () => {
    history.push("/login");
  };
  return <Link onClick={handelLogout}>ログアウト</Link>;
};
