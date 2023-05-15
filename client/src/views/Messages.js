import MessageItem from "./MessageItem"
import "./Messages.css"

export const Messages = ({ messages }) => {
  return (
    <div className="messages">
      {messages.map(message => (<MessageItem message={message} />))}
    </div>
  )
}

export default Messages;