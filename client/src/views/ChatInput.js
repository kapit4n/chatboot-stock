import "./ChatInput.css"

export const ChatInput = ({ sendMessage, message, setMessage }) => {

  return (
    <div className="chat-input">
      <textarea placeholder="Message #chat-room" value={message} onChange={e => setMessage(e.target.value)} />
      <div className="chatroom-actions">
        <button className="chatroom-actions-button" onClick={sendMessage}>Send</button>
      </div>
    </div>
  )
}

export default ChatInput;