package pkg

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mall-go/common/conf"
)

func NewGorm(c conf.DbConf) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.DSN))
	if err != nil {
		panic(err)
	}
	return db
}
