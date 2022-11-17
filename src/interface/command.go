package main

import (
	"fmt"
)


type CmdHandler func(handle *CmdHandle)

type CmdHandle struct {
	Name string
	Args []string
}

type Command struct {
	Name 		string
	Description string
	Usage       string
	Handle 	    CmdHandler
}

func LoadCommand(cmd Command) {
	_, ok := Commands[cmd.Name]

	if ok { //Command already loaded
		return;
	}

	Commands[cmd.Name] = cmd
}

func ExecCommand(handle *CmdHandle) {
	cmd, success := Commands[handle.Name]
	if !success {
		fmt.Println("This command does not exist.")
		return
	}

	cmd.Handle(handle)
}
