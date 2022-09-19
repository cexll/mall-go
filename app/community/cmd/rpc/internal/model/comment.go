package model

import "time"

// MallComment 评论表
type MallComment struct {
	ID        int64     `db:"id"`
	PostId    int64     `db:"post_id"`    // 文章ID
	UserId    int64     `db:"user_id"`    // 用户ID
	Ip        string    `db:"ip"`         // IP地址
	IpLoc     string    `db:"ip_loc"`     // IP城市
	CreatedAt time.Time `db:"created_at"` // 创建时间
	UpdatedAt time.Time `db:"updated_at"` // 更新时间
	IsDelete  int8      `db:"is_delete"`  // 是否删除
	DeletedAt time.Time `db:"deleted_at"` // 删除时间
}

// TableName 表名称
func (MallComment) TableName() string {
	return "mall_comment"
}
