package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserId(c *gin.Context) (int64, error) {
	value, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("获取用户ID失败")
	}
	userId, err := strconv.ParseInt(value.(string), 10, 64)
	if err != nil {
		return 0, errors.New("解析用户ID失败")
	}

	return userId, nil
}
