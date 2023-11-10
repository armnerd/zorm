package statement

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/armnerd/zorm/internal/element"
)

// field
func (s *Statement) Select(field []string) *Statement {
	s.SelectSet = append(s.SelectSet, field...)
	return s
}

// condition
func (s *Statement) Where(conditions [][]string) *Statement {
	for index := range conditions {
		one := conditions[index]
		condition := element.WhereEle{
			Column:    one[0],
			Condition: one[1],
			Value:     one[2],
		}
		s.WhereSet = append(s.WhereSet, condition)
	}
	return s
}

// order
func (s *Statement) Order(field map[string]string) *Statement {
	for k, v := range field {
		orderEle := element.OrderEle{
			Column:   k,
			Sequence: v,
		}
		s.OrderSet = append(s.OrderSet, orderEle)
	}
	return s
}

// offset
func (s *Statement) Offset(index int) *Statement {
	s.OffsetIndex = index
	return s
}

// limit
func (s *Statement) Limit(index int) *Statement {
	s.LimitIndex = index
	return s
}

// 组装sql
func (s *Statement) getSqlForSelect(table element.Table) {
	sql := "select "
	sql += strings.Join(s.SelectSet, ", ")
	wheres := make([]string, 0)
	for k := range s.WhereSet {
		v := s.WhereSet[k]
		temp := fmt.Sprintf("`%s` %s '%s'", v.Column, v.Condition, v.Value)
		wheres = append(wheres, temp)
	}
	tableName := table.TableName()
	sql += " from " + tableName + " where "
	sql += strings.Join(wheres, " and ")
	s.Sql = sql
}

// 清空选项
func (s *Statement) cleanUpForSelect() {
	s.Sql = ""
	s.SelectSet = []string{}
	s.WhereSet = []element.WhereEle{}
	s.OrderSet = []element.OrderEle{}
	s.OffsetIndex = 0
	s.LimitIndex = 0
}

// find
func (s *Statement) Find(dest interface{}) {
	destType := reflect.ValueOf(dest).Type().Elem().Elem()
	table, _ := reflect.New(destType).Elem().Interface().(element.Table)
	s.getSqlForSelect(table)
	fmt.Println(s.Sql)

	// 执行多条查询
	rows, err := s.Connection.Query(s.Sql)
	if err != nil {
		fmt.Println("多条数据错误", err)
		return
	}
	destValue := reflect.ValueOf(dest).Elem()
	for rows.Next() {
		destRes := reflect.New(destType).Elem()
		var values []interface{}
		for i := 0; i < destRes.NumField(); i++ {
			values = append(values, destRes.Field(i).Addr().Interface())
		}
		rows.Scan(values...)
		newArr := []reflect.Value{destRes}
		destValue.Set(reflect.Append(destValue, newArr...))
	}
	s.cleanUpForSelect()
}

// First
func (s *Statement) First(dest interface{}) {
	table, _ := dest.(element.Table)
	s.getSqlForSelect(table)
	s.Sql += " limit 1"
	fmt.Println(s.Sql)

	// 执行单条查询
	destRes := reflect.ValueOf(dest).Elem()
	var values []interface{}
	for i := 0; i < destRes.NumField(); i++ {
		values = append(values, destRes.Field(i).Addr().Interface())
	}
	row := s.Connection.QueryRow(s.Sql)
	row.Scan(values...)

	s.cleanUpForSelect()
}
