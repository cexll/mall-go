package controllers

import (
	"mall-go/app/user/cmd/api/logic"
	"mall-go/app/user/cmd/pb"
	"mall-go/common/di"
	"mall-go/common/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	resp  response.Response
	logic logic.UserLogic
}

var validate *validator.Validate

type RegisterValidate struct {
	Mobile   string `json:"mobile" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (t UserController) Register(c *gin.Context) {
	validate = validator.New()
	var req RegisterValidate
	err := c.BindJSON(&req)
	if err != nil {
		di.Logrus().Error(err)
		c.JSON(http.StatusOK, t.resp.Fail("参数异常"))
		return
	}
	err = validate.Struct(req)
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
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
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
	validate = validator.New()
	var req LoginValidate
	err := c.BindJSON(&req)
	if err != nil {
		di.Logrus().Error(err)
		c.JSON(http.StatusOK, t.resp.Fail("参数异常"))
		return
	}
	err = validate.Struct(req)
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
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(resp))
}

type GetUserValidate struct {
	Id string `json:"id" validate:"required"`
}

func (t UserController) GetUser(c *gin.Context) {
	validate = validator.New()
	var req GetUserValidate
	err := c.BindJSON(&req)
	if err != nil {
		di.Logrus().Error(err)
		c.JSON(http.StatusOK, t.resp.Fail("参数异常"))
		return
	}
	err = validate.Struct(req)
	if err != nil {
		di.Logrus().Error(err)
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
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
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
}

func (t UserController) SetUser(c *gin.Context) {
	value, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusOK, t.resp.Fail("获取用户ID失败"))
		return
	}
	userId, err := strconv.ParseInt(value.(string), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail("转换用户ID异常"))
		return
	}
	validate = validator.New()
	var req SetUserValidate
	err = c.BindJSON(&req)
	if err != nil {
		di.Logrus().Error(err)
		c.JSON(http.StatusOK, t.resp.Fail("参数异常"))
		return
	}
	err = validate.Struct(req)
	if err != nil {
		di.Logrus().Error(err)
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
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(resp))
}

func (t UserController) Logout(c *gin.Context) {
	value, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusOK, t.resp.Fail("获取用户ID失败"))
		return
	}
	userId := value.(int64)
	resp, err := t.logic.Logout(&pb.LogOutRequest{
		Id: userId,
	})
	if err != nil {
		c.JSON(http.StatusOK, t.resp.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, t.resp.Success(resp))
}
