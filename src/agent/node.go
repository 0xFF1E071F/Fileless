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

func initNode() {
	// Registering the node in the bootstrap server
	registerNode()
}
