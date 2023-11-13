package query

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/armnerd/zorm/internal/element"
)

// field
func (q *Query) Select(field []string) *Query {
	q.SelectSet = append(q.SelectSet, field...)
	return q
}

// condition
func (q *Query) Where(conditions [][]string) *Query {
	for index := range conditions {
		one := conditions[index]
		condition := element.WhereEle{
			Column:    one[0],
			Condition: one[1],
			Value:     one[2],
		}
		q.WhereSet = append(q.WhereSet, condition)
	}
	return q
}

// order
func (q *Query) Order(field map[string]string) *Query {
	for k, v := range field {
		orderEle := element.OrderEle{
			Column:   k,
			Sequence: v,
		}
		q.OrderSet = append(q.OrderSet, orderEle)
	}
	return q
}

// offset
func (q *Query) Offset(index int) *Query {
	q.OffsetIndex = index
	return q
}

// limit
func (q *Query) Limit(index int) *Query {
	q.LimitIndex = index
	return q
}

// 组装sql
func (q *Query) getSqlForSelect(table element.Table) {
	sql := "select "
	sql += strings.Join(q.SelectSet, ", ")
	wheres := make([]string, 0)
	for k := range q.WhereSet {
		v := q.WhereSet[k]
		temp := fmt.Sprintf("`%s` %s '%s'", v.Column, v.Condition, v.Value)
		wheres = append(wheres, temp)
	}
	tableName := table.TableName()
	sql += " from " + tableName + " where "
	sql += strings.Join(wheres, " and ")
	q.Sql = sql
}

// 清空选项
func (q *Query) cleanUpForSelect() {
	q.Sql = ""
	q.SelectSet = []string{}
	q.WhereSet = []element.WhereEle{}
	q.OrderSet = []element.OrderEle{}
	q.OffsetIndex = 0
	q.LimitIndex = 0
}

// find
func (q *Query) Find(dest interface{}) {
	destType := reflect.ValueOf(dest).Type().Elem().Elem()
	table, _ := reflect.New(destType).Elem().Interface().(element.Table)
	q.getSqlForSelect(table)
	fmt.Println(q.Sql)

	// 执行多条查询
	rows, err := q.Conn.Query(q.Sql)
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
	q.cleanUpForSelect()
}

// First
func (q *Query) First(dest interface{}) {
	table, _ := dest.(element.Table)
	q.getSqlForSelect(table)
	q.Sql += " limit 1"
	fmt.Println(q.Sql)

	// 执行单条查询
	destRes := reflect.ValueOf(dest).Elem()
	var values []interface{}
	for i := 0; i < destRes.NumField(); i++ {
		values = append(values, destRes.Field(i).Addr().Interface())
	}
	row := q.Conn.QueryRow(q.Sql)
	row.Scan(values...)

	q.cleanUpForSelect()
}
