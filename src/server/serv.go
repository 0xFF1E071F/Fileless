package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func wsHandle(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error while upgrading: ", err)
		return
	}

	defer ws.Close()

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Print("Error while reading: ", err)
			break
		}

		fmt.Println("Received: ", string(msg))
		err = ws.WriteMessage(msgType, []byte("Hello from server"))
		if err != nil {
			log.Print("Error during sending: ", err)
			break
		}
	}
}
func startServ() {
	http.HandleFunc("/init", wsHandle)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
