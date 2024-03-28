package service

import (
	socket "ChatDemo/service/websocket"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Connect(c *gin.Context) {
	userId, _ := c.Get("userId")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print(err)
		return
	}
	socket.NodeManagerApi.AddNode(userId.(uint), conn)
}
