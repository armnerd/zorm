package query

// 开始事务
func (q *Query) Begin() {
	tx, _ := q.Conn.Begin()
	q.tx = tx
}

// 事务回滚
func (q *Query) Rollback() {
	q.tx.Rollback()
}

// 事务提交
func (q *Query) Commit() {
	q.tx.Commit()
}
