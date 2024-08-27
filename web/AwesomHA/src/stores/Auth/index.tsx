import { createContext, useState } from "react";
import { SERVER_URL } from "../../libs/utils";

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
            const res = await fetch(`${SERVER_URL}/auth/login`, {
                method: "POST",
                mode: "no-cors",
                body: JSON.stringify({ email, password })
            });

            if (!res.ok) {
                throw new Error(`HTTP error! status: ${res.status}`);
            }

            return await res.json();
        } catch (error) {
            console.error(error);
        }
    }
    const handleSignup = async (name:string, email: string, password: string) => {
        try {
            const response = await fetch(`${SERVER_URL}/auth/register`, {
                method: "POST", 
                body: JSON.stringify({
                    name,
                    email,
                    password
                })
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const output = await response.json();
            setUserID(output)
            setIsAuthenticated(true)
            return true;
        } catch (error) {
            console.error(error);
            throw error;
        }
    }

    const handleLogout = async()=>{
        if(!isAuthenticated) return;
        setIsAuthenticated(false)
    }
    const output:AuthContext = {
        handleLogin,
        handleSignup,
        handleLogout,
        isAuthenticated,
        user: userID
    }

    return <AuthContext.Provider value={output}>{children}</AuthContext.Provider>
}