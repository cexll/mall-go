package model

import (
	"fmt"
	"mall-go/app/user/cmd/rpc/internal/svc"
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

func (m *MallUser) FindById(id int64) (*MallUser, error) {
	db := svc.Context.Xsql
	var user MallUser
	err := db.First(&user, fmt.Sprintf("SELECT * FROM %s WHERE `id` = ? LIMIT 1;", m.TableName()), id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *MallUser) FindByWhere(column []string, where []string, args []any, opts []string) (*MallUser, error) {
	db := svc.Context.Xsql
	var user MallUser
	sql := fmt.Sprintf("SELECT %s FROM `%s` WHERE %s %s LIMIT 1;",
		strings.Join(column, ", "),
		m.TableName(),
		strings.Join(where, " AND "),
		strings.Join(opts, " "))
	err := db.First(&user, sql,
		args...,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *MallUser) CreateOne(user *MallUser) (int64, error) {
	db := svc.Context.Xsql
	row, err := db.Insert(user)
	if err != nil {
		return 0, err
	}
	return row.LastInsertId()
}

func (m *MallUser) CreateAll(user []*MallUser) (int64, error) {
	db := svc.Context.Xsql
	row, err := db.BatchInsert(user)
	if err != nil {
		return 0, err
	}
	return row.LastInsertId()
}

func (m *MallUser) GetManyByWhere(column []string, where []string, args []any, opts []string) ([]*MallUser, error) {
	db := svc.Context.Xsql
	var user []*MallUser
	err := db.Find(&user, fmt.Sprintf("SELECT %s FROM `%s` WHERE %s %s;",
		strings.Join(column, ", "),
		m.TableName(),
		strings.Join(where, " AND "),
		strings.Join(opts, " ")),
		args...,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *MallUser) UpdateByWhere(user *MallUser, where []string, args []any) (int64, error) {
	db := svc.Context.Xsql
	rows, err := db.Update(user, strings.Join(where, ", "), args...)
	if err != nil {
		return 0, err
	}
	return rows.RowsAffected()
}

func (m *MallUser) DeleteByWhere(where []string, args []any) (int64, error) {
	db := svc.Context.Xsql
	result, err := db.Exec(fmt.Sprintf("DELETE  FROM `%s` WHERE %s;",
		m.TableName(), strings.Join(where, " AND ")), args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
