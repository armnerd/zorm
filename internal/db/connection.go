package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type Config struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

// 连接数据库
func Connect(config Config) (err error) {
	setup := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True"
	Db, err = sql.Open("mysql", fmt.Sprintf(setup, config.User, config.Pass, config.Host, config.Port, config.Database))
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

// 关闭连接
func Close() {
	Db.Close()
}
