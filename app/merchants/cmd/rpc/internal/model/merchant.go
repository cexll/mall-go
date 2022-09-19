package model

import "time"

// MallMerchant 商户表
type MallMerchant struct {
	ID        int64     `gorm:"id"`
	UserId    int64     `gorm:"user_id"`   // 用户ID
	ShopName  string    `gorm:"shop_name"` // 店铺名称
	ShopLogo  string    `gorm:"shop_logo"` // 店铺logo
	Mobile    string    `gorm:"mobile"`    // 店长手机
	Address   string    `gorm:"address"`   // 店铺地址
	Remark    string    `gorm:"remark"`    // 店铺描述
	Sort      int64     `gorm:"sort"`      // 店铺排序
	IsHide    int8      `gorm:"is_hide"`   // 是否隐藏
	Status    int8      `gorm:"status"`    // 店铺状态 \r\n1 正常\r\n2 禁用\r\n3 审核中\r\n4 审核拒绝 \r\n5 关闭店铺
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

// TableName 表名称
func (*MallMerchant) TableName() string {
	return "mall_merchant"
}
