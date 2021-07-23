package db

// 组装sql
func (m *Session) getSqlForInsert(table table) {
	sql := "insert info "
	m.Sql = sql
}

// field
func (m *Session) Field(field map[string]interface{}) *Session {
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
func (m *Session) cleanUpForInsert() {
	m.Sql = ""
	m.InsertSet = []insertEle{}
}

// find
func (m *Session) Save(table table) {
	m.getSqlForInsert(table)
	m.cleanUpForInsert()
}
