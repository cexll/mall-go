package controllers

import (
	"github.com/cexll/mall-go/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	logic logic.UserLogic
}

type FindRequest struct {
	Id int `form:"id" json:"id" binding:"required"`
}

func (t *UserController) Find(c *gin.Context) {
	var req FindRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    1403,
			"message": "参数异常",
			"data":    nil,
		})
		return
	}
	user := t.logic.Find(req.Id)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "message",
		"data":    user,
	})
}

func (t *UserController) Create(c *gin.Context) {
	id := t.logic.CreateUser()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "ok",
		"data":    id,
	})
}
