import { createContext } from "react";

type Device = {
    name: string
    url: string
    id: string
}

type UserMeta = {
    name: string
    email: string
    devices: Device[]
    id: string
}

const UserContext = createContext<UserMeta>({
    name: '',
    email: '',
    devices: [],
    id: ''
})

export default function UserProvider({ children }: any) {

    const getDeviceList = () =>{}
    const getDeviceCommands = () =>{}
    const output = { name: '', email: '', devices: [], id: '' }
    return <UserContext.Provider value={output}>{children}</UserContext.Provider>

}