package main

import (
	"strings"
)

var servAddr = "127.0.0.1:8082"

func arrayToString(x [65]int8) string {
   var buf [65]byte
   for i, b := range x {
      buf[i] = byte(b)
   }
   str := string(buf[:])
   if i := strings.Index(str, "\x00"); i != -1 {
      str = str[:i]
   }
   return str
}
 
