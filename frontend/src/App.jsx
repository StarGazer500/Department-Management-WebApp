import { useState } from 'react'

import './App.css'
import { Route, Routes } from "react-router-dom";

import {LoginForm,CreateRole,CreateUserAccount} from './accounts/Accounts';
import {LecturerPage} from './users_pages/Lecturer'
// import CreateRole from './accounts/CreateRoles'
// import  CreateUserAccount from './accounts/Accounts'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
    <Routes >
             
            <Route path="/login-user" element={<LoginForm />} />
            <Route path="/create-role" element={<CreateRole />} />
            <Route path="/create-user" element={<CreateUserAccount />} />
            <Route path="/is-user-valid" element={<LecturerPage/>} />
             
            
            
           
        </Routes>
     
    </>
  )
}

export default App
