package action

import (
	"fmt"
	"strings"

	"github.com/armnerd/zorm/internal/element"
)

// Set
func (s *Statement) Set(field map[string]string) *Statement {
	for k, v := range field {
		updateEle := element.UpdateEle{
			Column: k,
			Value:  v,
		}
		s.UpdateSet = append(s.UpdateSet, updateEle)
	}
	return s
}

// 组装sql
func (s *Statement) getSqlForUpdate(table element.Table) {
	tableName := table.TableName()
	sql := "update " + tableName + " set "
	update := make([]string, 0)
	for k := range s.UpdateSet {
		v := s.UpdateSet[k]
		temp := fmt.Sprintf("%s='%s'", v.Column, v.Value)
		update = append(update, temp)
	}
	sql += strings.Join(update, ",")
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
func (s *Statement) cleanUpForUpdate() {
	s.Sql = ""
	s.WhereSet = []element.WhereEle{}
	s.UpdateSet = []element.UpdateEle{}
}

// update
func (s *Statement) Update(table element.Table) {
	s.getSqlForUpdate(table)
	result, err := s.Connection.Exec(s.Sql)
	if err != nil {
		fmt.Println("修改数据错误", err)
		return
	}
	i, _ := result.RowsAffected() // 受影响行数
	fmt.Printf("受影响行数：%d \n", i)
	s.cleanUpForUpdate()
}
