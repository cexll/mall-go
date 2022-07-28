package model

import (
	"fmt"
	"github.com/cexll/mall-go/di"
	"github.com/mix-go/xsql"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	ID        int64  `gorm:"column:id" json:"id" form:"id" xsql:"id"`
	Nickname  string `gorm:"column:nickname" json:"nickname" form:"nickname" xsql:"nickname"`
	AvatarUrl string `gorm:"column:avatar_url" json:"avatar_url" form:"avatar_url" xsql:"avatar_url"`
	Password  string `gorm:"column:password" json:"password" form:"password" xsql:"password"`
	Mobile    string `gorm:"column:mobile" json:"mobile" form:"mobile" xsql:"mobile"`
	Signature string `gorm:"column:signature" json:"signature" form:"signature" xsql:"signature"`
	Status    int32  `gorm:"column:status" json:"status" form:"status" xsql:"status"`
	IsDelete  int32  `gorm:"column:is_delete" json:"is_delete" form:"is_delete" xsql:"is_delete"`
	CreatedAt string `gorm:"column:created_at" json:"created_at" form:"created_at" xsql:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at" form:"updated_at" xsql:"updated_at"`
}

const Table = "mall_user"

type UserModel struct {
}

func initialization() (db *gorm.DB) {
	return di.Gorm().Table("mall_user")
}

func (t *UserModel) FindById(Id int) *User {
	db := initialization()
	user := User{}
	db.First(&user, Id)
	return &user
}

func (t *UserModel) CreateOne(user *User) int64 {
	db := initialization()
	db.Create(&user)
	return user.ID
}

func CreateOne(data *any) (id int64, err error) {
	db := di.Xsql()
	row, err := db.Insert(data, xsql.Options{
		InsertKey: Table,
	})
	if err != nil {
		return 0, err
	}
	return row.LastInsertId()
}

func (t *UserModel) FindByWhere(column []string, where []string, args []any, opts []string) (User, error) {
	db := di.Xsql()
	var user User
	err := db.First(&user, fmt.Sprintf("SELECT %s FROM `%s` WHERE %s %s",
		strings.Join(column, ", "),
		Table,
		strings.Join(where, " AND "),
		strings.Join(opts, " ")),
		args...,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}
