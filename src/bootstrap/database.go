package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type database struct {
	db *sql.DB
}

type nodesTable struct {
	id int
	IP string
}

func initDatabase(user, password, IP, dbName string) *database {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, IP, dbName))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Println("[INFO] - Database connection established!")
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

func (handle *database) getNode() string {
	results, err := handle.db.Query("SELECT * FROM nodes")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var node nodesTable

	for results.Next() {
		err = results.Scan(&node.id, &node.IP)
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}

	return node.IP
}
