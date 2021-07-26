package db

import (
	"fmt"
	"strings"
)

// field
func (s *Session) Field(field map[string]string) *Session {
	for k, v := range field {
		fieldEle := insertEle{
			column: k,
			value:  v,
		}
		s.InsertSet = append(s.InsertSet, fieldEle)
	}
	return s
}

// 组装sql
func (s *Session) getSqlForInsert(table table) {
	tableName := table.TableName()
	sql := "insert into " + tableName
	keys := make([]string, 0)
	values := make([]string, 0)
	for k := range s.InsertSet {
		keys = append(keys, s.InsertSet[k].column)
		values = append(values, "'"+s.InsertSet[k].value+"'")
	}
	sql += fmt.Sprintf(" (%s)", strings.Join(keys, ","))
	sql += fmt.Sprintf(" VALUES (%s)", strings.Join(values, ","))
	fmt.Println(sql)
	s.Sql = sql
}

// 清空选项
func (s *Session) cleanUpForInsert() {
	s.Sql = ""
	s.InsertSet = []insertEle{}
}

// find
func (s *Session) Save(table table) {
	s.getSqlForInsert(table)
	result, err := Db.Exec(s.Sql)
	if err != nil {
		return
	}
	newID, _ := result.LastInsertId() // 新增数据的ID
	i, _ := result.RowsAffected()     // 受影响行数
	fmt.Printf("新增的数据ID：%d , 受影响行数：%d \n", newID, i)
	s.cleanUpForInsert()
}
