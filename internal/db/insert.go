package db

// 组装sql
func (m *Model) getSqlForInsert(table table) {
	sql := "insert info "
	m.Sql = sql
}

// field
func (m *Model) Field(field map[string]interface{}) *Model {
	for k, v := range field {
		fieldEle := insertEle{
			column: k,
			value:  v,
		}
		m.InsertSet = append(m.InsertSet, fieldEle)
	}
	return m
}

// 清空选项
func (m *Model) cleanUpForInsert() {
	m.Sql = ""
	m.InsertSet = []insertEle{}
}

// find
func (m *Model) Save(table table) {
	m.getSqlForInsert(table)
	m.cleanUpForInsert()
}
