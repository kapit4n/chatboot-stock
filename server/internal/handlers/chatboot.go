package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	domain "github.com/kapit4n/chat-ws-go/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	AMQP_URL = "amqp://guest:guest@localhost:5672"
)

func HandleChatBoot(c *websocket.Conn, done *chan struct{}) {
	defer c.Close()
	defer close(*done)

	conn, err := amqp.Dial(AMQP_URL)

	if err != nil {
		log.Println(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Println(err)
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"stockQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Hour)
	defer cancel()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("read message: %#v\n", err)
			return
		}
		log.Printf("recv: %s\n", message)
		var mData domain.WebsocketData

		if err := json.Unmarshal(message, &mData); err != nil {
			panic(err)
		}

		var messageInfo string
		msgInterface := mData.Message

		messageInfo = getMessageInfo(msgInterface)

		tokenIndex := strings.Index(messageInfo, "stock=")

		if messageInfo != "" && strings.Contains(messageInfo, "stock=") {

			stockAbr := getStockToken(messageInfo, tokenIndex)

			body := stockAbr

			err = ch.PublishWithContext(ctx,
				"",
				queue.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				},
			)

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func getStockToken(messageInfo string, tokenIndex int) string {
	splitedStrings := strings.Split(messageInfo, " ")
	stockTokenMessage := ""

	for _, str := range splitedStrings {
		if strings.Contains(str, "stock=") {
			stockTokenMessage = str
		}
	}

	stockAbr := stockTokenMessage[tokenIndex+len("stock="):]

	return stockAbr
}

func getMessageInfo(msgInterface interface{}) string {

	var messageInfo string
	msgMap, _ := msgInterface.(domain.DataMessage)

	fmt.Println(msgInterface)

	if msgMap.Message != "" {
		messageInfo = msgMap.Message
	}

	return messageInfo
}
