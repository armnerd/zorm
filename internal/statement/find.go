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
func (s *Statement) Find(table element.Table) []map[string]interface{} {
	res := make([]map[string]interface{}, 0)
	s.getSqlForSelect(table)
	fmt.Println(s.Sql)

	// 执行单条查询
	rows, err := s.Connection.Query(s.Sql)
	if err != nil {
		fmt.Println("多条数据错误", err)
		return res
	}
	destType := reflect.ValueOf(table).Type()
	destInfo := reflect.TypeOf(table)
	for rows.Next() {
		one := make(map[string]interface{})
		destRes := reflect.New(destType).Elem()
		var values []interface{}
		for i := 0; i < destRes.NumField(); i++ {
			values = append(values, destRes.Field(i).Addr().Interface())
		}
		rows.Scan(values...)
		for i := 0; i < destRes.NumField(); i++ {
			key := destInfo.Field(i).Tag.Get("json")
			kind := destRes.Field(i).Kind().String()
			var value interface{}
			switch kind {
			case "int":
				value = destRes.Field(i).Interface().(int)
			case "string":
				value = destRes.Field(i).Interface().(string)

			}
			one[key] = value
		}
		res = append(res, one)
	}
	s.cleanUpForSelect()
	return res
}

// First
func (s *Statement) First(table element.Table) map[string]interface{} {
	res := make(map[string]interface{})
	s.getSqlForSelect(table)
	s.Sql += " limit 1"
	fmt.Println(s.Sql)

	// 执行单条查询
	destType := reflect.ValueOf(table).Type()
	destInfo := reflect.TypeOf(table)
	destRes := reflect.New(destType).Elem()
	var values []interface{}
	for i := 0; i < destRes.NumField(); i++ {
		values = append(values, destRes.Field(i).Addr().Interface())
	}
	rows := s.Connection.QueryRow(s.Sql)
	rows.Scan(values...)
	for i := 0; i < destRes.NumField(); i++ {
		key := destInfo.Field(i).Tag.Get("json")
		kind := destRes.Field(i).Kind().String()
		var value interface{}
		switch kind {
		case "int":
			value = destRes.Field(i).Interface().(int)
		case "string":
			value = destRes.Field(i).Interface().(string)

		}
		res[key] = value
	}

	s.cleanUpForSelect()
	return res
}
