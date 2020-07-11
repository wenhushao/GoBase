package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	userName  string = "root"
	password  string = "201606"
	ipAddrees string = "127.0.0.1"
	port      int    = 3306
	dbName    string = "test"
	charset   string = "utf8"
)

//连接数据库
func connectMysql() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return Db
}

//添加数据
func add(Db *sqlx.DB) {
	for i := 0; i < 2; i++ {
		result, err := Db.Exec("insert into user  values(?,?,?)", 2, "johny", "123456")
		if err != nil {
			fmt.Printf("data insert faied, error:[%v]", err.Error())
			return
		}
		id, _ := result.LastInsertId()
		fmt.Printf("insert success, last id:[%d]\n", id)
	}
}

//更新uid=2的username
func update(Db *sqlx.DB) {
	result, err := Db.Exec("update user set username = 'anson' where userid = 2")
	if err != nil {
		fmt.Printf("update faied, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected()
	fmt.Printf("update success, affected rows:[%d]\n", num)
}

//删除uid=2的数据
func delete(Db *sqlx.DB) {
	result, err := Db.Exec("delete from user where userid = 2")
	if err != nil {
		fmt.Printf("delete faied, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected()
	fmt.Printf("delete success, affected rows:[%d]\n", num)
}

//查询数据
func query(Db *sqlx.DB) {
	rows, err := Db.Query("select * from user")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	for rows.Next() {
		//定义变量接收查询数据
		var userid int
		var username, password string

		err := rows.Scan(&userid, &username, &password)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
		}
		fmt.Println(userid, username, password)
	}

	//关闭结果集（释放连接）go
	rows.Close()
}

func main() {
	var Db *sqlx.DB = connectMysql()
	defer Db.Close()
	//add(Db)
	//update(Db)
	//delete(Db)
	query(Db)
}
