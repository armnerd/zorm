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
	list := []Demo{}
	z.Query.Select(fields).Where(wheres).Find(&list)
	fmt.Println(list)

	// 单条
	one := Demo{}
	z.Query.Select(fields).Where(wheres).First(&one)
	fmt.Println(one)
}

func TestUpdate(t *testing.T) {
	fields := map[string]string{
		"name": "frank",
	}
	wheres := [][]string{
		{"id", "=", "1"},
	}
	z.Query.Set(fields).Where(wheres).Update(Demo{})
}

func TestAdd(t *testing.T) {
	fields := map[string]string{
		"name": "zane",
	}
	z.Query.Field(fields).Save(Demo{})
}

// 删除
func TestDelete(t *testing.T) {
	wheres := [][]string{
		{"id", "=", "1"},
	}
	z.Query.Where(wheres).Delete(Demo{})
}
