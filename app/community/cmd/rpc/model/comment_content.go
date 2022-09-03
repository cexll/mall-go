package model

import "time"

// MallCommentContent 文章内容表
type MallCommentContent struct {
	ID        int64     `db:"id"`
	CommentId int64     `db:"comment_id"` // 评论ID
	UserId    int64     `db:"user_id"`    // 用户ID
	Content   string    `db:"content"`    // 评论内容
	Type      int8      `db:"type"`       // 类型\r\n1. 标题\r\n2. 文章段落\r\n3. 图片地址\r\n4. 视频地址\r\n5. 语音地址\r\n6. 链接地址
	Sort      int64     `db:"sort"`       // 排序 越小越靠前
	CreatedAt time.Time `db:"created_at"` // 创建时间
	UpdatedAt time.Time `db:"updated_at"` // 更新时间
	IsDelete  int8      `db:"is_delete"`  // 是否删除
	DeletedAt time.Time `db:"deleted_at"` // 删除时间
}

// TableName 表名称
func (MallCommentContent) TableName() string {
	return "mall_comment_content"
}
