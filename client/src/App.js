import React, { useState, useEffect, useCallback, useRef } from 'react';


import './App.css';

const SOCKET_URL = 'ws://localhost:8080/ws'
const STOCK_ROOM = 'stockRoom'
const SEND_EVENT_NAME = "send_message_text"

function App() {
  const [stockInfo, setStockInfo] = useState("")
  const ws = useRef(null)

  const [messages, setMessages] = useState([])
  const [users, setUsers] = useState([])
  const [rooms, setRooms] = useState([])
  const [lastMessage, setLastMessage] = useState(null)
  const [userInput, setUserInput] = useState(null)
  const [authenticated, setAuthenticated] = useState(false)

  const loginUser = () => {
    const timeStamp = new Date().getTime()
    ws.current = new WebSocket(`${SOCKET_URL}?name=${userInput}&uuid=${timeStamp}`);
    setAuthenticated(true)

    ws.current.addEventListener("error", (event) => {
      console.log("WebSocket error: ", event);
      setAuthenticated(false)
    });

    ws.current.addEventListener("open", (event) => {
      console.log("OPEN event goes here", event);
    });

    ws.current.addEventListener('message', (event) => {
      const { message, rooms, users } = JSON.parse(event.data)
      console.log(message)
      setLastMessage(message)
      setRooms(rooms)
      setUsers(users)
    })
  }

  useEffect(() => {
    if (lastMessage) setMessages([...messages.slice(-9), lastMessage])
    return () => setMessages([])
  }, [lastMessage])

  const sendMessage = useCallback(() => {
    const stockInputCasted = stockInfo.replace("/", "")
    ws.current.send(
      JSON.stringify({ channel: STOCK_ROOM, event: SEND_EVENT_NAME, message: { message_id: (new Date()).getTime, message: stockInputCasted, sender: userInput } })
    )
    setStockInfo("")
  }, [ws.current, stockInfo])

  return (
    <div className="App">
      <div className="login-box">
        <h1>
          Users
        </h1>
        {authenticated ? (
          <ul className="user-list">
            {users.map(user => (<li key={user.name} className="user-list-item">{user.name}</li>))}
          </ul>
        ) : (
          <>
            <input onChange={e => setUserInput(e.target.value)} value={userInput} />
            <button onClick={loginUser}>
              LOGIN
          </button>
          </>
        )}
      </div>
      {authenticated && (
        <div className="chatroom-container">
          <ul className='chatroom-box'>
            {messages.map(m => (
              <>
                {m?.sender === userInput ?
                  (<li key={m.id} className="chat-post-me">
                    {m?.sender}: {m?.message}
                  </li>)
                  :
                  (
                    <li key={m.message.id} className="chat-post">
                      {m?.sender}: {m?.message}
                    </li>
                  )
                }
              </>
            ))}
          </ul>
          <div className="chatroom-actions">
            <textarea rows="4" onChange={e => setStockInfo(e.target.value)} value={stockInfo} className="chat-message" />
            <button className="chatroom-actions-button" onClick={() => sendMessage()}>Send</button>
          </div>
        </div>
      )}
    </div >
  );
}

export default App;
