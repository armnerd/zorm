package zorm

import (
	"database/sql"
	"fmt"

	"github.com/armnerd/zorm/internal/statement"
	_ "github.com/go-sql-driver/mysql"
)

type Zorm struct {
	opts       *Options
	Connection *sql.DB
	Statement  statement.Statement
}

func NewZorm(opts ...OptionFunc) *Zorm {
	options := loadOptions(opts...)
	z := &Zorm{
		opts: options,
	}
	return z
}

func (z *Zorm) Connect() (err error) {
	setup := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True"
	Connection, err := sql.Open("mysql", fmt.Sprintf(setup, z.opts.User, z.opts.Pass, z.opts.Host, z.opts.Port, z.opts.Database))
	if err != nil {
		fmt.Println("数据库链接错误", err)
		return err
	}
	err = Connection.Ping()
	if err != nil {
		return err
	}
	z.Connection = Connection
	z.Statement = statement.Statement{
		Connection: Connection,
	}
	return err
}

// 关闭连接
func (z *Zorm) Close() {
	z.Connection.Close()
}
