const Users = ({ users }) => {
  return (
    <ul className="user-list">
      <h2>Users</h2>
      {users.map(user => user.name && (
        <li key={user.name} className="user-list-item">{user.name}</li>
      ))}
    </ul>
  )
}

export default Users;