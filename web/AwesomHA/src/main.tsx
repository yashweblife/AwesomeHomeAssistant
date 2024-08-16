import { ChakraProvider, extendTheme } from '@chakra-ui/react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'

const ThemeConfig = {
  initialColorMode: 'dark',
  useSystemColorMode: false
}

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <ChakraProvider theme={extendTheme(ThemeConfig)}>
    <App />
  </ChakraProvider>
)
