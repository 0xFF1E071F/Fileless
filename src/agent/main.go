package main

import (
	"fmt"
)
func main() {
	initRecon()
	if (arrayToString(sysInfo.Machine) == "x86" || arrayToString(sysInfo.Machine) == "x86_64") {
		fmt.Println("[INFO] - Calling initNode()")
		initNode()
	} else {
		initClient()
	}
}
