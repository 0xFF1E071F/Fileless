package main

import (
	"log"
	sys "syscall"
)

var sysInfo = sys.Utsname{}

func initRecon() {
	//Gathering system information
	err := sys.Uname(&sysInfo)
	if err != nil {
		log.Print("[ERROR] - ", err)
		return
	}
}
