package controllers

import (
	"fmt"
	"github.com/cexll/mall-go/logic"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"net/http"
	"time"
)

type HelloController struct {
	logic logic.UserLogic
}

type User struct {
	Id        int    `json:"id" xsql:"id" from:"id"`
	Nickname  string `json:"nickname" xsql:"nickname"`
	AvatarUrl string `json:"avatar_url" xsql:"avatar_url"`
	Password  string `json:"password" xsql:"password"`
}

type UserForm struct {
	Name     string    `validate:"required|min_len:7" message:"required:{field} is required" label:"用户名称"`
	Email    string    `validate:"email" message:"email is invalid" label:"用户邮箱"`
	Age      int       `validate:"required|int|min:1|max:99" message:"int:age must int|min:age min value is 1"`
	CreateAt int       `validate:"min:1"`
	Safe     int       `validate:"-"` // 标记字段安全无需验证
	UpdateAt time.Time `validate:"required" message:"update time is required"`
	Code     string    `validate:"customValidator"`
	// 结构体嵌套
	ExtInfo struct {
		Homepage string `validate:"required" label:"用户主页"`
		CityName string
	} `validate:"required" label:"扩展信息"`
}

type TestForm struct {
	T1 int    `validate:"required|int"`
	T2 int32  `validate:"required|int"`
	T3 int64  `validate:"required|int"`
	T4 string `validate:"required"`
}

func (t *HelloController) Index(c *gin.Context) {
	data, err := validate.FromRequest(c.Request)
	if err != nil {
		panic(err)
	}
	v := data.Src()
	fmt.Println(v)
	//v := data.Create()
	//testFrom := &TestForm{}
	//
	//v.BindSafeData(testFrom)
	////v := validate.Struct(testFrom)
	//if v.Validate() {
	//	//v.BindSafeData(testFrom)
	//	fmt.Println(testFrom)
	//} else {
	//	fmt.Println(v.Errors)
	//	fmt.Println(v.Errors.One())
	//}
	//data, err := validate.FromRequest(c.Request)
	//if err != nil {
	//	panic(err)
	//}
	//if err := c.ShouldBind(&test); err != nil {
	//	panic(err)
	//}
	//fmt.Println(test)
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": "hello, world!",
	//	"data":    test,
	//})
	//return
	//users, err := t.logic.FindByWhere()
	//if err != nil {
	//	log.Println(err)
	//}

	//var users []User
	//db := di.Xsql()
	//args := []string{"1", "cexll"}
	//err := db.Find(&users, "select * from mall_user where id = ? and nickname = ?", args)
	//if err != nil {
	//	log.Println(err)
	//}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "hello, world!",
		"data":    "",
	})
}
