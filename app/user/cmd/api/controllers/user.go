package controllers

import (
	"mall-go/app/user/cmd/api/logic"
	"mall-go/app/user/cmd/pb"
	"mall-go/common"
	"mall-go/common/di"
	"mall-go/common/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	resp  response.Response
	logic logic.UserLogic
}

type RegisterValidate struct {
	Mobile   string `json:"mobile" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (t UserController) Register(c *gin.Context) {
	var req RegisterValidate
	err := c.Bind(&req)
	if err != nil {
		di.Logrus().Error(err)
		c.JSON(http.StatusOK, t.resp.Fail("参数异常"))
		return
	}

	resp, err := t.logic.Register(&pb.RegisterRequest{
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(resp))
}

type LoginValidate struct {
	Mobile   string `json:"mobile" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Login 登陆
func (t UserController) Login(c *gin.Context) {
	var req LoginValidate
	err := c.Bind(&req)
	if err != nil {
		di.Logrus().Error(err)
		c.JSON(http.StatusOK, t.resp.Fail("参数异常"))
		return
	}

	resp, err := t.logic.Login(&pb.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(resp))
}

type GetUserValidate struct {
	Id string `json:"id" validate:"required"`
}

func (t UserController) GetUser(c *gin.Context) {
	var req GetUserValidate
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail("参数异常"))
		return
	}

	Id, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail("parse int error"))
		return
	}
	resp, err := t.logic.GetUser(&pb.GetUserRequest{
		Id: Id,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(resp))
}

type SetUserValidate struct {
	Nickname  string `json:"nickname" validate:"required"`
	AvatarUrl string `json:"avatar_url" validate:"required"`
	Mobile    string `json:"mobile" validate:"required"`
	Signature string `json:"signature" validate:"required"`
	Status    int64  `json:"status" validate:"required"`
	Password  string `json:"password" validate:"max=16"`
}

func (t UserController) SetUser(c *gin.Context) {
	userId, err := common.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}

	var req SetUserValidate
	err = c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail("参数异常"))
		return
	}

	resp, err := t.logic.SetUser(&pb.SetUserRequest{
		Id:        userId,
		Nickname:  req.Nickname,
		AvatarUrl: req.AvatarUrl,
		Mobile:    req.Mobile,
		Signature: req.Signature,
		Status:    req.Status,
		Password:  req.Password,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(resp))
}

func (t UserController) Logout(c *gin.Context) {
	userId, err := common.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}
	resp, err := t.logic.Logout(&pb.LogOutRequest{
		Id: userId,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Error(err))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(resp))
}
