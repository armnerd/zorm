package query

import (
	"fmt"
	"strings"

	"github.com/armnerd/zorm/internal/element"
)

// 组装sql
func (q *Query) getSqlForDelete(table element.Table) {
	tableName := table.TableName()
	sql := "delete from " + tableName
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
func (q *Query) cleanUpForDelete() {
	q.Sql = ""
	q.WhereSet = []element.WhereEle{}
}

// delete
func (q *Query) Delete(table element.Table) {
	q.getSqlForDelete(table)
	result, err := q.Conn.Exec(q.Sql)
	if err != nil {
		fmt.Println("删除数据错误", err)
		return
	}
	i, _ := result.RowsAffected()
	fmt.Printf("受影响行数：%d \n", i)
	q.cleanUpForDelete()
}
