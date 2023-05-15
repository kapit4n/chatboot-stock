import { useEffect, useRef } from 'react'
import "./ChatInput.css"

export const ChatInput = ({ sendMessage, message, setMessage }) => {


  const inputRef = useRef(null)

  useEffect(() => {
    if (inputRef) {
      inputRef.current.focus()
    }
  }, [])

  const sendMessageAnfFocus = () => {
    sendMessage();
    if (inputRef) {
      inputRef.current.focus()
    }
  }

  return (
    <div className="chat-input">
      <textarea ref={inputRef} placeholder="Message #chat-room" value={message} onChange={e => setMessage(e.target.value)} />
      <div className="chatroom-actions">
        <button className="chatroom-actions-button" onClick={sendMessageAnfFocus}>Send</button>
      </div>
    </div>
  )
}

export default ChatInput;