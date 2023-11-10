package test

import (
	"fmt"
	"testing"

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

func init() {
	config := db.Config{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Pass:     "123456",
		Database: "geek",
	}
	db.Connect(config)
}

func TestSearch(t *testing.T) {
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

func TestUpdate(t *testing.T) {
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

func TestAdd(t *testing.T) {
	session := db.Session{}
	fields := map[string]string{
		"name": "zane",
	}
	var Demo = Demo{}
	session.Field(fields).Save(Demo)
}

// 删除
func TestDelete(t *testing.T) {
	session := db.Session{}
	wheres := [][]string{
		{"id", "=", "1"},
	}
	var Demo = Demo{}
	session.Where(wheres).Delete(Demo)
}
