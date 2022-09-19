package model

import "time"

// MallPost 文章表
type MallPost struct {
	ID              int64     `db:"id"`
	UserId          int64     `db:"user_id"`           // 用户ID
	CommentCount    int64     `db:"comment_count"`     // 评论数
	UpvoteCount     int64     `db:"upvote_count"`      // 点赞数
	Visibility      int8      `db:"visibility"`        // 可见性 0公开 1私密 2好友可见
	IsTop           int8      `db:"is_top"`            // 是否置顶
	IsEssence       int8      `db:"is_essence"`        // 是否精华
	IsLock          int8      `db:"is_lock"`           // 是否锁定
	LatestRepliedOn int64     `db:"latest_replied_on"` // 最后回复时间
	Ip              string    `db:"ip"`                // IP地址
	IpLoc           string    `db:"ip_loc"`            // IP城市
	CreatedAt       time.Time `db:"created_at"`        // 创建时间
	UpdatedAt       time.Time `db:"updated_at"`        // 更新时间
	IsDelete        int8      `db:"is_delete"`         // 是否删除
	DeletedAt       time.Time `db:"deleted_at"`        // 删除时间
}

// TableName 表名称
func (MallPost) TableName() string {
	return "mall_post"
}
