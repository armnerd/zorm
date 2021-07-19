package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	var db *sql.DB
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/geek?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("数据库链接错误", err)
		return db, err
	}
	return db, err
}
