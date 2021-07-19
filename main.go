package main

import (
	"github.com/armnerd/zorm/internal/db"
)

type Demo struct {
	Id   int
	Name string
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
	db.M.Select(fields).Where(wheres).Find(Demo)
	db.M.Select(fields).Where(wheres).First(Demo)
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
