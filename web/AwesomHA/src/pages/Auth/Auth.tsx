import { useState } from "react"

function LoginBox({setState}:{setState:(val:boolean)=>void}){

    return <></>
}

function SignupBox({setState}:{setState:(val:boolean)=>void}){

    return <></>
}

export default function Auth() {
    const [login, setLogin] = useState(true)
    return(
        <>
        {
            login ? <LoginBox setState={setLogin}/> : <SignupBox setState={setLogin}/>
        }
        </>
    )
}
