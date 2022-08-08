package model

import (
	"fmt"
	"mall-go/common/di"
	"testing"
	"time"
)

func TestMallUser_CreateOne(t *testing.T) {
	db := di.Xsql()
	user := &MallUser{
		Nickname:  "cexll",
		Password:  "123123",
		Mobile:    "18882388888",
		Status:    1,
		IsDelete:  0,
		CreatedAt: time.Now(),
	}

	row, err := db.Insert(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(row.LastInsertId())
}

func TestMallUser_FindByWhere(t *testing.T) {
	user := MallUser{}

	user, err := user.FindByWhere([]string{
		"id", "nickname", "password", "mobile", "status", "is_delete",
	}, []string{
		"id = ?", "is_delete = ?",
	}, []any{
		1, 0,
	}, []string{
		"order by id",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
}
