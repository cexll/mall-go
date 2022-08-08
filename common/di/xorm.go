package di

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/mix-go/xdi"
	"mall-go/common/config"
	"xorm.io/xorm"
)

func init() {
	obj := xdi.Object{
		Name: "xorm",
		New: func() (i interface{}, e error) {
			return xorm.NewEngine("mysql", config.Conf.DB.DSN)
		},
	}
	if err := xdi.Provide(&obj); err != nil {
		panic(err)
	}
}

func Xorm() (db *xorm.Engine) {
	if err := xdi.Populate("xorm", &db); err != nil {
		panic(err)
	}
	return
}
