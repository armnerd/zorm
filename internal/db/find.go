package db

import (
	"fmt"
	"reflect"
	"strings"
)

// field
func (m *Model) Select(field []string) *Model {
	m.SelectSet = append(m.SelectSet, field...)
	return m
}

// condition
func (m *Model) Where(conditions [][]string) *Model {
	for index := range conditions {
		one := conditions[index]
		condition := whereEle{
			column:    one[0],
			condition: one[1],
			value:     one[2],
		}
		m.WhereSet = append(m.WhereSet, condition)
	}
	return m
}

// 组装sql
func (m *Model) getSqlForSelect(table table) {
	sql := "select "
	sql += strings.Join(m.SelectSet, ", ")
	wheres := make([]string, 0)
	for k := range m.WhereSet {
		v := m.WhereSet[k]
		temp := fmt.Sprintf("`%s` %s '%s'", v.column, v.condition, v.value)
		wheres = append(wheres, temp)
	}
	tableName := table.TableName()
	sql += " from " + tableName + " where "
	sql += strings.Join(wheres, " and ")
	m.Sql = sql
}

// 清空选项
func (m *Model) cleanUpForSelect() {
	m.Sql = ""
	m.SelectSet = []string{}
	m.WhereSet = []whereEle{}
	m.OrderSet = []orderEle{}
	m.Offset = 0
	m.Limit = 0
}

// find
func (m *Model) Find(table table) {
	m.getSqlForSelect(table)
	fmt.Println(m.Sql)
	// 执行单条查询
	rows, err := Db.Query(m.Sql)
	if err != nil {
		fmt.Println("多条数据错误", err)
		return
	}
	dest := reflect.ValueOf(table)
	destType := dest.Type()
	for rows.Next() {
		destRes := reflect.New(destType).Elem()
		var values []interface{}
		for i := 0; i < destRes.NumField(); i++ {
			values = append(values, destRes.Field(i).Addr().Interface())
		}
		rows.Scan(values...)
		fmt.Println("单行数据结果", destRes)
	}
	m.cleanUpForSelect()
}

// First
func (m *Model) First(table table) {
	m.getSqlForSelect(table)
	fmt.Println(m.Sql)
	// 执行多条查询
	dest := reflect.ValueOf(table)
	destType := dest.Type()
	destRes := reflect.New(destType).Elem()
	var values []interface{}
	for i := 0; i < destRes.NumField(); i++ {
		values = append(values, destRes.Field(i).Addr().Interface())
	}
	rows := Db.QueryRow(m.Sql)
	rows.Scan(values...)
	fmt.Println("单条数据结果：", destRes)
	m.cleanUpForSelect()
}
