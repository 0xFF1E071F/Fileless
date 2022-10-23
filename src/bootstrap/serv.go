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
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error while upgrading: ", err)
		return
	}

	defer ws.Close()

	for { // Handling of the request will go here
		msgType, _, err := ws.ReadMessage()
		if err != nil {
			log.Print("Error while reading: ", err)
			break
		}


		err = ws.WriteMessage(msgType, []byte("Hello from server"))
		if err != nil {
			log.Print("Error during sending: ", err)
			break
		}
	}
}

func regNode(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "POST") {
		fmt.Fprintf(w, ".")
		return
	}

	if err := r.ParseForm(); err != nil {
		return
	}

	go registerNode(r.FormValue("Arch"), strings.Split(r.RemoteAddr, ":")[0])
}

func startServ(IP string, port int) {
	http.HandleFunc("/init", wsHandle)
	http.HandleFunc("/registerNode", regNode)
	fmt.Println("[INFO] - Started webserver successfully")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", IP, port), nil))
}
