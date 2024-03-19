package wss

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/IgnacyBialobrzewski/go-chat/views/fragments"
	"golang.org/x/net/websocket"
)

type ChatMessage struct {
	Content string `json:"message"`
}

type WsServer struct {
	clients map[*websocket.Conn][]ChatMessage
}

func NewChatWss() *WsServer {
	return &WsServer{
		clients: make(map[*websocket.Conn][]ChatMessage, 1024),
	}
}

func (self *WsServer) HandleWs(client *websocket.Conn) {
	log.Println("incoming connection from: ", client.RemoteAddr())

	self.clients[client] = []ChatMessage{}
	self.readWs(client)
}

func (self *WsServer) readWs(client *websocket.Conn) {
	msgBuf := make([]byte, 1024)

	defer client.Close()

	for {
		bytesRead, err := client.Read(msgBuf)

		if err != nil {
			log.Println("failed to read: ", err)
			break
		}

		var msg ChatMessage

		err = json.Unmarshal(msgBuf[:bytesRead], &msg)

		if err != nil {
			log.Println("failed to unmarshal: ", err)
			continue
		}

		if strings.Trim(msg.Content, " ") == "" {
			return
		}

		self.clients[client] = append(self.clients[client], msg)
		self.emitMsg(msg)

		log.Println("new message: ", msg)
	}
}

func (self *WsServer) emitMsg(msg ChatMessage) {
	for client := range self.clients {
		views.Message(msg.Content).Render(context.Background(), client)
	}
}
