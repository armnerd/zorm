package test

import (
	"fmt"
	"testing"

	"github.com/armnerd/zorm"
)

type Demo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (d Demo) TableName() string {
	return "demo"
}

var z *zorm.Zorm

func init() {
	z = zorm.NewZorm(
		zorm.WithHost("127.0.0.1"),
		zorm.WithPort("3306"),
		zorm.WithUser("root"),
		zorm.WithPass("123456"),
		zorm.WithDatabase("snail"),
	)
	z.Connect()
}

func TestSearch(t *testing.T) {
	fields := []string{
		"id",
		"name",
	}
	wheres := [][]string{
		{"name", "=", "zane"},
	}

	// 多条
	resultList := z.Statement.Select(fields).Where(wheres).Find(Demo{})
	fmt.Println(resultList)

	// 单条
	resultOne := z.Statement.Select(fields).Where(wheres).First(Demo{})
	fmt.Println(resultOne)
}

func TestUpdate(t *testing.T) {
	fields := map[string]string{
		"name": "frank",
	}
	wheres := [][]string{
		{"id", "=", "1"},
	}
	var Demo = Demo{}
	z.Statement.Set(fields).Where(wheres).Update(Demo)
}

func TestAdd(t *testing.T) {
	fields := map[string]string{
		"name": "zane",
	}
	var Demo = Demo{}
	z.Statement.Field(fields).Save(Demo)
}

// 删除
func TestDelete(t *testing.T) {
	wheres := [][]string{
		{"id", "=", "1"},
	}
	var Demo = Demo{}
	z.Statement.Where(wheres).Delete(Demo)
}
