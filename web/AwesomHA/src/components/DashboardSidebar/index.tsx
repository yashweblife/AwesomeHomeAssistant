import { IconButton, Tooltip, VStack } from "@chakra-ui/react";
import { BsFillDeviceSsdFill } from "react-icons/bs";
import { FaCog, FaPlus, FaRegUser, FaRobot } from "react-icons/fa";

type SidebarProps = {
    onOpen: () => void
}

export default function AddDeviceModal({ onOpen }: SidebarProps) {
    return (
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
    )
}