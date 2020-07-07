package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func insertData(db *sql.DB) {
	stmt, _ := db.Prepare(`INSERT INTO user (userid, username, password) VALUES (?, ?, ?)`)
	rows, err := stmt.Query(1, "Arthur", "A2020")
	defer stmt.Close()
	rows.Close()
	if err != nil {
		log.Fatalf("insert data error: %v\n", err)
	}
	rows, err = stmt.Query(2, "test", 19)
	var result int
	rows.Scan(&result)
	log.Printf("insert result %v\n", result)
	rows.Close()
}

func selectData(db *sql.DB) {
	var userid int
	var username string
	var password string
	rows, err := db.Query("select * from user")
	defer rows.Close()
	if err != nil {
		log.Fatalf("select data error: %v\n", err)
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
func deleteData(db *sql.DB) {
	stmt, _ := db.Prepare(`DELETE FROM user WHERE userid = ?`)

	rows, err := stmt.Query(1)
	defer stmt.Close()

	rows.Close()
	if err != nil {
		log.Fatalf("delete data error: %v\n", err)
	}

	rows, err = stmt.Query(2)
	rows.Close()
	if err != nil {
		log.Fatalf("delete data error: %v\n", err)
	}
}
func main() {
	db, err := sql.Open("mysql", "root:201606@tcp(127.0.0.1:3306)/test")

	defer db.Close()

	if err != nil {
		fmt.Printf("connect to db 127.0.0.1:3306 error: %v\n", err)
		return
	}

	//deleteData(db)

	//insertData(db)

	selectData(db)
}
