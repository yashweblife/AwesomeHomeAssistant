import { Card, CardBody, IconButton, Text } from "@chakra-ui/react"
import { useEffect, useState } from "react"
import { CircularProgressbarWithChildren } from 'react-circular-progressbar'
import "react-circular-progressbar/dist/styles.css"
import { FaPause, FaPlay } from "react-icons/fa"
type SensorDisplayProps = {
    name: string
    url: string
}

export function SensorDisplay({ name, url }: SensorDisplayProps) {
    const [data, setData] = useState(0)
    const [isURLValid, setIsURLValid] = useState(false)
    const [intervalId, setIntervalId] = useState<number | undefined>(undefined)
    const [isPaused, setIsPaused] = useState(true)
    useEffect(() => {
        const validateUrl = async (): Promise<number | undefined> => {
            try {
                const URL = url + "device/";
                const response = await fetch(URL, { method: 'GET' });
                console.log(response)
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }

                setIsURLValid(true);
                setIntervalId(setInterval(handleGetData, 1000));
            } catch (error) {
                console.error(error);
                return undefined;
            }
        }
        const result = validateUrl().catch(err => err);
    }, [])

    async function handleGetData() {
        try {
            const URL = url + 'device/value';
            const response = await fetch(URL, { cache: 'no-store' });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            if (response.status === 204) {
                setData(1024)

                return
            }
            const { value } = await response.json();
            if (value === null) {
                throw new Error('Received null value from server');
            }
            setData(Number(value));

        } catch (error) {
            if (!(error instanceof Error)) {
                throw new Error('Unexpected error');
            }
            console.log('error', error);
        }
    }
    const handlePause = () => {
        if (intervalId) {
            clearInterval(intervalId);
            setIntervalId(undefined);
            setIsPaused(true);
        } else {
            setIntervalId(setInterval(handleGetData, 1000));
            setIsPaused(false);
        }
    }
    return (
        <Card backgroundColor={'rgb(22,22,22)'} p={2}>
            {
                isURLValid ?
                    <CardBody width={300} height={300}>

                        <CircularProgressbarWithChildren value={data} strokeWidth={5} styles={{
                            path: {
                                stroke: '#00d493'
                            },
                            trail: {
                                stroke: '#292f56'
                            }
                        }}
                            minValue={0}
                            maxValue={1024}
                        >
                            <Text color={'rgb(255,255,255)'} fontSize={20}>{name.toUpperCase()}</Text>
                            <Text color={'rgb(255,255,255)'}>
                                {Math.round(data / 1024 * 100)}%
                            </Text>
                            <IconButton onClick={handlePause} aria-label="" color='white' fontSize={'20'} isRound variant='none' icon={isPaused ? <FaPlay /> : <FaPause />}></IconButton>
                        </CircularProgressbarWithChildren>
                    </CardBody>
                    :
                    <CardBody>
                        <Text color={'rgb(255,255,255)'}>Connection Failed</Text>1
                    </CardBody>
            }
        </Card>
    )
}