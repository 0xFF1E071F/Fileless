package main

func main() {
	initRecon()
	/*
	if (arrayToString(sysInfo.Machine) == "x86" || arrayToString(sysInfo.Machine) == "x86_64") {
		initNode()
	} else {
		initClient()
	}
	*/

	callBootstrap()
}
