import { createContext, useState } from "react";


type AuthContext = {
    handleLogin: (email: string, password: string) => Promise<any>
    handleSignup: (name:string, email: string, password: string) => Promise<any>
    handleLogout: () => void
    isAuthenticated: boolean
    user: any
}


export const AuthContext = createContext<AuthContext>(
    {
        handleLogin: async (email: string, password: string) => false,
        handleSignup: async (name:string, email: string, password: string) => false,
        handleLogout: () => {},
        isAuthenticated: false,
        user: {}
    }
)

export default function AuthContextProvider({children}: {children: any}){
    const [isAuthenticated, setIsAuthenticated] = useState(false)
    const [userID, setUserID] = useState({})
    const handleLogin = async (email: string, password: string) => {
        try {
            const res = await fetch("http://localhost:8080/auth/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    email,
                    password
                })
            })
            return res
        } catch (error) {
            console.log(error)
        }
        
    }
    const handleSignup = async (name:string, email: string, password: string) => {
        try {
            const res = await fetch("http://localhost:8080/auth/signup", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    name,
                    email,
                    password
                })
            })
            return res
        } catch (error) {
            console.log(error)
        }
    }
    const output:AuthContext = {
        handleLogin,
        handleSignup,
        handleLogout: () => setIsAuthenticated(false),
        isAuthenticated,
        user: userID
    }

    return <AuthContext.Provider value={output}>{children}</AuthContext.Provider>
}