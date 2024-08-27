import { Box, HStack, useDisclosure, VStack } from "@chakra-ui/react";
import AddDeviceModal from "../../components/AddDeviceModal";
import Chart from "../../components/Chart";
import DashboardSidebar from "../../components/DashboardSidebar";
import { SensorDisplay } from "../../components/SensorDisplay";
export default function Dashboard() {
  const { isOpen, onOpen, onClose } = useDisclosure();
  return (
    <Box width="100vw" height="100vh" backgroundColor={'rgb(32,32,32)'}>
      <HStack width="100vw" height="100vh">
        <DashboardSidebar sidebarOpen={onOpen} />
        <VStack height={'100%'} p={2}>
          <SensorDisplay name="light" url="http://localhost:8080/"></SensorDisplay>
        </VStack>
        <VStack>
          <Chart />
        </VStack>
      </HStack>
      <AddDeviceModal isOpen={isOpen} onClose={onClose} />
    </Box>
  )
}
