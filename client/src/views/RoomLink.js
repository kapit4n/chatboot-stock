import "./RoomLink.css"

export const RoomLink = ({ roomInfo, onClick, roomSelected }) => {
  return (
    <div className={roomSelected?.name === roomInfo?.name ? "room-link-selected": "room-link"}>
      # <a onClick={onClick}>{roomInfo.name}</a>
      <span>{roomInfo.newMessages}</span>
    </div>
  )
}

export default RoomLink;