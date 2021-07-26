package db

import (
	"fmt"
	"strings"
)

// 组装sql
func (s *Session) getSqlForDelete(table table) {
	tableName := table.TableName()
	sql := "delete from " + tableName
	wheres := make([]string, 0)
	for k := range s.WhereSet {
		v := s.WhereSet[k]
		temp := fmt.Sprintf("`%s` %s '%s'", v.column, v.condition, v.value)
		wheres = append(wheres, temp)
	}
	sql += " where " + strings.Join(wheres, " and ")
	fmt.Println(sql)
	s.Sql = sql
}

// 清空选项
func (s *Session) cleanUpForDelete() {
	s.Sql = ""
	s.WhereSet = []whereEle{}
}

// delete
func (s *Session) Delete(table table) {
	s.getSqlForDelete(table)
	result, err := Db.Exec(s.Sql)
	if err != nil {
		return
	}
	i, _ := result.RowsAffected()
	fmt.Printf("受影响行数：%d \n", i)
	s.cleanUpForDelete()
}
