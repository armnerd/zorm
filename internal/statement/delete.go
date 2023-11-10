package action

import (
	"fmt"
	"strings"

	"github.com/armnerd/zorm/internal/element"
)

// 组装sql
func (s *Statement) getSqlForDelete(table element.Table) {
	tableName := table.TableName()
	sql := "delete from " + tableName
	wheres := make([]string, 0)
	for k := range s.WhereSet {
		v := s.WhereSet[k]
		temp := fmt.Sprintf("`%s` %s '%s'", v.Column, v.Condition, v.Value)
		wheres = append(wheres, temp)
	}
	sql += " where " + strings.Join(wheres, " and ")
	fmt.Println(sql)
	s.Sql = sql
}

// 清空选项
func (s *Statement) cleanUpForDelete() {
	s.Sql = ""
	s.WhereSet = []element.WhereEle{}
}

// delete
func (s *Statement) Delete(table element.Table) {
	s.getSqlForDelete(table)
	result, err := s.Connection.Exec(s.Sql)
	if err != nil {
		fmt.Println("删除数据错误", err)
		return
	}
	i, _ := result.RowsAffected()
	fmt.Printf("受影响行数：%d \n", i)
	s.cleanUpForDelete()
}
