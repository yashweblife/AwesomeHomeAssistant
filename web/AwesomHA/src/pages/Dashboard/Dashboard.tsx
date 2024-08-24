import { Box, HStack, IconButton, Tooltip, useDisclosure, VStack } from "@chakra-ui/react";
import { BsFillDeviceSsdFill } from 'react-icons/bs';
import { FaCog, FaPlus, FaRegUser, FaRobot } from 'react-icons/fa';
import AddDeviceModal from "../../components/AddDeviceModal";
import { SensorDisplay } from "../../components/SensorDisplay";
export default function Dashboard() {
  const {isOpen, onOpen, onClose} =  useDisclosure();
  return (
    <Box width="100vw" height="100vh" backgroundColor={'rgb(32,32,32)'}>
      <HStack width="100vw" height="100vh">
        <VStack backgroundColor={'rgb(12,12,12)'} height={'100%'} p={2} spacing={10}>
          <Tooltip label="Profile" placement='right'>
            <IconButton aria-label="" color='white' fontSize={'20'} mb={10} isRound variant='none' icon={<FaRegUser />}>1</IconButton>
          </Tooltip>
          <Tooltip label="Devices" placement='right'>
            <IconButton aria-label="" color='white' fontSize={'20'} isRound variant='none' icon={<BsFillDeviceSsdFill />}>1</IconButton>
          </Tooltip>
          <Tooltip label="AI" placement='right'>
            <IconButton aria-label="" color='white' fontSize={'20'} isRound variant='none' icon={<FaRobot />}>2</IconButton>
          </Tooltip>
          <Tooltip label="Add" placement='right'>
            <IconButton aria-label="" color='white' onClick={onOpen} fontSize={'20'} isRound variant='none' icon={<FaPlus />}>1</IconButton>
          </Tooltip>
          <Tooltip label="Settings" placement='right'>
            <IconButton aria-label="" color='white' fontSize={'20'} isRound variant='none' icon={<FaCog />}>1</IconButton>
          </Tooltip>
        </VStack>
        <VStack height={'100%'} p={2}>
          <SensorDisplay name="light" url="http://localhost:8080/"></SensorDisplay>
        </VStack>
      </HStack>
      <AddDeviceModal isOpen={isOpen} onClose={onClose} />
    </Box>
  )
}
