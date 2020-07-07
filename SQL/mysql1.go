package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func insertData(db *sql.DB) {
	rows, err := db.Query(`INSERT INTO user (userid, username, password) VALUES (3, "xys", "123456")`)
	defer rows.Close()

	if err != nil {
		log.Fatalf("insert data error: %v\n", err)
	}

	var result int
	rows.Scan(&result)
	log.Printf("insert result %v\n", result)
}

func selectData(db *sql.DB) {
	var userid int
	var username string
	var password string
	rows, err := db.Query(`SELECT * From user where userid = 1`)
	if err != nil {
		log.Fatalf("insert data error: %v\n", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&userid, &username, &password)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("get data, userid: %d, username: %s, password: %s", userid, username, password)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:201606@tcp(127.0.0.1:3306)/test")

	defer db.Close()

	if err != nil {
		fmt.Printf("connect to db 127.0.0.1:3306 error: %v\n", err)
		return
	}

	insertData(db)

	selectData(db)
}
