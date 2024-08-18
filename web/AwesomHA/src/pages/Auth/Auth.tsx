import { Button, ButtonGroup, Card, CardBody, Flex, Heading, Input, Stack, Text } from "@chakra-ui/react"
import { useContext, useState } from "react"
import { AuthContext } from "../../stores/Auth"

type AuthBoxProps = {
    setState: (val: boolean) => void
    setError: (val: Error) => void
    handleAuth: any
}

function LoginBox({ setState, setError, handleAuth }:AuthBoxProps) {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    async function handleLoginButton(){
        try {
            const output = await handleAuth(email, password)
            if (output) {
                return(true)
            }
        } catch (error: any) {
            setError(error)
            return(false)
        }
    }
    return (
        <Card>
            <CardBody>
                <Stack spacing={3}>
                    <Heading>Login</Heading>
                    <Input value={email} onChange={(e)=>{setEmail(e.target.value)}} placeholder='Email' variant={'filled'} />
                    <Input value={password} onChange={(e)=>{setPassword(e.target.value)}} placeholder='Password' variant={'filled'} />
                    <ButtonGroup justifyContent={'space-between'}>
                        <Button onClick={() => setState(false)}>Signup</Button>
                        <Button colorScheme="teal" onClick={handleLoginButton}>Login</Button>
                    </ButtonGroup>
                </Stack>
            </CardBody>
        </Card>
    )
}

function SignupBox({ setState, setError, handleAuth }: AuthBoxProps) {
    const [name, setName] = useState("")
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [confirmPassword, setConfirmPassword] = useState("")
    
    async function handleSignupButton(){
        if (password !== confirmPassword) {
            setError(new Error("Passwords do not match"))
            return
        }
        try {
            const output = await handleAuth(name, email, password)
            if (output) {
                return(true)
            }
        } catch (error: any) {
            setError(error)
            return;
        }    
    }
    return (
        <Card>
            <CardBody>
                <Stack spacing={3}>
                    <Heading>Signup</Heading>
                    <Input value={name} onChange={(e)=>{setName(e.target.value)}} placeholder='Name' variant={'filled'} />
                    <Input value={email} onChange={(e)=>{setEmail(e.target.value)}} placeholder='Email' variant={'filled'} />
                    <Input value={password} onChange={e=>{setPassword(e.target.value)}} placeholder='Password' variant={'filled'} />
                    <Input value={confirmPassword} onChange={(e)=>{setConfirmPassword(e.target.value)}} placeholder='ConfirmPassword' variant={'filled'} />
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
    const [error, setError] = useState<Error|null>(null)
    const {handleLogin, handleSignup} = useContext(AuthContext)
    return (
        <Flex justifyContent={'center'} alignItems={'center'} height={'100vh'} flexDirection={'column'}>
            {
                login ? <LoginBox setState={setLogin} setError={setError} handleAuth={handleLogin} /> : <SignupBox setState={setLogin} setError={setError} handleAuth={handleSignup} />
            }
            {
                error ? <Text color="rgb(255,100,100)">{error.message}</Text> : null
            }
        </Flex>
    )
}
