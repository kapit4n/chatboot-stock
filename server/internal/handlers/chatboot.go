package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	domain "github.com/kapit4n/chat-ws-go/internal/domain"
)

func HandleChatBoot(c *websocket.Conn, done *chan struct{}) {
	defer c.Close()
	defer close(*done)

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

			processingMessageJson, err := buildMessage(fmt.Sprintf(CHATBOOTPROCESSINGMESSAGE, stockAbr))
			c.WriteMessage(1, processingMessageJson)

			requestUrl := fmt.Sprintf("https://stooq.com/q/l/?s=%s.us&f=sd2t2ohlcv&h&e=json", stockAbr)

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
				processedMessageJson, err := buildMessage(fmt.Sprintf("%s quote is unknow", stockAbr))

				if err != nil {
					log.Fatal(err)
				}
				c.WriteMessage(1, processedMessageJson)
			} else {
				stockClosedFloat := symbols["close"].(float64)
				stockClosed := fmt.Sprintf("%f", stockClosedFloat)

				processedMessageJson, err := buildMessage(fmt.Sprintf("%s quote is $%s per share", stockAbr, stockClosed))

				if err != nil {
					log.Fatal(err)
				}

				c.WriteMessage(1, processedMessageJson)
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
	msgMap, _ := msgInterface.(map[string]interface{})

	if msgMap["message"] != nil {
		messageInfo = msgMap["message"].(string)
	}

	return messageInfo
}

func buildMessage(message string) ([]byte, error) {
	responseMessage := domain.WebsocketData{
		Channel: CHANNELNAME,
		Event:   EVENTNAME,
		Message: domain.DataMessage{
			MessageID: strconv.FormatInt(int64(time.Now().Unix()), 10),
			Message:   message,
			Sender:    CHATBOOTNAME,
		},
	}
	return json.Marshal(responseMessage)
}
