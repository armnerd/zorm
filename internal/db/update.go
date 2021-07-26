package db

import (
	"fmt"
	"strings"
)

// 组装sql
func (m *Session) getSqlForUpdate(table table) {
	tableName := table.TableName()
	sql := "update " + tableName + " set "
	update := make([]string, 0)
	for k := range m.UpdateSet {
		v := m.UpdateSet[k]
		temp := fmt.Sprintf("%s='%s'", v.column, v.value)
		update = append(update, temp)
	}
	sql += strings.Join(update, ",")
	wheres := make([]string, 0)
	for k := range m.WhereSet {
		v := m.WhereSet[k]
		temp := fmt.Sprintf("`%s` %s '%s'", v.column, v.condition, v.value)
		wheres = append(wheres, temp)
	}
	sql += " where " + strings.Join(wheres, " and ")
	fmt.Println(sql)
	m.Sql = sql
}

// Set
func (m *Session) Set(field map[string]string) *Session {
	for k, v := range field {
		updateEle := updateEle{
			column: k,
			value:  v,
		}
		m.UpdateSet = append(m.UpdateSet, updateEle)
	}
	return m
}

// 清空选项
func (m *Session) cleanUpForUpdate() {
	m.Sql = ""
	m.WhereSet = []whereEle{}
	m.UpdateSet = []updateEle{}
}

// update
func (m *Session) Update(table table) {
	m.getSqlForUpdate(table)
	result, err := Db.Exec(m.Sql)
	if err != nil {
		return
	}
	i, _ := result.RowsAffected() // 受影响行数
	fmt.Printf("受影响行数：%d \n", i)
	m.cleanUpForUpdate()
}
