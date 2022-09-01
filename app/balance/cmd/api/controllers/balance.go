package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall-go/app/balance/cmd/api/logic"
	"mall-go/app/balance/cmd/pb"
	"mall-go/common"
	"mall-go/common/response"
	"net/http"
	"strconv"
)

type BalanceController struct {
	resp  response.Response
	logic logic.BalanceLogic
}

// GetBalance 查询余额
func (t *BalanceController) GetBalance(c *gin.Context) {
	userId, err := common.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}
	ptype, _ := strconv.Atoi(c.PostForm("type"))
	resp, err := t.logic.GetBalance(&pb.GetBalanceRequest{
		UserId: userId,
		Type:   int32(ptype),
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(resp))
}

type GetBalanceChangeListRequest struct {
	Id         int64 `json:"id" validate:"required"`
	Type       int32 `json:"type" validate:"min=0"`
	TypeAmount int32 `json:"type_amount" validate:"min=0"`
	Page       int32 `json:"page" validate:"required"`
	PageSize   int32 `json:"page_size" validate:"required"`
}

// GetBalanceChangeList 查询余额账单
func (t *BalanceController) GetBalanceChangeList(c *gin.Context) {
	var req GetBalanceChangeListRequest
	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, t.resp.Fail("参数异常"))
		return
	}

	resp, err := t.logic.GetBalanceChangeList(&pb.GetBalanceChangeListRequest{
		Id:         req.Id,
		Type:       req.Type,
		TypeAmount: req.TypeAmount,
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(resp))
}

// Withdraw 提现
func (t *BalanceController) Withdraw(c *gin.Context) {

}

// Recharge 充值
func (t *BalanceController) Recharge(c *gin.Context) {

}
