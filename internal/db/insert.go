package db

import (
	"fmt"
	"strings"
)

// field
func (m *Session) Field(field map[string]string) *Session {
	for k, v := range field {
		fieldEle := insertEle{
			column: k,
			value:  v,
		}
		m.InsertSet = append(m.InsertSet, fieldEle)
	}
	return m
}

// 组装sql
func (m *Session) getSqlForInsert(table table) {
	tableName := table.TableName()
	sql := "insert into " + tableName
	keys := make([]string, 0)
	values := make([]string, 0)
	for k := range m.InsertSet {
		keys = append(keys, m.InsertSet[k].column)
		values = append(values, "'"+m.InsertSet[k].value+"'")
	}
	sql += fmt.Sprintf(" (%s)", strings.Join(keys, ","))
	sql += fmt.Sprintf(" VALUES (%s)", strings.Join(values, ","))
	fmt.Println(sql)
	m.Sql = sql
}

// 清空选项
func (m *Session) cleanUpForInsert() {
	m.Sql = ""
	m.InsertSet = []insertEle{}
}

// find
func (m *Session) Save(table table) {
	m.getSqlForInsert(table)
	result, err := Db.Exec(m.Sql)
	if err != nil {
		return
	}
	newID, _ := result.LastInsertId() // 新增数据的ID
	i, _ := result.RowsAffected()     // 受影响行数
	fmt.Printf("新增的数据ID：%d , 受影响行数：%d \n", newID, i)
	m.cleanUpForInsert()
}
