package main


var db *database;

func main() {
	// Initializing the webserver
	db = initDatabase("bootstrap", "bootstrap", "127.0.0.1:3306", "bootstrap")
	startServ()
	
}
