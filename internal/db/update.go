package db

// 组装sql
func (m *Session) getSqlForUpdate(table table) {
	sql := "update "
	m.Sql = sql
}

// Set
func (m *Session) Set(field map[string]interface{}) *Session {
	for k, v := range field {
		updateEle := updateEle{
			column: k,
			value:  v,
		}
		m.UpdateSet = append(m.UpdateSet, updateEle)
	}
	return m
}

// 清空选项
func (m *Session) cleanUpForUpdate() {
	m.Sql = ""
	m.WhereSet = []whereEle{}
	m.UpdateSet = []updateEle{}
}

// update
func (m *Session) Update(table table) {
	m.getSqlForUpdate(table)
	m.cleanUpForUpdate()
}
