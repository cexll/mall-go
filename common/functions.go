package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"os"
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

func GetCWD() string {
	cwd, _ := os.Getwd()
	return cwd
}

func SubStr(str string, start int, length int) string {
	if length == 0 {
		return str[start:]
	}
	if start == 0 {
		return str[:length]
	}
	return str[start:length]
}
