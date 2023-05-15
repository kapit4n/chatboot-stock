import React, { useState, useEffect, useCallback, useRef } from 'react';
import LeftSide from './views/LeftSide'


import './App.css';
import ChatContainer from './views/ChatContainer';
import MainContainer from './views/MainContainer';
import Messages from './views/Messages';
import ChatInput from './views/ChatInput';
import Rooms from './views/Rooms';
import Users from './views/Users';
import Search from './views/Search';

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
  const [selectedRoom, setSelectedRoom] = useState("")

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
      console.log(rooms)
      setLastMessage(message)
      setRooms(rooms)
      setUsers(users)
      setSelectedRoom(rooms[0])
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
    <div className="app-container">
      {authenticated && (
        <ChatContainer>
          <LeftSide>
            <Rooms rooms={rooms} setSelectedRoom={setSelectedRoom} roomSelected={selectedRoom} />
            <Users users={users} />
          </LeftSide>
          <MainContainer>
            <Search />
            <Messages messages={messages} />
            <ChatInput sendMessage={sendMessage} message={stockInfo} setMessage={setStockInfo} />
          </MainContainer>
        </ChatContainer>
      )}

      <div className="login-box">
        {!authenticated && (
          <>
            <input onChange={e => setUserInput(e.target.value)} value={userInput} />
            <button onClick={loginUser}>
              LOGIN
            </button>
          </>
        )}
      </div>

    </div >
  );
}

export default App;