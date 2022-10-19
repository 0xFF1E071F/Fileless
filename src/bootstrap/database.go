package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type database struct {
	db *sql.DB
}

func initDatabase(user, password, IP, dbName string) *database {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, IP, dbName))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &database{db}
}

func (handle *database) insertNode(IP string) bool {
	stmt, err := handle.db.Prepare("INSERT INTO nodes(IP) VALUES(?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer stmt.Close()
	_, err = stmt.Exec(IP)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
