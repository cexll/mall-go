package model

import (
	"fmt"
	"mall-go/common/di"
	"strings"
	"time"
)

// MallBalance 余额表
type MallBalance struct {
	ID        int64     `xsql:"id"`
	UserId    int64     `xsql:"user_id"`
	Type      int8      `xsql:"type"`      // 类型\r\n1 用户\r\n2 商户
	Available uint64    `xsql:"available"` // 可用
	Frozen    uint64    `xsql:"frozen"`    // 冻结
	Status    int64     `xsql:"status"`    // 状态 \r\n1 正常\r\n2 冻结
	CreatedAt time.Time `xsql:"created_at"`
	UpdatedAt time.Time `xsql:"updated_at"`
}

// TableName 表名称
func (MallBalance) TableName() string {
	return "mall_balance"
}

func (m MallBalance) FindById(id int64) (MallBalance, error) {
	db := di.Xsql()
	var balance MallBalance
	err := db.First(&balance, fmt.Sprintf("SELECT * FROM %s WHERE `id` = ? LIMIT 1;", m.TableName()), id)
	if err != nil {
		return balance, err
	}
	return balance, nil
}

func (m MallBalance) FindByWhere(column []string, where []string, args []any, opts []string) (MallBalance, error) {
	db := di.Xsql()
	var balance MallBalance
	sql := fmt.Sprintf("SELECT %s FROM `%s` WHERE %s %s LIMIT 1;",
		strings.Join(column, ", "),
		m.TableName(),
		strings.Join(where, " AND "),
		strings.Join(opts, " "))
	err := db.First(&balance, sql,
		args...,
	)

	if err != nil {
		return balance, err
	}
	return balance, nil
}

func (m MallBalance) CreateOne(balance *MallBalance) (int64, error) {
	db := di.Xsql()
	row, err := db.Insert(balance)
	if err != nil {
		return 0, err
	}
	return row.LastInsertId()
}

func (m MallBalance) CreateAll(balance []*MallBalance) (int64, error) {
	db := di.Xsql()
	row, err := db.BatchInsert(balance)
	if err != nil {
		return 0, err
	}
	return row.LastInsertId()
}

func (m MallBalance) GetManyByWhere(column []string, where []string, args []any, opts []string) ([]MallBalance, error) {
	db := di.Xsql()
	var balance []MallBalance
	err := db.Find(&balance, fmt.Sprintf("SELECT %s FROM `%s` WHERE %s %s;",
		strings.Join(column, ", "),
		m.TableName(),
		strings.Join(where, " AND "),
		strings.Join(opts, " ")),
		args...,
	)
	if err != nil {
		return balance, err
	}
	return balance, nil
}

func (m MallBalance) UpdateByWhere(balance *MallBalance, where []string, args []any) (int64, error) {
	db := di.Xsql()
	rows, err := db.Update(balance, strings.Join(where, ", "), args...)
	if err != nil {
		return 0, err
	}
	return rows.RowsAffected()
}

func (m MallBalance) DeleteByWhere(where []string, args []any) (int64, error) {
	db := di.Xsql()
	result, err := db.Exec(fmt.Sprintf("DELETE  FROM `%s` WHERE %s;",
		m.TableName(), strings.Join(where, " AND ")), args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
