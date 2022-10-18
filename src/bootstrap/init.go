package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

func handleInit(ws *websocket.Conn, cmd []byte, msgType int) {
	if (string(cmd) == "x86_64" || string(cmd) == "x86") {
		fmt.Println("[INFO] - Node recognized")
	}
}
