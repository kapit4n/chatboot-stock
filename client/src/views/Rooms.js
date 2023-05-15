import RoomLink from "./RoomLink"
import "./Rooms.css"

export const Rooms = ({ rooms, roomSelected }) => {
  return (
    <div className="rooms">
      <h2>Rooms</h2>
      {rooms.map(room => room.name && <RoomLink roomInfo={room} roomSelected={roomSelected}/>)}
    </div>
  )
}

export default Rooms;