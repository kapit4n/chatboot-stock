import RoomLink from "./RoomLink"

export const Rooms = ({ rooms }) => {
  return (
    <div className="rooms">
      <h2>Rooms</h2>
      {rooms.map(room => room.name && <RoomLink roomInfo={room} />)}
    </div>
  )
}

export default Rooms;