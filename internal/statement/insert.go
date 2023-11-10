package action

import (
	"fmt"
	"strings"

	"github.com/armnerd/zorm/internal/element"
)

// field
func (s *Statement) Field(field map[string]string) *Statement {
	for k, v := range field {
		fieldEle := element.InsertEle{
			Column: k,
			Value:  v,
		}
		s.InsertSet = append(s.InsertSet, fieldEle)
	}
	return s
}

// 组装sql
func (s *Statement) getSqlForInsert(table element.Table) {
	tableName := table.TableName()
	sql := "insert into " + tableName
	keys := make([]string, 0)
	values := make([]string, 0)
	for k := range s.InsertSet {
		keys = append(keys, s.InsertSet[k].Column)
		values = append(values, "'"+s.InsertSet[k].Value+"'")
	}
	sql += fmt.Sprintf(" (%s)", strings.Join(keys, ","))
	sql += fmt.Sprintf(" VALUES (%s)", strings.Join(values, ","))
	fmt.Println(sql)
	s.Sql = sql
}

// 清空选项
func (s *Statement) cleanUpForInsert() {
	s.Sql = ""
	s.InsertSet = []element.InsertEle{}
}

// find
func (s *Statement) Save(table element.Table) {
	s.getSqlForInsert(table)
	result, err := s.Connection.Exec(s.Sql)
	if err != nil {
		fmt.Println("新增数据错误", err)
		return
	}
	newID, _ := result.LastInsertId() // 新增数据的ID
	i, _ := result.RowsAffected()     // 受影响行数
	fmt.Printf("新增的数据ID：%d , 受影响行数：%d \n", newID, i)
	s.cleanUpForInsert()
}
