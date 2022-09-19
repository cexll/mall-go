package controllers

import (
	"github.com/gin-gonic/gin"
	balancepb "mall-go/app/balance/cmd/pb"
	"mall-go/app/merchants/cmd/api/internal/logic"
	"mall-go/app/merchants/cmd/pb"
	"mall-go/common"
	"mall-go/common/response"
	"net/http"
)

type MerchantsController struct {
	resp  response.Response
	logic logic.MerchantsLogic
}

type JoinMerchantRequest struct {
	ShopName string `json:"shop_name" validate:"required"`
	ShopLogo string `json:"shop_logo" validate:"required"`
	Mobile   string `json:"mobile" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Remark   string `json:"remark" validate:"min=0"`
}

func (t *MerchantsController) JoinMerchant(c *gin.Context) {
	var req JoinMerchantRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}
	userId, err := common.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}
	resp, err := t.logic.JoinMerchant(&pb.JoinMerchantRequest{
		ShopName: req.ShopName,
		ShopLogo: req.ShopLogo,
		Mobile:   req.Mobile,
		Address:  req.Address,
		Remark:   req.Remark,
		UserId:   userId,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(resp))
}

type UpdateMerchantRequest struct {
	ShopId   int64  `json:"shop_id" validate:"required"`
	ShopName string `json:"shop_name" validate:"required"`
	ShopLogo string `json:"shop_logo" validate:"required"`
	Mobile   string `json:"mobile" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Remark   string `json:"remark" validate:"min=0"`
}

func (t *MerchantsController) UpdateMerchant(c *gin.Context) {
	var req UpdateMerchantRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}
	userId, err := common.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}
	resp, err := t.logic.UpdateMerchant(&pb.UpdateMerchantRequest{
		ShopName: req.ShopName,
		ShopLogo: req.ShopLogo,
		Mobile:   req.Mobile,
		Address:  req.Address,
		Remark:   req.Remark,
		UserId:   userId,
		ShopId:   req.ShopId,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(resp))
}

func (t *MerchantsController) CloseMerchant(c *gin.Context) {
	userId, err := common.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}
	merchant, err := t.logic.GetMerchant(&pb.GetMerchantRequest{
		UserId: userId,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}
	resp, err := t.logic.CloseMerchant(&pb.CloseMerchantRequest{
		ShopId: merchant.Id,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(resp))

}

func (t *MerchantsController) GetMerchant(c *gin.Context) {
	userId, err := common.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}
	merchant, err := t.logic.GetMerchant(&pb.GetMerchantRequest{
		UserId: userId,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}
	c.JSON(http.StatusOK, t.resp.Success(merchant))
}

func (t *MerchantsController) GetBalance(c *gin.Context) {
	userId, err := common.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}
	balance, err := t.logic.GetBalance(&balancepb.GetBalanceRequest{
		UserId: userId,
		Type:   2,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}
	c.JSON(http.StatusOK, t.resp.Success(balance))
}

type GetBalanceChangeListRequest struct {
	Type       int32 `json:"type" validate:"min=0"`
	TypeAmount int32 `json:"type_amount" validate:"min=0"`
	Page       int32 `json:"page" validate:"required"`
	PageSize   int32 `json:"page_size" validate:"required"`
}

func (t *MerchantsController) GetBalanceChangeList(c *gin.Context) {
	var req GetBalanceChangeListRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}
	userId, err := common.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}
	balance, err := t.logic.GetBalance(&balancepb.GetBalanceRequest{
		UserId: userId,
		Type:   2,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}
	item, err := t.logic.GetBalanceChangeList(&balancepb.GetBalanceChangeListRequest{
		Id:         balance.Id,
		Type:       req.Type,
		TypeAmount: req.TypeAmount,
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(item))
}

func (t *MerchantsController) Withdrawals(c *gin.Context) {

}
