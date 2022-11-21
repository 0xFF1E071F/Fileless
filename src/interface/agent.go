package main

type AgentCommand struct {
  ID      int
  Args    []string
  Padding string
}

var AgentCommands map[string]AgentCommand = map[string]AgentCommand {
  "ping": {
    ID: 0x01,
    Args: make([]string, 2),
    Padding: "test", //Should be generated when handling command
  },

  "die": {
    ID: 0xDEAD,
    Args: make([]string, 0),
    Padding: "test", //Should be generated when handling command
  },
}

