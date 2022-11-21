package main

import (
  "fmt"
  "encoding/json"
)

type AgentCommand struct {
  ID        int
  Args      []string
  Salt      string
}

var AgentCommands map[string]AgentCommand = map[string]AgentCommand {
  "ping": {
    ID: 0x01,
    Args: make([]string, 0),
    Salt: "salt", //Should be generated when handling command
  },

  "die": {
    ID: 0xDEAD,
    Args: make([]string, 0),
    Salt: "salt", //Should be generated when handling command
  },
}


func BuildAgentCmd(handle *CmdHandle) (string, string){
  agentCmd, _ := AgentCommands[handle.Name]
  copy(agentCmd.Args, handle.Args) //Copying over the arguments

  // Generating random salt
  var salt string = GenerateRandomString(16)
  agentCmd.Salt = salt 

  cmd, err := json.Marshal(agentCmd)
  CheckErr(err)

  var sig string = SignCmd(string(cmd))

  return string(cmd), sig
}

func BroadcastCommand(handle *CmdHandle) {
  _, sig := BuildAgentCmd(handle)

  fmt.Println("Sig:", sig)
}
