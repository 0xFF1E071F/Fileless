package main

import (
	"fmt"
	"log"
	"net/http"
)

func initAgent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, new agent!")
	fmt.Println("[INFO] - Received new request from agent")
}

func startServ() {
	http.HandleFunc("/init", initAgent)

	log.Fatal(http.ListenAndServe(":8082", nil))
}
