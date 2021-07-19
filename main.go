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
		{"name", "=", "zane"},
	}
	var Demo = Demo{}
	model.Select(fields).Where(wheres).Find(Demo)
	model.Select(fields).Where(wheres).First(Demo)
}
