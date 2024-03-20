package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IgnacyBialobrzewski/go-chat/controllers"
	"github.com/IgnacyBialobrzewski/go-chat/wss"
	"golang.org/x/net/websocket"
)

var (
	ServerHost string = "0.0.0.0"
	ServerPort string = os.Getenv("PORT")
)

func main() {
	if ServerPort == "" {
		ServerPort = "3000"
	}

	fmt.Printf("Starting a HTTP server at http://%s:%s\n", ServerHost, ServerPort)

	controller.Index{}.Routes()
	chatWss := wss.NewChatWss()

	http.Handle("/ws", websocket.Handler(chatWss.HandleWs))
	http.ListenAndServe(ServerHost+":"+ServerPort, nil)
}
