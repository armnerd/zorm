package db

import (
	"fmt"
	"strings"
)

// 组装sql
func (m *Session) getSqlForDelete(table table) {
	tableName := table.TableName()
	sql := "delete from " + tableName
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

// 清空选项
func (m *Session) cleanUpForDelete() {
	m.Sql = ""
	m.WhereSet = []whereEle{}
}

// delete
func (m *Session) Delete(table table) {
	m.getSqlForDelete(table)
	result, err := Db.Exec(m.Sql)
	if err != nil {
		return
	}
	i, _ := result.RowsAffected()
	fmt.Printf("受影响行数：%d \n", i)
	m.cleanUpForDelete()
}
