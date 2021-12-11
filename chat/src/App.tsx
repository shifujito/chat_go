import { useState } from 'react'
import './App.css'
import './index.css'
import { apiClient, aCl } from './api-client'
import { Route, BrowserRouter }  from "react-router-dom"
import Login from './login'
import Post from './post';


function App() {
  const [count, setCount] = useState(0)

  const handleClick = () => {
    aCl.get("http://127.0.0.1:8080/users").then(res => console.log(res))
    // apiClient.users.get().then(res => {console.log(res)})
  }
  return (
    <BrowserRouter>
      <Route path="/login" component={Login}/>
      <Route path="/main" component={Post}/>
    </BrowserRouter>
  )
}

export default App
