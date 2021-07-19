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

func (b Demo) TableName() string {
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
	db.ConnectDB(config)
	Search()
	Add()
	Update()
	Delete()
}

// 查询
func Search() {
	fields := []string{
		"id",
		"name",
	}
	wheres := [][]string{
		{"name", "=", "zane"},
	}
	var Demo = Demo{}

	// 查询单个
	resultList := db.M.Select(fields).Where(wheres).Find(Demo)
	fmt.Println(resultList)

	// 查询多个
	resultOne := db.M.Select(fields).Where(wheres).First(Demo)
	fmt.Println(resultOne)
}

// 新增
func Add() {
	fields := map[string]interface{}{
		"name": "zane",
	}
	var Demo = Demo{}
	db.M.Field(fields).Save(Demo)
}

// 修改
func Update() {
	fields := map[string]interface{}{
		"name": "frank",
	}
	wheres := [][]string{
		{"id", "=", "1"},
	}
	var Demo = Demo{}
	db.M.Set(fields).Where(wheres).Update(Demo)
}

// 删除
func Delete() {
	wheres := [][]string{
		{"id", "=", "1"},
	}
	var Demo = Demo{}
	db.M.Where(wheres).Delete(Demo)
}
