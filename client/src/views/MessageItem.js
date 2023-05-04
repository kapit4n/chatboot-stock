export const MessageItem = ({ message }) => {
  return (
    <div className="message-item">
      <img width={50} height={50} src="https://www.mtsolar.us/wp-content/uploads/2020/04/avatar-placeholder.png" />
      <div className="message-info">
        <h3>{message?.sender}</h3>
        <div>{message.message}</div>
      </div>
    </div>
  )
}

export default MessageItem;