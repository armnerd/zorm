package db

// 组装sql
func (m *Model) getSqlForUpdate(table table) {
	sql := "update "
	m.Sql = sql
}

// Set
func (m *Model) Set(field map[string]interface{}) *Model {
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
