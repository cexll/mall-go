package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mall-go/app/user/api/logic"
	"mall-go/app/user/pb"
	"mall-go/common/di"
	"mall-go/common/response"
	"net/http"
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

func (t UserController) GetUser(c *gin.Context) {

}

func (t UserController) SetUser(c *gin.Context) {

}

func (t UserController) Logout(c *gin.Context) {

}
