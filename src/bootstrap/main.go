package main

import (
	"fmt"
)

type Config struct {
	Webserver struct {
		IP string `json:IP`
		Port int `json:Port`
	}

	Database struct {
		IP string `json:IP`
		User string `json:User`
		Password string `json:Password`
		DbName string `json:DbName`
	}
}

var cfg Config = Config{}
var db *database;

func main() {
	if (parseConfig()) {
		fmt.Println("[INFO] - settings.json parsed successfully!")
	}

	db = initDatabase(cfg.Database.User, cfg.Database.Password, cfg.Database.IP, cfg.Database.DbName)
	startServ(cfg.Webserver.IP, cfg.Webserver.Port)
	
}
