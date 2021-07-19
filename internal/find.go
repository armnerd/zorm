package internal

import (
	"fmt"
	"strings"
)

type Demo struct {
	Id   int
	Name string
}

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
func getSqlForSelect(m *Model, table table) {
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
	getSqlForSelect(m, table)
	//执行单条查询
	db, err := ConnectDB()
	if err != nil {
		fmt.Println("数据库链接错误", err)
		return
	}
	fmt.Println(m.Sql)
	rows, err := db.Query(m.Sql)
	if err != nil {
		fmt.Println("多条数据查询错误", err)
		return
	}
	var docList []Demo
	for rows.Next() {
		var doc Demo
		rows.Scan(&doc.Id, &doc.Name)
		//加入数组
		docList = append(docList, doc)
	}
	fmt.Println("多条数据查询结果", docList)
	m.cleanUpForSelect()
}

// First
func (m *Model) First(table table) {
	getSqlForSelect(m, table)
	var doc Demo
	//执行单条查询
	db, err := ConnectDB()
	if err != nil {
		fmt.Println("数据库链接错误", err)
		return
	}
	fmt.Println(m.Sql)
	rows := db.QueryRow(m.Sql)
	rows.Scan(&doc.Id, &doc.Name)
	fmt.Println("单条数据结果：", doc)
	m.cleanUpForSelect()
}
