import "./ChatContainer.css"

export const ChatContainer = ({ children }) => {
  return (
    <div className="chat-container">
      {children}
    </div>
  )
}

export default ChatContainer;