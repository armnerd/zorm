package query

import (
	"fmt"
	"strings"

	"github.com/armnerd/zorm/internal/element"
)

// field
func (q *Query) Field(field map[string]string) *Query {
	for k, v := range field {
		fieldEle := element.InsertEle{
			Column: k,
			Value:  v,
		}
		q.InsertSet = append(q.InsertSet, fieldEle)
	}
	return q
}

// 组装sql
func (q *Query) getSqlForInsert(table element.Table) {
	tableName := table.TableName()
	sql := "insert into " + tableName
	keys := make([]string, 0)
	values := make([]string, 0)
	for k := range q.InsertSet {
		keys = append(keys, q.InsertSet[k].Column)
		values = append(values, "'"+q.InsertSet[k].Value+"'")
	}
	sql += fmt.Sprintf(" (%s)", strings.Join(keys, ","))
	sql += fmt.Sprintf(" VALUES (%s)", strings.Join(values, ","))
	fmt.Println(sql)
	q.Sql = sql
}

// 清空选项
func (q *Query) cleanUpForInsert() {
	q.Sql = ""
	q.InsertSet = []element.InsertEle{}
}

// save
func (q *Query) Save(table element.Table) {
	q.getSqlForInsert(table)
	result, err := q.Conn.Exec(q.Sql)
	if err != nil {
		fmt.Println("新增数据错误", err)
		return
	}
	newID, _ := result.LastInsertId() // 新增数据的ID
	i, _ := result.RowsAffected()     // 受影响行数
	fmt.Printf("新增的数据ID：%d , 受影响行数：%d \n", newID, i)
	q.cleanUpForInsert()
}
