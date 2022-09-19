package pkg

import (
	"database/sql"
	"github.com/mix-go/xsql"
	"github.com/zeromicro/go-zero/core/logx"
	"mall-go/common/conf"
)

func NewXsql(c conf.DbConf) *xsql.DB {
	db, err := sql.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}
	return xsql.New(db, xsql.Options{
		DebugFunc: func(l *xsql.Log) {
			logx.Infof("SQL Time: %s  Sql: %s  Args: %s", l.Time, l.SQL, l.Bindings)
		},
	})
}
