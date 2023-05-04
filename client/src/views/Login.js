import { useState } from "react"

export const Login = () => {

  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")

  return (
    <div className="login">
      <label for="username" />
      <input id="username" onChange={e => setUsername(e.target.value)} value={username} />
      <label for="password" />
      <input type="password" id="password" onChange={e => setPassword(e.target.value)} value={password} />
      
      <div className="login-actions">
        <button onClick={loginUser}>LOGIN</button>
        <button onClick={loginUser}>CANCEL</button>
      </div>
    </div>
  )
}