package pkg

type WebsocketData struct {
	Channel string      `json:"channel"`
	Event   string      `json:"event"`
	Message DataMessage `json:"message"`
	Rooms   []RoomInfo  `json:"rooms"`
	Users   []UserInfo  `json:"users"`
}

type DataMessage struct {
	MessageID string `json:"message_id"`
	Name      string `json:"name"`
	Message   string `json:"message"`
	Sender    string `json:"sender"`
	Time      string `json:"time"`
}

type UserInfo struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type RoomInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
