package action

import (
	"database/sql"

	"github.com/armnerd/zorm/internal/element"
)

// 会话
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
