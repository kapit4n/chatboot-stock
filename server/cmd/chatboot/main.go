package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	domain "github.com/kapit4n/chat-ws-go/internal/domain"
	"github.com/kapit4n/chat-ws-go/internal/handlers"
)

const (
	USERNAME                  = "bootUser"
	UUID                      = "0001"
	SERVER                    = "localhost"
	PORT                      = "8080"
	WSURL                     = "ws://" + SERVER + ":" + PORT + "/ws?user=" + USERNAME + "&uuid=" + UUID
	CHATBOOTCONNECTIONMESSAGE = "Chatboot is connected"
)

func main() {
	ctx := context.Background()

	dialer := websocket.Dialer{}

	c, _, err := dialer.DialContext(ctx, WSURL, nil)
	if err != nil {
		log.Panicf("Dial failed: %#v\n", err)
	}
	defer c.Close()

	token, _ := uuid.NewUUID()

	data := domain.WebsocketData{
		Channel: handlers.CHANNELNAME,
		Event:   handlers.EVENTNAME,
		Message: domain.DataMessage{
			MessageID: token.String(),
			Message:   CHATBOOTCONNECTIONMESSAGE,
			Sender:    USERNAME,
		},
	}
	d, err := json.Marshal(data)

	c.WriteMessage(1, []byte(d))

	stillConected := make(chan bool)

	done := make(chan struct{})

	go func() {
		handlers.HandleChatBoot(c, &done)
	}()

	fmt.Println(<-stillConected)
}
