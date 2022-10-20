package main

import (
	"fmt"
	"log"
	"strings"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func wsHandle(w http.ResponseWriter, r *http.Request) {
	IP := strings.Split(r.RemoteAddr, ":")[0]
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error while upgrading: ", err)
		return
	}

	defer ws.Close()

	for { // Handling of the request will go here
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Print("Error while reading: ", err)
			break
		}

		go handleInit(ws, msg, msgType, IP)
		/*
		err = ws.WriteMessage(msgType, []byte("Hello from server"))
		if err != nil {
			log.Print("Error during sending: ", err)
			break
		}
		*/
	}
}
func startServ(IP string, port int) {
	http.HandleFunc("/init", wsHandle)
	fmt.Println("[INFO] - Started webserver successfully")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", IP, port), nil))
}
