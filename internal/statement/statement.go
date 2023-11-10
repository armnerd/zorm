package statement

import (
	"database/sql"

	"github.com/armnerd/zorm/internal/element"
)

// sql 语句
type Statement struct {
	Connection  *sql.DB
	Sql         string
	SelectSet   []string
	WhereSet    []element.WhereEle
	OrderSet    []element.OrderEle
	OffsetIndex int
	LimitIndex  int
	InsertSet   []element.InsertEle
	UpdateSet   []element.UpdateEle
}
