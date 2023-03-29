import MessageItem from "./MessageItem"

export const Messages = ({ messages }) => {
  return (
    <div className="messages">
      {messages.map(message => (<MessageItem message={message} />))}
    </div>
  )
}

export default Messages;