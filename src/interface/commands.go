package main

import (
	"fmt"
	"os"
)

var (
	Commands = make(map[string]Command)
)

func LoadCommands() {
	helpCommand := Command {
		Name: "help",
		Description: "Shows all possible commands",
		Usage: "help",
		Handle: Help,
	}

	LoadCommand(helpCommand)

	exitCommand := Command {
		Name: "exit",
		Description: "Exists the current session",
		Usage: "exit",
		Handle: Exit,
	}

	LoadCommand(exitCommand)
}

func Help(handle *CmdHandle) {
	if (len(handle.Args) > 0) {
		fmt.Println(fmt.Sprintf("--- Usage of command %s ---", handle.Args[0]))
		cmd, ok := Commands[handle.Args[0]]
		if !ok {
			fmt.Println("The entered command does not exist.")
			return
		}

		fmt.Println(fmt.Sprintf("Usage: %s", cmd.Usage))
		return
	}

	fmt.Println("List of possible commands")
	for _, cmd := range Commands {
		fmt.Println(fmt.Sprintf("%s: %s", cmd.Name, cmd.Description))
	}
}

func Exit(handle *CmdHandle) {
	fmt.Println("Byebye, until the next time!")
	os.Exit(0)
}

