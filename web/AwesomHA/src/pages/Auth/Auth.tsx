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
    const handleLoginButton = async () => {
        console.log('handleLoginButton called');
        if (!handleAuth) {
            console.log('handleAuth is not defined');
            throw new Error('handleAuth is not defined');
        }

        try {
            console.log('trying to login');
            if (!email || !password) {
                console.log('email or password is empty');
                throw new Error('Email and password are required');
            }

            const output = await handleAuth(email, password);
            console.log('login output', output);
            return !!output;
        } catch (error) {
            console.log('login error', error);
            if (!(error instanceof Error)) {
                throw new Error('Unexpected error');
            }

            setError(error);
            return false;
        }
    };
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
    
    const handleSignupButton = async () => {
        console.log('handleSignupButton called');
        if (password == null || confirmPassword == null) {
            console.log('password or confirmPassword is null');
            const error = new Error("Password and confirm password cannot be null");
            console.log('throwing error:', error);
            return setError(error);
        }
        console.log('password and confirmPassword are not null');
        if (password !== confirmPassword) {
            console.log('passwords do not match');
            const error = new Error("Passwords do not match");
            console.log('throwing error:', error);
            return setError(error);
        }

        try {
            console.log('trying to signup');
            const output = await handleAuth(name, email, password);
            console.log('signup output', output);
            if (output == null) {
                console.log('handleAuth returned null');
                const error = new Error("handleAuth returned null");
                console.log('throwing error:', error);
                throw error;
            }
            return !!output;
        } catch (error) {
            console.log('signup error', error);
            if (error == null) {
                console.log('handleAuth threw null exception');
                const error = new Error("handleAuth threw null exception");
                console.log('throwing error:', error);
                throw error;
            }
            if (!(error instanceof Error)) {
                console.log('handleAuth threw something that was not an instance of Error');
                const error = new Error("handleAuth threw something that was not an instance of Error");
                console.log('throwing error:', error);
                throw error;
            }
            console.log('setting error', error);
            setError(error);
        }
    };
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
