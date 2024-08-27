export type Command = {
    name: string,
    info: string,
    command: string
}

export type Device = {
    name: string
    id: string
    status: string
    type: string
    room: string
    url: string
    commands:Command[]
}

export type User = {
    name: string
    email: string
    devices:Device
    id: string
}

export type Todo = {
    name: string
    description: string
    id: string
    user: User
    status: string
    date: string
}
export const SERVER_URL = "http://localhost:8080"