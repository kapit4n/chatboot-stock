export const Profile = ({userInfo}) => {
  return (
    <div className="profile">
      <img src={userInfo.image} />
      <h3>{userInfo.name}</h3>
      <p>{userInfo.profession}</p>
      <span>{userInfo.status}</span>
      <button>Message</button>
      <button>Huddle</button>
    </div>
  )
}

export default Profile;