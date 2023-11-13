package query

import (
	"fmt"
	"strings"

	"github.com/armnerd/zorm/internal/element"
)

// Set
func (q *Query) Set(field map[string]string) *Query {
	for k, v := range field {
		updateEle := element.UpdateEle{
			Column: k,
			Value:  v,
		}
		q.UpdateSet = append(q.UpdateSet, updateEle)
	}
	return q
}

// 组装sql
func (q *Query) getSqlForUpdate(table element.Table) {
	tableName := table.TableName()
	sql := "update " + tableName + " set "
	update := make([]string, 0)
	for k := range q.UpdateSet {
		v := q.UpdateSet[k]
		temp := fmt.Sprintf("%s='%s'", v.Column, v.Value)
		update = append(update, temp)
	}
	sql += strings.Join(update, ",")
	wheres := make([]string, 0)
	for k := range q.WhereSet {
		v := q.WhereSet[k]
		temp := fmt.Sprintf("`%s` %s '%s'", v.Column, v.Condition, v.Value)
		wheres = append(wheres, temp)
	}
	sql += " where " + strings.Join(wheres, " and ")
	fmt.Println(sql)
	q.Sql = sql
}

// 清空选项
func (q *Query) cleanUpForUpdate() {
	q.Sql = ""
	q.WhereSet = []element.WhereEle{}
	q.UpdateSet = []element.UpdateEle{}
}

// update
func (q *Query) Update(table element.Table) {
	q.getSqlForUpdate(table)
	result, err := q.Conn.Exec(q.Sql)
	if err != nil {
		fmt.Println("修改数据错误", err)
		return
	}
	i, _ := result.RowsAffected() // 受影响行数
	fmt.Printf("受影响行数：%d \n", i)
	q.cleanUpForUpdate()
}
