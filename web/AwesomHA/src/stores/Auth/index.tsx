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
                mode: "no-cors",
                headers: {
                    "Content-Type": "application/json"
                },
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
            const response = await fetch("http://localhost:8080/auth/register", {
                method: "POST",
                mode: "no-cors",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    name,
                    email,
                    password
                })
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return await response.json();
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