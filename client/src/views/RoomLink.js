export const RoomLink = ({ roomInfo, onClick }) => {
  return (
    <div className="room-link">
      # <a onClick={onClick}>{roomInfo.name}</a>
      <span>{roomInfo.newMessages}</span>
    </div>
  )
}

export default RoomLink;