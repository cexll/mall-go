package model

import "time"

// MallPostCollection 文章收藏表
type MallPostCollection struct {
	ID        int64     `db:"id"`
	PostId    int64     `db:"post_id"`    // 文章ID
	UserId    int64     `db:"user_id"`    // 用户ID
	CreatedAt time.Time `db:"created_at"` // 创建时间
	UpdatedAt time.Time `db:"updated_at"` // 更新时间
	IsDelete  int8      `db:"is_delete"`  // 是否删除
	DeletedAt time.Time `db:"deleted_at"` // 删除时间
}

// TableName 表名称
func (MallPostCollection) TableName() string {
	return "mall_post_collection"
}
