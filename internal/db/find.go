package db

import (
	"fmt"
	"reflect"
	"strings"
)

// field
func (s *Session) Select(field []string) *Session {
	s.SelectSet = append(s.SelectSet, field...)
	return s
}

// condition
func (s *Session) Where(conditions [][]string) *Session {
	for index := range conditions {
		one := conditions[index]
		condition := whereEle{
			column:    one[0],
			condition: one[1],
			value:     one[2],
		}
		s.WhereSet = append(s.WhereSet, condition)
	}
	return s
}

// order
func (s *Session) Order(field map[string]string) *Session {
	for k, v := range field {
		orderEle := orderEle{
			column:   k,
			sequence: v,
		}
		s.OrderSet = append(s.OrderSet, orderEle)
	}
	return s
}

// offset
func (s *Session) Offset(index int) *Session {
	s.OffsetIndex = index
	return s
}

// limit
func (s *Session) Limit(index int) *Session {
	s.LimitIndex = index
	return s
}

// 组装sql
func (s *Session) getSqlForSelect(table table) {
	sql := "select "
	sql += strings.Join(s.SelectSet, ", ")
	wheres := make([]string, 0)
	for k := range s.WhereSet {
		v := s.WhereSet[k]
		temp := fmt.Sprintf("`%s` %s '%s'", v.column, v.condition, v.value)
		wheres = append(wheres, temp)
	}
	tableName := table.TableName()
	sql += " from " + tableName + " where "
	sql += strings.Join(wheres, " and ")
	s.Sql = sql
}

// 清空选项
func (s *Session) cleanUpForSelect() {
	s.Sql = ""
	s.SelectSet = []string{}
	s.WhereSet = []whereEle{}
	s.OrderSet = []orderEle{}
	s.OffsetIndex = 0
	s.LimitIndex = 0
}

// find
func (s *Session) Find(table table) []map[string]interface{} {
	res := make([]map[string]interface{}, 0)
	s.getSqlForSelect(table)
	fmt.Println(s.Sql)

	// 执行单条查询
	rows, err := Db.Query(s.Sql)
	if err != nil {
		fmt.Println("多条数据错误", err)
		return res
	}
	dest := reflect.ValueOf(table)
	destInfo := reflect.TypeOf(table)
	destType := dest.Type()
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
func (s *Session) First(table table) map[string]interface{} {
	res := make(map[string]interface{})
	s.getSqlForSelect(table)
	fmt.Println(s.Sql)

	// 执行多条查询
	dest := reflect.ValueOf(table)
	destInfo := reflect.TypeOf(table)
	destType := dest.Type()
	destRes := reflect.New(destType).Elem()
	var values []interface{}
	for i := 0; i < destRes.NumField(); i++ {
		values = append(values, destRes.Field(i).Addr().Interface())
	}
	rows := Db.QueryRow(s.Sql)
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
