package db

// 组装sql
func (m *Model) getSqlForDelete(table table) {
	sql := "delete "
	m.Sql = sql
}

// 清空选项
func (m *Model) cleanUpForDelete() {
	m.Sql = ""
	m.WhereSet = []whereEle{}
}

// delete
func (m *Model) Delete(table table) {
	m.getSqlForDelete(table)
	m.cleanUpForDelete()
}
