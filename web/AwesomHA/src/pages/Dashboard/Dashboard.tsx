import { Box, HStack, IconButton, VStack } from "@chakra-ui/react";
import { BsFillDeviceSsdFill } from 'react-icons/bs';
import { FaCog, FaRegUser, FaRobot } from 'react-icons/fa';
export default function Dashboard() {
    return (
        <Box width="100vw" height="100vh" backgroundColor={'rgb(32,32,32)'}>
        <HStack width="100vw" height="100vh">
          <VStack backgroundColor={'rgb(12,12,12)'} height={'100%'} p={2} spacing={10}>
            <IconButton aria-label="" color='white' fontSize={'3xl'} isRound variant='none' icon={<FaRegUser />}>1</IconButton>
            <IconButton aria-label="" color='white' fontSize={'3xl'} isRound variant='none' icon={<BsFillDeviceSsdFill />}>1</IconButton>
            <IconButton aria-label="" color='white' fontSize={'3xl'} isRound variant='none' icon={<FaRobot />}>2</IconButton>
            <IconButton aria-label="" color='white' fontSize={'3xl'} isRound variant='none' icon={<FaCog />}>1</IconButton>
          </VStack>
        </HStack>
      </Box>
    )
}
