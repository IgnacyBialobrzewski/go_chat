package wss

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/IgnacyBialobrzewski/go-chat/models"
	"github.com/IgnacyBialobrzewski/go-chat/views/fragments"
	"golang.org/x/net/websocket"
)

type WsServer struct {
	clients map[*websocket.Conn][]models.ChatMessage
}

func NewChatWss() *WsServer {
	return &WsServer{
		clients: make(map[*websocket.Conn][]models.ChatMessage, 1024),
	}
}

func (self *WsServer) HandleWs(client *websocket.Conn) {
	log.Println("incoming connection from: ", client.RemoteAddr())

	self.clients[client] = []models.ChatMessage{}
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

		msg := models.ChatMessage{
			Sender: client,
		}

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

func (self *WsServer) emitMsg(msg models.ChatMessage) {
	for client := range self.clients {
		isSenderReceiver := client == msg.Sender
		views.Message(msg.Content, isSenderReceiver).Render(context.Background(), client)
	}
}
