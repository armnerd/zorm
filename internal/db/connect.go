package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var M Model

type Config struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func ConnectDB(config Config) (err error) {
	setup := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True"
	M = Model{}
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
