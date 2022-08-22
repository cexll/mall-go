package logic

import (
	"errors"
	"mall-go/app/merchants/cmd/pb"
	"mall-go/app/merchants/cmd/rpc/model"
	"mall-go/common/di"
	"time"
)

type MerchantsLogic struct {
}

func (l *MerchantsLogic) JoinMerchant(in *pb.JoinMerchantRequest) (bool, error) {
	// 查询是否存在
	var merchant model.MallMerchant
	db := di.Gorm()
	db.First(&merchant, "mobile = ?", in.Mobile)
	if merchant.ID != 0 {
		return false, errors.New("手机号已绑定商户")
	}

	merchant.UserId = in.UserId
	merchant.ShopName = in.ShopName
	merchant.ShopLogo = in.ShopLogo
	merchant.Mobile = in.Mobile
	address := in.Address
	merchant.Address = address
	merchant.Remark = in.Remark
	merchant.Sort = 1
	merchant.IsHide = 0
	merchant.Status = 3
	merchant.CreatedAt = time.Now()
	db.Create(&merchant)
	if merchant.ID != 0 {
		return true, nil
	}
	return false, db.Error
}

func (l *MerchantsLogic) UpdateMerchant(in *pb.UpdateMerchantRequest) (bool, error) {
	var merchant model.MallMerchant
	db := di.Gorm()
	db.First(&merchant, "id = ?", in.ShopId)
	if merchant.ID == 0 {
		return false, errors.New("商户不存在")
	}
	if merchant.Status != 1 {
		return false, errors.New("商户状态异常")
	}
	merchant.ID = in.ShopId
	merchant.UserId = in.UserId
	merchant.ShopName = in.ShopName
	merchant.ShopLogo = in.ShopLogo
	merchant.Mobile = in.Mobile
	merchant.Address = in.Address
	merchant.Remark = in.Remark
	merchant.UpdatedAt = time.Now()

	db.Save(&merchant)
	if db.Error != nil {
		return false, db.Error
	}
	return true, nil
}

func (l *MerchantsLogic) CloseMerchant(in *pb.CloseMerchantRequest) (bool, error) {
	var merchant model.MallMerchant
	db := di.Gorm()
	db.First(&merchant, "id = ?", in.ShopId)
	if merchant.ID == 0 {
		return false, errors.New("商户不存在")
	}
	merchant.Status = 5
	merchant.UpdatedAt = time.Now()
	db.Save(&merchant)

	if db.Error != nil {
		return false, db.Error
	}
	return true, nil
}
