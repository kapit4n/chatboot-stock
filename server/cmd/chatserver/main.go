package main

import (
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
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true } // origni check
	dataUsers := make(map[*melody.Session]*domain.UserInfo)

	lock := new(sync.Mutex)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.html", nil)
	})

	r.GET("sw.js", func(c *gin.Context) {
		c.Header("Content-Type", "application/javascript")
		c.HTML(http.StatusOK, "sw.js", nil)
	})
	r.GET("/ws", func(c *gin.Context) {
		fmt.Print("REACH HERE")
		err := m.HandleRequest(c.Writer, c.Request)
		if err != nil {
			fmt.Println(err.Error())
		}
	})

	r.GET("/uuid", func(c *gin.Context) {
		token, _ := uuid.NewUUID()
		result := map[string]interface{}{
			"uuid": token.String(),
		}
		c.JSON(http.StatusOK, result)
	})

	m.HandleConnect(func(s *melody.Session) {
		handlers.HandleConnection(s, m, lock, dataUsers)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		handlers.HandleDisconnect(s, m, lock, dataUsers)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		handlers.HandleMessage(s, m, lock, dataUsers, msg)
	})

	_ = r.Run(":8080")
}
