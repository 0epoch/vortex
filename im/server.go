package im

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"vortex/model"
	"vortex/service"
)

type Server struct {
	clients map[int64] *Client
	cLock sync.RWMutex
	broadcast chan []byte //接收消息通道
}

func NewServer () *Server {
	return &Server{
		clients:   make(map[int64]*Client),
		broadcast: make(chan []byte),
	}
}

func(s *Server) Start() {
	for {
		select {
		case message := <- s.broadcast:
			fmt.Println("接收消息...")
			err := s.Send(message)
			if err != nil {
				continue
			}
		}
	}
}

//发送好友消息
func (s *Server) friendMessage(friendID int64, message []byte) error {
	client, ok := s.clients[friendID]
	if !ok {
		//TODO 发送离线消息
		return errors.New("好友未上线")
	}
	fmt.Println(friendID, "给好友消息...")
	client.msgChan <- message
	return nil
}

//发送群消息
func (s *Server) groupMessage(groupID int64, message []byte) error {
	userGroup, _ := service.FindGroupUserByID(groupID)
	for _, v := range userGroup {
		if client, ok := s.clients[v.UserID]; ok {
			client.msgChan <- message
		}
	}
	return nil
}


//同步离线消息
func (s *Server) SyncMessage() {

}

func (s *Server) Send(data []byte) error {
	message := model.Message{}
	fmt.Println("run send .....")
	err := json.Unmarshal(data, &message)
	if err != nil {
		//log.Fatal("消息格式错误", err)
		return err
	}
	switch message.Type {
	case model.MSG_TYPE_F:
		err = s.friendMessage(message.ToID, data)
	case model.MSG_TYPE_G:
		err = s.groupMessage(message.ToID, data)
	}
	return err
}

func (s *Server) AddClient(client *Client) {
	fmt.Println("添加连接....")
	s.cLock.Lock()
	s.clients[client.userID] = client
	fmt.Println("当前全部客户端：", s.clients)
	s.cLock.Unlock()
}

func (s *Server) RemoveClient(client *Client) {
	s.cLock.Lock()
	delete(s.clients, client.userID)
	s.cLock.Unlock()
}