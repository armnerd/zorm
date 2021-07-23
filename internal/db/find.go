package db

import (
	"fmt"
	"reflect"
	"strings"
)

// field
func (m *Session) Select(field []string) *Session {
	m.SelectSet = append(m.SelectSet, field...)
	return m
}

// condition
func (m *Session) Where(conditions [][]string) *Session {
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

// order
func (m *Session) Order(field map[string]string) *Session {
	for k, v := range field {
		orderEle := orderEle{
			column:   k,
			sequence: v,
		}
		m.OrderSet = append(m.OrderSet, orderEle)
	}
	return m
}

// offset
func (m *Session) Offset(index int) *Session {
	m.OffsetIndex = index
	return m
}

// limit
func (m *Session) Limit(index int) *Session {
	m.LimitIndex = index
	return m
}

// 组装sql
func (m *Session) getSqlForSelect(table table) {
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
func (m *Session) cleanUpForSelect() {
	m.Sql = ""
	m.SelectSet = []string{}
	m.WhereSet = []whereEle{}
	m.OrderSet = []orderEle{}
	m.OffsetIndex = 0
	m.LimitIndex = 0
}

// find
func (m *Session) Find(table table) []map[string]interface{} {
	res := make([]map[string]interface{}, 0)
	m.getSqlForSelect(table)
	fmt.Println(m.Sql)

	// 执行单条查询
	rows, err := Db.Query(m.Sql)
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
	m.cleanUpForSelect()
	return res
}

// First
func (m *Session) First(table table) map[string]interface{} {
	res := make(map[string]interface{})
	m.getSqlForSelect(table)
	fmt.Println(m.Sql)

	// 执行多条查询
	dest := reflect.ValueOf(table)
	destInfo := reflect.TypeOf(table)
	destType := dest.Type()
	destRes := reflect.New(destType).Elem()
	var values []interface{}
	for i := 0; i < destRes.NumField(); i++ {
		values = append(values, destRes.Field(i).Addr().Interface())
	}
	rows := Db.QueryRow(m.Sql)
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

	m.cleanUpForSelect()
	return res
}
