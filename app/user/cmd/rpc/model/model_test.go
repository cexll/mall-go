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
	// 一系列判断

	//user.Nickname = "张大仙"
	//update := MallUser{
	//	Nickname: "张大仙",
	//	Mobile:   "8008208820",
	//	Status:   1,
	//}

	//value := reflect.ValueOf(user)
	//key := reflect.TypeOf(user)
	//count := value.NumField()
	//fmt.Println("value is count ", count)
	//var set []string
	//for i := 0; i < count; i++ {
	//	f := value.Field(i)
	//	switch f.Kind() {
	//	case reflect.String:
	//		if f.String() != "" {
	//			set = append(set, fmt.Sprintf("%s = %s", key.Field(i).Name, f.String()))
	//		}
	//	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	//		if f.Int() != 0 {
	//			set = append(set, fmt.Sprintf("%s = %d", key.Field(i).Name, f.Int()))
	//		}
	//	case reflect.Interface:
	//		if f.Interface() != nil {
	//			set = append(set, fmt.Sprintf("%s = %d", key.Field(i).Name, f.Interface()))
	//		}
	//	}
	//}

	//fmt.Println(set)
}

func TestMallUser_DeleteByWhere(t *testing.T) {

}
