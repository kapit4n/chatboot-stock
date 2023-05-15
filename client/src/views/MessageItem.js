
import "./MessageItem.css"

export const MessageItem = ({ message }) => {
  return (
    <div className="message-item">
      <div className="message-item-image">
       <img width={50} height={50} src="https://www.mtsolar.us/wp-content/uploads/2020/04/avatar-placeholder.png" />
      </div>
      <div className="message-info">
        <div className="message-info-title">
          <h3>{message?.sender}</h3>
          <span>11:44</span>
        </div>
        <div>{message.message}</div>
      </div>
    </div>
  )
}

export default MessageItem;