import { Button, ButtonGroup, Card, CardBody, Flex, Heading, Input, Stack, Text } from "@chakra-ui/react"
import { useState } from "react"

function LoginBox({ setState }: { setState: (val: boolean) => void, setError: (val: string) => void }) {
    function handleLoginButton(){

    }
    return (
        <Card>
            <CardBody>
                <Stack spacing={3}>
                    <Heading>Login</Heading>
                    <Input placeholder='Email' variant={'filled'} />
                    <Input placeholder='Password' variant={'filled'} />
                    <ButtonGroup justifyContent={'space-between'}>
                        <Button onClick={() => setState(false)}>Signup</Button>
                        <Button colorScheme="teal" onClick={handleLoginButton}>Login</Button>
                    </ButtonGroup>
                </Stack>
            </CardBody>
        </Card>
    )
}

function SignupBox({ setState }: { setState: (val: boolean) => void, setError: (val: string) => void }) {
    function handleSignupButton(){
        
    }
    return (
        <Card>
            <CardBody>
                <Stack spacing={3}>
                    <Heading>Signup</Heading>
                    <Input placeholder='Email' variant={'filled'} />
                    <Input placeholder='Name' variant={'filled'} />
                    <Input placeholder='Password' variant={'filled'} />
                    <Input placeholder='ConfirmPassword' variant={'filled'} />
                    <ButtonGroup justifyContent={'space-between'}>
                        <Button onClick={() => setState(true)}>Login</Button>
                        <Button colorScheme="teal" onClick={handleSignupButton}>Enter</Button>
                    </ButtonGroup>
                </Stack>
            </CardBody>
        </Card>
    )
}

export default function Auth() {
    const [login, setLogin] = useState(true)
    const [error, setError] = useState("")
    return (
        <Flex justifyContent={'center'} alignItems={'center'} height={'100vh'} flexDirection={'column'}>
            {
                login ? <LoginBox setState={setLogin} setError={setError} /> : <SignupBox setState={setLogin} setError={setError} />
            }
            {
                error ? <Text color="rgb(255,100,100)">{error}</Text> : null
            }
        </Flex>
    )
}
