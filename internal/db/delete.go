package db

// 组装sql
func (m *Session) getSqlForDelete(table table) {
	sql := "delete "
	m.Sql = sql
}

// 清空选项
func (m *Session) cleanUpForDelete() {
	m.Sql = ""
	m.WhereSet = []whereEle{}
}

// delete
func (m *Session) Delete(table table) {
	m.getSqlForDelete(table)
	m.cleanUpForDelete()
}
