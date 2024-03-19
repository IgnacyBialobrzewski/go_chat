package main

import (
	"fmt"
	"net/http"

	"github.com/IgnacyBialobrzewski/go-chat/controllers"
	"github.com/IgnacyBialobrzewski/go-chat/wss"
	"golang.org/x/net/websocket"
)

const (
	ServerHost string = "localhost"
	ServerPort string = "3000"
	ServerAddr string = ServerHost + ":" + ServerPort
)

func main() {
	fmt.Printf("Starting a HTTP server at http://%s\n", ServerAddr)

	controller.Index{}.Routes()
	chatWss := wss.NewChatWss()

	http.Handle("/ws", websocket.Handler(chatWss.HandleWs))
	http.ListenAndServe(ServerAddr, nil)
}
