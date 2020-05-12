package im

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},

}

func ServeWs(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	userID, _ := c.Get("userID")
	client := &Client{userID:userID.(int64),conn: conn, msgChan: make(chan []byte, 256)}
	imServer.AddClient(client)
	go client.writePump()
	go client.readPump()
}