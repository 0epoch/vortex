package vortexnet

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Client struct {
	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("连接成功...", conn.RemoteAddr())
	client := &Client{conn: conn}
	fmt.Println(client)
}