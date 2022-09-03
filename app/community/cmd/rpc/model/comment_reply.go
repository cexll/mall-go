package model

import "time"

// MallCommentReply 评论回复表
type MallCommentReply struct {
	ID        int64     `db:"id"`
	CommentId int64     `db:"comment_id"` // 评论ID
	UserId    int64     `db:"user_id"`    // 用户ID
	AtUserId  int64     `db:"at_user_id"` // @用户ID
	Content   string    `db:"content"`    // 内容
	Ip        string    `db:"ip"`         // IP地址
	IpLoc     string    `db:"ip_loc"`     // IP城市
	CreatedAt time.Time `db:"created_at"` // 创建时间
	UpdatedAt time.Time `db:"updated_at"` // 更新时间
	IsDelete  int8      `db:"is_delete"`  // 是否删除
	DeletedAt time.Time `db:"deleted_at"` // 删除时间
}

// TableName 表名称
func (MallCommentReply) TableName() string {
	return "mall_comment_reply"
}
