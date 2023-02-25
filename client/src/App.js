import React, { useState, useEffect, useCallback, useRef } from 'react';


import './App.css';

const SOCKET_URL = 'ws://localhost:8080/ws'
const STOCK_ROOM = 'stockRoom'
const SEND_EVENT_NAME = "send_message_text"

function App() {
  const [stockInfo, setStockInfo] = useState("")
  const ws = useRef(null)

  const [messages, setMessages] = useState([])
  const [lastMessage, setLastMessage] = useState(null)
  const [userInput, setUserInput] = useState(null)
  const [authenticated, setAuthenticated] = useState(false)


  const loginUser = () => {
    const timeStamp = new Date().getTime()
    ws.current = new WebSocket(`${SOCKET_URL}?name=${userInput}&uuid=${timeStamp}`);
    ws.current.addEventListener('message', (event) => {
      const message = JSON.parse(event.data)
      setLastMessage(message)
    })
    setAuthenticated(true)
  }

  useEffect(() => {
    if (lastMessage) setMessages([...messages, lastMessage])
    return () => setMessages([])
  }, [lastMessage])

  const sendMessage = useCallback(() => {
    console.log(stockInfo)
    ws.current.send(
      JSON.stringify({ channel: STOCK_ROOM, event: SEND_EVENT_NAME, message: { message_id: (new Date()).getTime, message: stockInfo, sender: userInput } })
    )
  }, [ws.current, stockInfo])

  console.log(messages)

  return (
    <div className="App">
      <div className="login-box">
        User: <input onChange={e => setUserInput(e.target.value)} value={userInput} />
        {!authenticated && (
          <button onClick={loginUser}>
            LOGIN
          </button>

        )}
      </div>
      <ul className='chatroom-box'>
        {messages.map(m => (
          <>
            {m?.message?.sender === userInput ?
              (<li key={m.message.id} className="chat-post-me">
                {m?.message?.message}
              </li>)
              :
              (
                <li key={m.message.id} className="chat-post">
                  {m?.message?.sender}: {m?.message?.message}
                </li>)}
          </>

        ))}
        <input onChange={e => setStockInfo(e.target.value)} value={stockInfo} />
        <button onClick={() => sendMessage()}>Send</button>
      </ul>
    </div >
  );
}

export default App;
