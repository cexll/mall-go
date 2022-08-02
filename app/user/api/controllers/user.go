package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"mall-go/app/user/pb"
	"mall-go/common/di"
	"mall-go/common/response"
	"time"
)

var userClient pb.UserClient
var ctx context.Context

func init() {
	ctx, _ = context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	conn, err := grpc.DialContext(ctx, ":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	//defer func() {
	//	_ = conn.Close()
	//}()
	userClient = pb.NewUserClient(conn)
}

type UserController struct {
	resp response.Response
}

var validate *validator.Validate

type RegisterValidate struct {
	Mobile   string `json:"mobile" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (t *UserController) Register(c *gin.Context) {
	validate = validator.New()
	var req RegisterValidate
	err := c.BindJSON(&req)
	if err != nil {
		di.Logrus().Error(err)
		c.JSON(200, t.resp.Fail("参数异常"))
		return
	}
	err = validate.Struct(req)
	if err != nil {
		di.Logrus().Error(err)
		c.JSON(200, t.resp.Fail("参数异常"))
		return
	}
	rpcReq := pb.RegisterRequest{
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Password: req.Password,
	}

	rpcResp, err := userClient.Register(ctx, &rpcReq)
	if err != nil {
		di.Logrus().Error(err)
		c.JSON(200, t.resp.Fail("服务调用异常"))
		return
	}

	c.JSON(200, t.resp.Success(rpcResp))
}

func (t *UserController) GetUser(c *gin.Context) {

}

func (t *UserController) SetUser(c *gin.Context) {

}

func (t *UserController) Logout(c *gin.Context) {

}
