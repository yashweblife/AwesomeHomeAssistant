import { Box, Button, HStack, IconButton, Input, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay, Tooltip, useDisclosure, VStack } from "@chakra-ui/react";
import { useState } from "react";
import { BsFillDeviceSsdFill } from 'react-icons/bs';
import { FaCog, FaPlus, FaRegUser, FaRobot } from 'react-icons/fa';
import { SensorDisplay } from "../../components/SensorDisplay";
export default function Dashboard() {
  const {isOpen, onOpen, onClose} = useDisclosure()
  const [modalDeviceName, setModalDeviceName] = useState('')
  const [modalDeviceUrl, setModalDeviceUrl] = useState('http://192.168.')
  const handleAddButton = async ()=>{
    const output = {
      name: modalDeviceName,
      url: modalDeviceUrl
    }
    try {
      const test = await fetch("http://localhost:8080/device/register", {
        method: "POST",
        body: JSON.stringify(output),
      }) 
      if(!test.ok){
        throw new Error("error")
      }
      onClose()
    } catch (error) {
    }
  }
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
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Add Device</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Input placeholder="Device Name" value={modalDeviceName} onChange={(e) => setModalDeviceName(e.target.value)}></Input>
            <Input placeholder="Device URL" value={modalDeviceUrl} onChange={(e) => setModalDeviceUrl(e.target.value)}></Input>
          </ModalBody>
          <ModalFooter>
            <Button colorScheme='blue' mr={3} onClick={onClose}>
              Close
            </Button>
            <Button variant='ghost' onClick={handleAddButton}>Add</Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </Box>
  )
}
