import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Auth from './pages/Auth/Auth'
import Dashboard from './pages/Dashboard/Dashboard'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route index element={<Auth />}/>
        <Route path="/dashboard" element={<Dashboard />}/>
      </Routes>
    </BrowserRouter>
  )
}

export default App
