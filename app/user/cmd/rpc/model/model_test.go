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

func TestMallUser_UpdateByWhere(t *testing.T) {
	user := MallUser{}

	// where update
	rows, err := user.UpdateByWhere([]string{
		"id = ?",
	}, []any{
		1,
	}, []string{
		"nickname = ?",
		"mobile = ?",
	}, []any{
		"张大仙",
		"19992399999",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rows)
	// 2022/08/05 17:40:41 SQL Time: 1.5781ms  Sql: UPDATE `mall_user` SET nickname = ?, mobile = ? WHERE id = ?;  Args: [张大仙 1999239999   9 %!s(int=1)]
}

func TestMallUser_DeleteByWhere(t *testing.T) {

}
