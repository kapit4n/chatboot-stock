const Users = ({ users }) => {
  return (
    <ul className="user-list">
      {users.map(user => (
        <li key={user.name} className="user-list-item">{user.name}</li>
      ))}
    </ul>
  )
}

export default Users;