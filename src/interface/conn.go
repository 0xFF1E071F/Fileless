package main

import (
  "net"
)

func InitTCPServ() {
  listen, err := net.Listen("tcp", "127.0.0.1:3000")
  CheckErr(err)

  defer listen.Close()

  for {
    // Accepting new connections
    conn, err := listen.Accept()
    CheckErr(err)
    conn.Write([]byte("Welcome, new agent!"))

    // Handling new agent
    NewAgent(1, conn).Handle()
  }
}
