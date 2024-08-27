import { Button, Input, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay } from "@chakra-ui/react";
import { useState } from "react";

type AddDeviceModalProps = {
    isOpen: boolean
    onClose: () => void
}
export default function AddDeviceModal({ isOpen, onClose }: AddDeviceModalProps) {
    const [modalDeviceName, setModalDeviceName] = useState('')
    const [modalDeviceUrl, setModalDeviceUrl] = useState('http://192.168.')

    const checkIfURLIsValid = async () => {
        try {
            const test = await fetch("http://localhost:8080/device/validity", {method:"GET", body: JSON.stringify({url: modalDeviceUrl})})
            const {isValid} = await test.json()
            return isValid
        } catch (error) {
            throw error
        }
    }

    const handleAddButton = async () => {
        const output = {
            name: modalDeviceName,
            url: modalDeviceUrl
        }
        try {
            const isValid = await checkIfURLIsValid()
            if (!isValid) {    
                const test = await fetch("http://localhost:8080/device/register", {
                method: "POST",
                body: JSON.stringify(output),
            })
            if (!test.ok) {
                throw new Error("error")
            }
            onClose()
        }
        } catch (error) {
        }
    }
    return (
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
    )
}