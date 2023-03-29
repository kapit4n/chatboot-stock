import RoomLink from "./RoomLink"

export const Rooms = ({ rooms }) => {
  return (
    <div className="rooms">
      {rooms.map(room => <RoomLink roomInfo={room} />)}
    </div>
  )
}

export default Rooms;