package main

import (
  "net"
)

var (
  SessionList = make(map[int]*Agent)
)

type Agent struct {
  SessionID int
  conn      net.Conn
}

func NewAgent(id int, c net.Conn) *Agent {
  return &Agent{id, c}
}

func AddToSessions(agent *Agent) {
  _, success := SessionList[agent.SessionID]
  if success {
    // What if ID already exists?
    return 
  }

  SessionList[agent.SessionID] = agent
}

func (this *Agent) Handle() {
  this.conn.Write([]byte("Handling initiated...\n"))
  /*
    TODO:
         - Adding transfering of public key
         - Adding agent to sessions list (done)
         - Removing agent when session closes
  */

  AddToSessions(this)
}
