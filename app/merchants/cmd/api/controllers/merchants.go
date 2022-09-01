package controllers

import (
	"github.com/gin-gonic/gin"
	"mall-go/app/merchants/cmd/rpc/logic"
	"mall-go/common/response"
)

type MerchantsController struct {
	resp  response.Response
	logic logic.MerchantsLogic
}

func (t *MerchantsController) JoinMerchant(c *gin.Context) {

}

func (t *MerchantsController) UpdateMerchant(c *gin.Context) {

}

func (t *MerchantsController) CloseMerchant(c *gin.Context) {

}

func (t *MerchantsController) GetBalance(c *gin.Context) {

}

func (t *MerchantsController) GetBalanceChangeList(c *gin.Context) {

}

func (t *MerchantsController) Withdrawals(c *gin.Context) {

}
