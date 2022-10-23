package main

import (
	"fmt"
	"log"
	"net/url"
	"net/http"
	
	"github.com/gorilla/websocket"
)

func recvBootstrap(conn *websocket.Conn) string {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Print("Error while reading: ", err)
		return ""
	}
	
	conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	conn.Close()
	return string(msg)
}

func connectNode(nodeIP string) {
	// Using HTTP for now, to keep it simple (Testing purposes)
	// TODO: Upgrading this to WebSockets

	_, err := http.Get(fmt.Sprintf("http://%s:1337/cmd", nodeIP))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func callBootstrap() {
	servURL := url.URL{Scheme: "ws", Host: servAddr, Path: "/init"} 
	ws, _, err := websocket.DefaultDialer.Dial(servURL.String(), nil)
	if err != nil {
		log.Print("Error while dialing: ", err)
		return
	}

	defer ws.Close()
	
	cmd := fmt.Sprintf(arrayToString(sysInfo.Machine))
	err = ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s", cmd)))
	if err != nil {
		log.Print("Error while sending: ", err)
		return	
	}

	nodeIP := recvBootstrap(ws)
	if nodeIP != "" {
		connectNode(nodeIP)
	}
	return
}
