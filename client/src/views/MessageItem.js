export const MessageItem = ({ message }) => {
  return (
    <div className="message-item">
      <img src={message.sender.image} />
      <div className="message-info">
        <h3>{message.sender.name}</h3>
        <div>{message.text}</div>
      </div>
    </div>
  )
}

export default MessageItem;