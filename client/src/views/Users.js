import "./Users.css"

const Users = ({ users }) => {
  return (
    <div className="users-list-container">
      <h2>Users</h2>
      <ul className="user-list">
        {users.map(user => user.name && (
          <li key={user.name} className="user-list-item">{user.name}</li>
          ))}
      </ul>
    </div>
  )
}

export default Users;