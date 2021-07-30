package main

import (
	"fmt"

	"github.com/armnerd/zorm/internal/db"
)

/*----------------------------示例表---------------------------------*/

type Demo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (d Demo) TableName() string {
	return "demo"
}

/*----------------------------示例表----------------------------------*/

func main() {
	config := db.Config{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Pass:     "123456",
		Database: "geek",
	}
	db.Connect(config)
	defer func() {
		db.Close()
	}()
	Search()
	Add()
	Update()
	Delete()
}

// 查询
func Search() {
	session := db.Session{}
	fields := []string{
		"id",
		"name",
	}
	wheres := [][]string{
		{"name", "=", "zane"},
	}
	var Demo = Demo{}

	// 多条
	resultList := session.Select(fields).Where(wheres).Find(Demo)
	fmt.Println(resultList)

	// 单条
	resultOne := session.Select(fields).Where(wheres).First(Demo)
	fmt.Println(resultOne)
}

// 新增
func Add() {
	session := db.Session{}
	fields := map[string]string{
		"name": "zane",
	}
	var Demo = Demo{}
	session.Field(fields).Save(Demo)
}

// 修改
func Update() {
	session := db.Session{}
	fields := map[string]string{
		"name": "frank",
	}
	wheres := [][]string{
		{"id", "=", "1"},
	}
	var Demo = Demo{}
	session.Set(fields).Where(wheres).Update(Demo)
}

// 删除
func Delete() {
	session := db.Session{}
	wheres := [][]string{
		{"id", "=", "1"},
	}
	var Demo = Demo{}
	session.Where(wheres).Delete(Demo)
}
