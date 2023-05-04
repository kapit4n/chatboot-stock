package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	domain "github.com/kapit4n/chat-ws-go/internal/domain"
	"github.com/kapit4n/chat-ws-go/internal/handlers"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	USERNAME                  = "bootUser"
	UUID                      = "0001"
	SERVER                    = "localhost"
	PORT                      = "8080"
	WSURL                     = "ws://" + SERVER + ":" + PORT + "/ws?user=" + USERNAME + "&uuid=" + UUID
	CHATBOOTCONNECTIONMESSAGE = "Chatboot is connected"
	REQUEST_URL               = "https://stooq.com/q/l/?s=%s.us&f=sd2t2ohlcv&h&e=json"
	AMQP_URL                  = "amqp://guest:guest@localhost:5672"
)

func main() {
	conn, err := amqp.Dial(AMQP_URL)

	if err != nil {
		log.Panicln(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

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
		log.Panicln(err)
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Panicln(err)
	}

	var runChan chan struct{}

	ctx := context.Background()

	dialer := websocket.Dialer{}

	c, _, err := dialer.DialContext(ctx, WSURL, nil)
	if err != nil {
		log.Panicf("Dial failed: %#v\n", err)
	}
	defer c.Close()

	go func() {
		for stockToken := range msgs {
			processingMessageJson, err := buildMessage(fmt.Sprintf(handlers.CHATBOOTPROCESSINGMESSAGE, string(stockToken.Body)))
			c.WriteMessage(1, processingMessageJson)

			requestUrl := fmt.Sprintf(REQUEST_URL, string(stockToken.Body))

			resp, err := http.Get(requestUrl)

			if err != nil {
				log.Fatal(err)
			}

			b, err := io.ReadAll(resp.Body)

			var respBody map[string]interface{}
			var symbols map[string]interface{}

			err = json.Unmarshal(b, &respBody)

			if err != nil {
				log.Fatal(err)
			}

			symbols = respBody["symbols"].([]interface{})[0].(map[string]interface{})

			if symbols["close"] == nil {
				processedMessageJson, err := buildMessage(fmt.Sprintf("%s quote is unknow", string(stockToken.Body)))

				if err != nil {
					log.Fatal(err)
				}
				c.WriteMessage(1, processedMessageJson)
			} else {
				stockClosedFloat := symbols["close"].(float64)
				stockClosed := fmt.Sprintf("%f", stockClosedFloat)

				processedMessageJson, err := buildMessage(fmt.Sprintf("%s quote is $%s per share", string(stockToken.Body), stockClosed))

				if err != nil {
					log.Fatal(err)
				}

				c.WriteMessage(1, processedMessageJson)
			}
		}
	}()

	<-runChan
}

func buildMessage(message string) ([]byte, error) {
	responseMessage := domain.WebsocketData{
		Channel: handlers.CHANNELNAME,
		Event:   handlers.EVENTNAME,
		Message: domain.DataMessage{
			MessageID: strconv.FormatInt(int64(time.Now().Unix()), 10),
			Message:   message,
			Sender:    handlers.CHATBOOTNAME,
		},
	}
	return json.Marshal(responseMessage)
}
