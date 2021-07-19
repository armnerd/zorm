package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var M Model

func ConnectDB() (err error) {
	M = Model{}
	Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/geek?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("数据库链接错误", err)
		return err
	}
	err = Db.Ping()
	if err != nil {
		return err
	}
	return err
}
