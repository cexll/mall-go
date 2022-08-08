package di

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mix-go/xdi"
	"github.com/mix-go/xsql"
	"log"
	"mall-go/common/config"
)

func init() {
	obj := xdi.Object{
		Name: "xsql",
		New: func() (i interface{}, e error) {
			db, err := sql.Open("mysql", config.Conf.DB.DSN)
			if err != nil {
				return nil, err
			}
			return xsql.New(db, xsql.Options{
				DebugFunc: func(l *xsql.Log) {
					log.Printf("SQL Time: %s  Sql: %s  Args: %s", l.Time, l.SQL, l.Bindings)
				},
			}), nil
		},
	}
	if err := xdi.Provide(&obj); err != nil {
		panic(err)
	}
}

func Xsql() (db *xsql.DB) {
	if err := xdi.Populate("xsql", &db); err != nil {
		panic(err)
	}
	return
}
