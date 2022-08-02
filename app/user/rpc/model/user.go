package model

import (
	"fmt"
	"mall-go/common/di"
	"strings"
	"time"
)

// MallUser 用户表
type MallUser struct {
	ID        int64     `xsql:"id"`
	Nickname  string    `xsql:"nickname"`   // 用户名
	AvatarUrl string    `xsql:"avatar_url"` // 头像
	Password  string    `xsql:"password"`   // 密码
	Mobile    string    `xsql:"mobile"`     // 手机号
	Signature string    `xsql:"signature"`  // 个性签名
	Status    int8      `xsql:"status"`     // 状态 1正常 2禁用
	IsDelete  int8      `xsql:"is_delete"`  // 是否删除 1删除
	CreatedAt time.Time `xsql:"created_at"`
	UpdatedAt time.Time `xsql:"updated_at"`
}

// TableName 表名称
func (MallUser) TableName() string {
	return "mall_user"
}

func (m MallUser) FindByWhere(column []string, where []string, args []any, opts []string) (MallUser, error) {
	db := di.Xsql()
	var user MallUser
	err := db.First(&user, fmt.Sprintf("SELECT %s FROM `%s` WHERE %s %s",
		strings.Join(column, ", "),
		m.TableName(),
		strings.Join(where, " AND "),
		strings.Join(opts, " ")),
		args...,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (m MallUser) CreateOne(user *MallUser) (int64, error) {
	db := di.Xsql()
	row, err := db.Insert(user)
	if err != nil {
		return 0, err
	}
	return row.LastInsertId()
}
