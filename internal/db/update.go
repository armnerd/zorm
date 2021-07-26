package db

import (
	"fmt"
	"strings"
)

// Set
func (s *Session) Set(field map[string]string) *Session {
	for k, v := range field {
		updateEle := updateEle{
			column: k,
			value:  v,
		}
		s.UpdateSet = append(s.UpdateSet, updateEle)
	}
	return s
}

// 组装sql
func (s *Session) getSqlForUpdate(table table) {
	tableName := table.TableName()
	sql := "update " + tableName + " set "
	update := make([]string, 0)
	for k := range s.UpdateSet {
		v := s.UpdateSet[k]
		temp := fmt.Sprintf("%s='%s'", v.column, v.value)
		update = append(update, temp)
	}
	sql += strings.Join(update, ",")
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
func (s *Session) cleanUpForUpdate() {
	s.Sql = ""
	s.WhereSet = []whereEle{}
	s.UpdateSet = []updateEle{}
}

// update
func (s *Session) Update(table table) {
	s.getSqlForUpdate(table)
	result, err := Db.Exec(s.Sql)
	if err != nil {
		fmt.Println("修改数据错误", err)
		return
	}
	i, _ := result.RowsAffected() // 受影响行数
	fmt.Printf("受影响行数：%d \n", i)
	s.cleanUpForUpdate()
}
