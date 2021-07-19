package db

// 组装sql
func (m *Model) getSqlForUpdate(table table) {
	sql := "update "
	m.Sql = sql
}

// field
func (m *Model) Set(field []string) *Model {
	m.SelectSet = append(m.SelectSet, field...)
	return m
}

// 清空选项
func (m *Model) cleanUpForUpdate() {
	m.Sql = ""
	m.WhereSet = []whereEle{}
	m.UpdateSet = []updateEle{}
}

// update
func (m *Model) Update(table table) {
	m.getSqlForUpdate(table)
	m.cleanUpForUpdate()
}
