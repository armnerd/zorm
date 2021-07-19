package main

import (
	"fmt"

	"github.com/armnerd/zorm/internal/db"
)

type Demo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (b Demo) TableName() string {
	return "demo"
}

func main() {
	db.ConnectDB()
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

}

// 修改
func Update() {

}

// 删除
func Delete() {

}
