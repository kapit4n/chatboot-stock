package handlers

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	domain "github.com/kapit4n/chat-ws-go/internal/domain"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/olahol/melody.v1"
)

var (
	key   = []byte("testChatBoot")
	store = sessions.NewCookieStore(key)
)

func HandleConnection(s *melody.Session, m *melody.Melody, lock *sync.Mutex, dataUsers map[*melody.Session]*domain.UserInfo) {
	lock.Lock()

	var token string
	name := s.Request.URL.Query().Get("name")
	token = s.Request.URL.Query().Get("uuid")

	if token == "" {
		t, _ := uuid.NewUUID()
		token = t.String()
	}

	chatUser := domain.UserInfo{
		Name: name,
		UUID: token,
	}

	fmt.Println(chatUser)

	s.Set("chat-user", chatUser)

	message := domain.DataMessage{
		Message: strings.TrimLeft(name, " ") + " join stockRoom",
		Sender:  chatUser.Name,
	}
	lock.Unlock()

	data := domain.WebsocketData{
		Channel: "stockRoom",
		Event:   "status",
		Message: message,
	}
	b, _ := json.Marshal(data)
	_ = m.Broadcast(b)
}

func HandleDisconnect(s *melody.Session, m *melody.Melody, lock *sync.Mutex, dataUsers map[*melody.Session]*domain.UserInfo) {
	lock.Lock()

	fmt.Println(dataUsers)
	fmt.Println("DATA USERS BEFORE EACH DISCONNECTION")

	delete(dataUsers, s)

	chatUser := s.Keys["chat-user"].(domain.UserInfo)
	name := chatUser.Name
	message := domain.DataMessage{
		Message: name + " left stockRoom",
		Sender:  name,
	}
	lock.Unlock()

	data := domain.WebsocketData{
		Channel: CHANNELNAME,
		Event:   "status",
		Message: message,
	}
	b, _ := json.Marshal(data)
	_ = m.Broadcast(b)
}

func HandleMessage(s *melody.Session, m *melody.Melody, lock *sync.Mutex, dataUsers map[*melody.Session]*domain.UserInfo, msg []byte) {
	var WsData domain.WebsocketData

	if err := json.Unmarshal(msg, &WsData); err != nil {
		panic(err)
	}

	if WsData.Channel == ROOMNAME {
		if WsData.Event == EVENTNAME {
			var message domain.DataMessage
			_ = mapstructure.Decode(WsData.Message, &message)

			token, _ := uuid.NewUUID()

			message.MessageID = token.String()
			message.Message = strings.Trim(message.Message, " ")
			message.Message = strings.Trim(message.Message, "\n")

			data := domain.WebsocketData{
				Channel: WsData.Channel,
				Event:   WsData.Event,
				Message: message,
			}
			b, _ := json.Marshal(data)
			_ = m.Broadcast(b)

		}
	}
}
