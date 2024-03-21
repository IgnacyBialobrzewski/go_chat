package models

import "golang.org/x/net/websocket"

type ChatMessage struct {
	Content string `json:"message"`
	Sender  *websocket.Conn
}
