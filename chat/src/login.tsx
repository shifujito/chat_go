import { format } from 'path'
import { useState } from 'react'
import { useHistory, Route }  from "react-router-dom"
import { apiClient, aCl } from './api-client'

type User = {
    id: number
    name: string
    match: boolean
}


function Login(){
    const [data, setData] = useState<User[]>([])

    const history = useHistory();

    const [inputname, setInputName] = useState<string>("")
    const [inputpass, setInputPass] = useState<string>("")

    const handleInputName = (e: React.ChangeEvent<HTMLInputElement>) => {
        setInputName(e.target.value)
    }

    const handleInputPass = (e: React.ChangeEvent<HTMLInputElement>) => {
        setInputPass(e.target.value)
    }

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        aCl.post("http://127.0.0.1:8080/logins", {
            name: inputname,
            password: inputpass
        }).then(res => {
            console.log(res)
            setData(res.data)
            history.push("/main")
        }).catch((err) => {
            console.log(err.response.status)
        })
    }

    return (
        <div>
            <a href="/create">sign up</a>
            <h1>Login</h1>
            {/* <h2>{{ .Message }}</h2> */}
            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor="name">User Name</label>
                    <input type="text" name="name" id="name" value={inputname} onChange={handleInputName}/>
                </div>
                <div>
                    <label htmlFor="name">password</label>
                    <input type="password" name="password" value={inputpass} onChange={handleInputPass}/>
                </div>
                <button type='submit'>ログイン</button>
            </form>
        </div>
    )
}

export default Login;


