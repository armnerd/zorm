package internal

import (
	"fmt"
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

// find
func (m *Model) Find(table table) {
	sql := "select "
	sql += strings.Join(m.SelectSet, ", ")
	wheres := make([]string, 0)
	for k := range m.WhereSet {
		v := m.WhereSet[k]
		temp := fmt.Sprintf("%s %s %s", v.column, v.condition, v.value)
		wheres = append(wheres, temp)
	}
	tableName := table.TableName()
	sql += " from " + tableName + " where "
	sql += strings.Join(wheres, " and ")
	m.Sql = sql
	fmt.Println(sql)
}
