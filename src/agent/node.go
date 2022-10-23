package main

import (
	"fmt"
	"net/url"
	"net/http"
)

func registerNode() bool {
	data := url.Values {
		"Arch": {arrayToString(sysInfo.Machine)},
	}

	_, err := http.PostForm(fmt.Sprintf("http://%s/registerNode", servAddr), data)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func handleCmd(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, to your node new agent!")
	fmt.Println("[INFO] ⁻ An agent sent something, whoo!")
}

func initServer() {
	http.HandleFunc("/cmd", handleCmd)
	http.ListenAndServe("0.0.0.0:1337", nil)
}

func initNode() {
	// Registering the node in the bootstrap server
	if(registerNode()) {
		initServer()
	}
}
