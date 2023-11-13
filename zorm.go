package zorm

import (
	"database/sql"
	"fmt"

	"github.com/armnerd/zorm/internal/query"
	_ "github.com/go-sql-driver/mysql"
)

type Zorm struct {
	opts  *Options
	Conn  *sql.DB
	Query query.Query
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
	Conn, err := sql.Open("mysql", fmt.Sprintf(setup, z.opts.User, z.opts.Pass, z.opts.Host, z.opts.Port, z.opts.Database))
	if err != nil {
		fmt.Println("数据库链接错误", err)
		return err
	}
	err = Conn.Ping()
	if err != nil {
		return err
	}
	z.Conn = Conn
	z.Query = query.Query{
		Conn: Conn,
	}
	return err
}

// 关闭连接
func (z *Zorm) Close() {
	z.Conn.Close()
}
