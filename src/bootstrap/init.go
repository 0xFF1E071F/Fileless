package main

import (
	"fmt"
	//"github.com/gorilla/websocket"
)

func registerNode(msg, IP string) {
	if (msg == "x86_64" || msg == "x86") {
		db.insertNode(IP)
		fmt.Println("[INFO] - Node recognized")
	}
}
