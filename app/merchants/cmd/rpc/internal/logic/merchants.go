package logic

import (
	"errors"
	"gorm.io/gorm"
	"mall-go/app/merchants/cmd/pb"
	"mall-go/app/merchants/cmd/rpc/internal/model"
	"mall-go/app/merchants/cmd/rpc/internal/svc"
	"time"
)

type MerchantsLogic struct {
}

func (l *MerchantsLogic) JoinMerchant(in *pb.JoinMerchantRequest) (bool, error) {
	// 查询是否存在
	var merchant model.MallMerchant
	db := svc.Context.Db
	err := db.First(&merchant, "mobile = ? AND user_id = ?", in.Mobile, in.UserId).Error
	if err == gorm.ErrRecordNotFound {
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
		if err = db.Create(&merchant).Error; err != nil {
			return false, err
		}
		return true, nil
	}

	if merchant.ID != 0 {
		return false, errors.New("已经申请入驻了")
	}

	return false, err
}

func (l *MerchantsLogic) UpdateMerchant(in *pb.UpdateMerchantRequest) (bool, error) {
	var merchant model.MallMerchant
	db := svc.Context.Db
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
	db := svc.Context.Db
	db.First(&merchant, "id = ?", in.ShopId)
	if merchant.ID == 0 {
		return false, errors.New("商户不存在")
	}
	merchant.Status = 5
	merchant.UpdatedAt = time.Now()
	if err := db.Save(&merchant).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (l *MerchantsLogic) GetMerchant(in *pb.GetMerchantRequest) (*model.MallMerchant, error) {
	var merchant model.MallMerchant
	db := svc.Context.Db
	if err := db.First(&merchant, "user_id = ?", in.UserId).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}
