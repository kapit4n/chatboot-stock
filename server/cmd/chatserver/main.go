package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	domain "github.com/kapit4n/chat-ws-go/internal/domain"
	handlers "github.com/kapit4n/chat-ws-go/internal/handlers"
	"gopkg.in/olahol/melody.v1"
)

const (
	ALLOWORIGINS = "http://localhost:3000"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{ALLOWORIGINS},
		AllowMethods:     []string{"PUT", "PATCH", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	m := melody.New()
	m.Config.MaxMessageSize = 2000
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	dataUsers := make(map[*melody.Session]*domain.UserInfo)
	rooms := make([]domain.RoomInfo, 3)

	rooms = append(rooms, domain.RoomInfo{Name: "General", Description: "The general goom where all people are located"})
	rooms = append(rooms, domain.RoomInfo{Name: "stockToom", Description: "Stock room"})

	lock := new(sync.Mutex)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Main Page")
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.GET("/ws", func(c *gin.Context) {
		fmt.Print("REACH HERE")
		err := m.HandleRequest(c.Writer, c.Request)
		if err != nil {
			fmt.Println(err.Error())
		}
	})

	r.GET("/users", func(c *gin.Context) {
		fmt.Println(dataUsers)

		for k, v := range dataUsers {
			fmt.Println(k)
			fmt.Println(v)
		}

		b, err := json.Marshal(&dataUsers)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, b)
	})

	r.GET("/uuid", func(c *gin.Context) {
		uuid := getUUID()
		c.JSON(http.StatusOK, uuid)
	})

	m.HandleConnect(func(s *melody.Session) {
		handlers.HandleConnection(s, m, lock, dataUsers, rooms)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		handlers.HandleDisconnect(s, m, lock, dataUsers, rooms)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		handlers.HandleMessage(s, m, lock, dataUsers, msg, rooms)
	})

	_ = r.Run(":8080")
}

func getUUID() map[string]interface{} {
	token, _ := uuid.NewUUID()
	result := map[string]interface{}{
		"uuid": token.String(),
	}
	return result
}
