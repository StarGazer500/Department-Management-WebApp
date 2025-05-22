import { useState } from 'react'

import './App.css'
import { Route, Routes } from "react-router-dom";

import AdminLoginForm from './accounts/Login';

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
    <Routes >
             
            <Route path="/login-user" element={<AdminLoginForm />} />
            
            
           
        </Routes>
     
    </>
  )
}

export default App
