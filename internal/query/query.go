package query

import (
	"database/sql"

	"github.com/armnerd/zorm/internal/element"
)

// sql 语句
type Query struct {
	Conn        *sql.DB
	Sql         string
	SelectSet   []string
	WhereSet    []element.WhereEle
	OrderSet    []element.OrderEle
	OffsetIndex int
	LimitIndex  int
	InsertSet   []element.InsertEle
	UpdateSet   []element.UpdateEle
}
