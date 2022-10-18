package main

import (
	"fmt"
	"log"
	"flag"
	"os"
	"time"
	"net/url"
	
	"github.com/gorilla/websocket"
)

var servAddr = flag.String("addr", "192.168.146.133:8082", "http service address")
var done chan interface{}
var interrupt chan os.Signal

func recvHandler(conn *websocket.Conn) {
	defer close(done)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Print("Error while reading: ", err)
			return
		}

		fmt.Println("Received: ", string(msg))
	}
}

func initClient() {
	done = make(chan interface{})
	interrupt = make(chan os.Signal)
	
	servURL := url.URL{Scheme: "ws", Host: *servAddr, Path: "/init"} 
	ws, _, err := websocket.DefaultDialer.Dial(servURL.String(), nil)
	if err != nil {
		log.Print("Error while dialing: ", err)
		return
	}

	defer ws.Close()
	go recvHandler(ws)

	for {
		select {
		case <-time.After(time.Duration(1) * time.Millisecond * 1000):
			err := ws.WriteMessage(websocket.TextMessage, []byte("Hello from agent"))
			if err != nil {
				log.Print("Error while sending: ", err)
				return
			}

		case <-interrupt:
			log.Print("Received SIGINT interrupt, closing")

			err := ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error while sending: ", err)
				return
			}

			select {
			case <-done:
				log.Print("Closing recvHandler...")
			case <-time.After(time.Duration(1) * time.Second):
				log.Print("Timeout in closing receiving channel, exiting...")
			}

			return
		}
	}
}
