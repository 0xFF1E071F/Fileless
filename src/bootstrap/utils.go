package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func parseConfig() bool {
	file, err := ioutil.ReadFile("config/settings.json")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	err = json.Unmarshal([]byte(file), &cfg)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
