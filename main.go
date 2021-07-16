package main

import "github.com/armnerd/zorm/internal"

type Demo struct {
	Id   int
	Name string
}

func (b Demo) TableName() string {
	return "demo"
}

func main() {
	model := internal.Model{}
	fields := []string{
		"id",
		"name",
	}
	wheres := [][]string{
		{"id", "=", "1"},
		{"name", "=", "zane"},
	}
	var Demo = Demo{}
	model.Select(fields).Where(wheres).Find(Demo)
}
