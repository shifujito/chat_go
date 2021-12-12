import { HStack, Text } from "@chakra-ui/react";

type Props = {
    userId: number
    postUserId: number
}

export const PostIcons: React.VFC<Props> = ({postUserId, userId}) => {
    console.log("a")
    console.log(postUserId)
    console.log(userId)
    const deleteActivate = () =>{
        console.log("a")
        console.log(postUserId)
        console.log(userId)
        if (postUserId == userId)(<Text>削除</Text>)
    }

    return (
        <HStack>
            <Text>いいね</Text>
            <Text>リツイート</Text>
            {deleteActivate}
        </HStack>
    )
}
