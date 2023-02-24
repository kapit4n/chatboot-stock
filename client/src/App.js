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
  

  const loginUser = () => {
    const timeStamp = new Date().getTime()
    ws.current  = new WebSocket(`${SOCKET_URL}?name=${userInput}&uuid=${timeStamp}`);
    ws.current.addEventListener('message', (event) => {
      const message = JSON.parse(event.data)
      setLastMessage(message)
    })
  }

  useEffect(() => {
    if (lastMessage) setMessages([...messages, lastMessage])
    return () => setMessages([])
  }, [lastMessage])

  const sendMessage = useCallback((msg) => {
    ws.current.send(
      JSON.stringify({channel: STOCK_ROOM, event: SEND_EVENT_NAME, message: {message_id: (new Date()).getTime, message: msg}})
    )
  }, [ws.current])


  return (
    <div className="App">
      <header className="App-header">
      <div>
        User: <input onChange={e => setUserInput(e.target.value)} value={userInput} />
        <button onClick={loginUser}>
          LOGIN
          </button>
      </div>
      <div>
      {messages.map(m => (
        <li>{m?.message?.message || m?.message?.info}</li>
      ))}
      <input onChange={e => setStockInfo(e.target.value)} value={stockInfo} />
      <button onClick={() => sendMessage(`/stock=${stockInfo}`)}>Send</button>
    </div>
      </header>
    </div>
  );
}

export default App;
