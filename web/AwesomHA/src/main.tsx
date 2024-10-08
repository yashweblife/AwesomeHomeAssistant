import { ChakraProvider, extendTheme } from '@chakra-ui/react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'
import AuthContextProvider from './stores/Auth'
const ThemeConfig = {
  initialColorMode: 'dark',
  useSystemColorMode: false
}

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <ChakraProvider theme={extendTheme(ThemeConfig)}>
    <AuthContextProvider>
      <App />
    </AuthContextProvider>
  </ChakraProvider>
)
