package db

// 组装sql
func (m *Model) getSqlForInsert(table table) {
	sql := "insert info "
	m.Sql = sql
}

// field
func (m *Model) Field(field []string) *Model {
	m.SelectSet = append(m.SelectSet, field...)
	return m
}

// 清空选项
func (m *Model) cleanUpForInsert() {
	m.Sql = ""
	m.InsertSet = []insertEle{}
}

// find
func (m *Model) Insert(table table) {
	m.getSqlForInsert(table)
	m.cleanUpForInsert()
}
